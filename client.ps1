
$client = [System.Net.Sockets.TcpClient]::new("127.0.0.1", 6379)
$stream = $client.GetStream()
$reader = [System.IO.StreamReader]::new($stream)
$writer = [System.IO.StreamWriter]::new($stream)
$writer.AutoFlush = $true

try {
    while ($true) {
        $cmd = Read-Host "redis-go>"

        if ($cmd -eq "exit") {
            break
        }

        if ([string]::IsNullOrWhiteSpace($cmd)) {
            continue
        }

        # Write-Host "Sending: $cmd"
        $writer.WriteLine($cmd)

        $response = $reader.ReadLine()

        if ($null -eq $response) {
            Write-Host "Server closed the connection."
            break
        }

        Write-Host $response
    }
}
finally {
    $client.Close()
}
