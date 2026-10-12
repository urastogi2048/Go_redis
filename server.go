package main

import (
	"fmt"
	"net"
	"bufio"
)

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
		if err != nil {
			fmt.Print(err)

		}
		fmt.Print("go-redis >")
		n, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println(err)
		}
		parsed_cmd:=readCmd(n)
		implCmd(parsed_cmd, database)

		//fmt.Print(n)
		// for i := 0; i < len(n); i++ {
		// 	fmt.Print(buffer[i], " ")
		// }

	}

}
