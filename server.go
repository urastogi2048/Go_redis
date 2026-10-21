package main

import (
	"bufio"
	"fmt"
	"net"
	"time"
)
type Connection struct {
	conn net.Conn
	
} 
func handleconn(conn net.Conn, database *Database){
		reader:=bufio.NewReader(conn)
		defer conn.Close()
		for {
			parsedCmd,e:= readRESPcmd(reader)
			if e!=nil {
				 return 
			}
			cmd:=Command{
				name: parsedCmd[0],
				args: parsedCmd[1:],
			}
			implCmd(cmd, database, Connection{conn:conn})
			

		}
		

}
func dummyconn() {
	database:= Database{}
	database.db = make(map[string]Entry)
	ln, err := net.Listen("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Print(err)

	}
	database.StartActiveExpiration(1*time.Second)
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Print(err)

		}
		
		go handleconn(conn, &database)
	

	}

}
