package main

import (
	"fmt"
	"strings"
)

type Command struct {
	name string
	args []string
}

func readCmd(input string) (c Command) {

	//cmd:=[]string
	cmd := strings.Fields(input)
	//x:=0
	for i := 0; i < len(cmd); i++ {
		if i == 0 {
			c.name = cmd[i]
		} else {
			c.args = append(c.args, cmd[i])
		}

	}
	return c
}
func implCmd(input Command, db *Database,connection Connection) {
	
	if input.name == "SET" {
		if len(input.args) != 2 {
			fmt.Fprintln(connection.conn, "Invalid command")
			return
		}
		key := input.args[0]
		val := input.args[1]

		db.Set(key, val)
		//fmt.Fprintln(connection.conn, "OK")
		connection.conn.Write([]byte("+OK\r\n"))

	} else if input.name == "GET" {
		if len(input.args) != 1 {
			writeString(connection.conn, "length of input not matched")
			return
		}
		key := input.args[0]
		val, e := db.Get(key)
		result := val
		if !e {
			writeString(connection.conn,"invalid request")
		} else {
			writeBulkString(connection.conn,result)
		}
	} else if input.name == "DEL" {
		if len(input.args) != 1 {
			writeString(connection.conn, "invalid request length")
			return
		}
		key := input.args[0]
		f := db.Del(key)
		if f {
			connection.conn.Write([]byte("+OK\r\n"))

		} else {
			writeString(connection.conn, "Invalid key")
		}
	} else {
		writeString(connection.conn, "Invalid key")
		return
	}

}
