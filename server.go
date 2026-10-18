package main

import (
	"fmt"
	"net"
	"bufio"
)
type Connection struct {
	conn net.Conn
	
} 
func handleconn(conn net.Conn, database *Database){
		reader:=bufio.NewReader(conn)
		defer conn.Close()
		//fmt.Println("Waiting for command...")
		for {
			n,err:=reader.ReadString('\n')
			//fmt.Println("2. Received:", n)
			if err!=nil {
				break
			}
			parsedCmd:= readCmd(n)
			//fmt.Println("3. Executing command")
			implCmd(parsedCmd, database, Connection{conn:conn})
			//fmt.Println("3. Executing command")
			//fmt.Println("4. Command completed")

		}
		

}
func dummyconn() {
	database:= Database{}
	database.db = make(map[string]string)
	ln, err := net.Listen("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Print(err)

	}
	//buffer := make([]byte, 1024)
	for {
		conn, err := ln.Accept()
		//fmt.Println("1. Client connected")
		if err != nil {
			fmt.Print(err)

		}
		//fmt.Print("go-redis >")
		// n, err := bufio.NewReader(conn).ReadString('\n')
		// if err != nil {
		// 	fmt.Println(err)
		// }
		// parsed_cmd:=readCmd(n)
		// implCmd(parsed_cmd, database, Connection{conn: conn})
		go handleconn(conn, &database)
		//conn.Close()


		//fmt.Print(n)
		// for i := 0; i < len(n); i++ {
		// 	fmt.Print(buffer[i], " ")
		// }

	}

}
