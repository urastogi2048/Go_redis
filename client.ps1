
$client = [System.Net.Sockets.TcpClient]::new("127.0.0.1", 6379)
$stream = $client.GetStream()
$utf8 = [System.Text.Encoding]::UTF8
$ascii = [System.Text.Encoding]::ASCII

function Encode-Request([string[]]$argsList) {
    $memory = [System.IO.MemoryStream]::new()

    try {
        $header = $ascii.GetBytes("*$($argsList.Count)`r`n")
        $memory.Write($header, 0, $header.Length)

        foreach ($arg in $argsList) {
            $data = $utf8.GetBytes($arg)
            $header = $ascii.GetBytes("`$$($data.Length)`r`n")

            $memory.Write($header, 0, $header.Length)
            $memory.Write($data, 0, $data.Length)

            $crlf = $ascii.GetBytes("`r`n")
            $memory.Write($crlf, 0, $crlf.Length)
        }

        return ,$memory.ToArray()
    }
    finally {
        $memory.Dispose()
    }
}

function Read-RespLine {
    $bytes = [System.Collections.Generic.List[byte]]::new()

    while ($true) {
        $b = $stream.ReadByte()

        if ($b -lt 0) {
            throw "Server closed the connection."
        }

        if ($b -eq 13) {
            if ($stream.ReadByte() -ne 10) {
                throw "Invalid RESP CRLF terminator."
            }
            break
        }

        $bytes.Add([byte]$b)
    }

    return $ascii.GetString($bytes.ToArray())
}

function Read-ExactBytes([int]$length) {
    $buffer = [byte[]]::new($length)
    $offset = 0

    while ($offset -lt $length) {
        $count = $stream.Read($buffer, $offset, $length - $offset)

        if ($count -eq 0) {
            throw "Server closed the connection."
        }

        $offset += $count
    }

    return ,$buffer
}

function Decode-Response {
    $prefix = $stream.ReadByte()

    if ($prefix -lt 0) {
        throw "Server closed the connection."
    }

    $line = Read-RespLine

    switch ([char]$prefix) {
        '+' {
            return $line
        }
        '-' {
            return "Server error: $line"
        }
        ':' {
            return [long]::Parse(
                $line,
                [System.Globalization.CultureInfo]::InvariantCulture
            )
        }
        '$' {
            $length = [int]::Parse($line)

            if ($length -eq -1) {
                return "(nil)"
            }

            if ($length -lt 0) {
                throw "Invalid RESP bulk-string length."
            }

            $data = Read-ExactBytes $length

            if ($stream.ReadByte() -ne 13 -or
                $stream.ReadByte() -ne 10) {
                throw "Invalid bulk-string terminator."
            }

            return $utf8.GetString($data)
        }
        default {
            throw "Unsupported RESP response type: $([char]$prefix)"
        }
    }
}

try {
    while ($true) {
        $cmd = Read-Host "redis-go>"

        if ($cmd -eq "exit") {
            break
        }

        if ([string]::IsNullOrWhiteSpace($cmd)) {
            continue
        }

        # Initial version: split commands on whitespace.
        $parts = @($cmd.Trim() -split '\s+')

        # Encode the request into RESP bytes.
        $request = Encode-Request -argsList $parts
        $stream.Write($request, 0, $request.Length)

        # Read and decode the server's RESP response.
        $response = Decode-Response
        Write-Host $response
    }
}
catch {
    Write-Host "Client error: $($_.Exception.Message)"
}
finally {
    $client.Close()
}
