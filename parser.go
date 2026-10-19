package main

//resp
import (
	"bufio"
	"errors"
	"io"
	

	"fmt"
	"strconv"
	"strings"
)

func readRESPLine(input *bufio.Reader) (string,error){
	//var l=len(input)
	var output string
	//var line string
	line,e:=input.ReadString('\n')
	if e!=nil {
		return "",e
	}else if len(line)<2 {return "",errors.New("err in readRESPLine")}else if !strings.HasSuffix(line, "\r\n") {
		return "", errors.New("err in readRESPLine")
	}
	output = line[:len(line)-2]
	return output,nil
}

func readLen (input string, symbol string)(int, error){
	
	if !strings.HasPrefix(input, symbol){
		return 0,errors.New("err in readLen")
	}
	line ,e := strconv.Atoi(input[1:])
	if(e!=nil){
		return 0,errors.New("err in readLen")
	}else if strings.HasPrefix(input, symbol) && line>0{
		return  line,nil
	}else {
		return 0,errors.New("err in readLen")
	}
}

func readBulkString(input *bufio.Reader) (string,error){
	header,e:=readRESPLine(input)
	if(e!=nil){
		return "",e
	}
	if !strings.HasPrefix(header, "$"){
		return "",errors.New("invalid bulk string header")

	}
	length, e:=strconv.Atoi(header[1:])
	if e!=nil || length<0{
		return "", errors.New("invalid bulk str len")

	}
	data:=make([]byte,length)
	_,err:=io.ReadFull(input,data)
	if err!=nil{
		return "",err
	}
	terminator:=make([]byte,2)
	_,err=io.ReadFull(input,terminator)
	if err!=nil {
		return "",err
	}
	if string(terminator)!="\r\n"{
		return "",errors.New("invalid terminator")
	}
	return string(data),nil



}
func readRESPcmd(input *bufio.Reader)([]string,error){
	header,e:=readRESPLine(input)
	if e!=nil{
		return nil, e
	}
	length,e1:=readLen(header,"*")
	if e1!=nil{
		return nil, e1
	}
	data:=[]string{}
	for i:=0; i<length; i++{

		a,e2:=readBulkString(input)
		if e2!=nil{
		return nil, e2
		}
		data=append(data,a)

	}
	return data,nil




}

func writeString(w io.Writer, value string ) error {
	_,err:=fmt.Fprintf(w,"+%s\r\n",value)
	return err
}
func writeError(w io.Writer, message string)error {
	_,err:=fmt.Fprintf(w,"-ERR %s\r\n",message)
	return err
}
func writeInteger(w io.Writer, value int) error {
	_, err := fmt.Fprintf(w, ":%d\r\n", value)
	return err
}

func writeBulkString(w io.Writer, value string) error {
	_, err := fmt.Fprintf(w, "$%d\r\n%s\r\n", len([]byte(value)), value)
	return err
}

func writeNullBulkString(w io.Writer) error {
	_, err := io.WriteString(w, "$-1\r\n")
	return err
}
