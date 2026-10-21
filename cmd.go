package main

import (
	"strconv"
	"strings"
	"time"
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
func implCmd(input Command, db *Database, connection Connection) {

	if input.name == "SET" {
		if len(input.args) != 2 {
			if len(input.args) == 4 && input.args[2] == "EX" {
				a, e := strconv.Atoi(input.args[3])
				//a=time(a)
				if e == nil {
					entry := Entry{
						value:     input.args[0],
						expiresIn: time.Now().Add(time.Duration(a) * time.Second),
					}
					db.SetWithTTL(input.args[0], input.args[1], entry.expiresIn)
					connection.conn.Write([]byte("+OK\r\n"))
					return
				}
				writeString(connection.conn, "Invalid command")
				return
			}
			writeString(connection.conn, "Invalid command")
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
			writeString(connection.conn, "The given key does not exist")
		} else {
			writeBulkString(connection.conn, result)
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
	} else if input.name == "EXPIRE" {
		if len(input.args) != 2 {
			writeString(connection.conn, "invalid request length")
			return
		}
		key := input.args[0]
		//a:=input.args[1]
		a, e := strconv.Atoi(input.args[1])
		if e == nil {
			if db.TTLApply(key, time.Now().Add(time.Duration(a)*time.Second)) {
				writeString(connection.conn, "SUCCESS")
			} else {
				writeInteger(connection.conn, 0)
			}
			return
		} else {
			writeString(connection.conn, "Invalid key")
			return
		}
	} else if input.name == "TTL" {
		if len(input.args) != 1 {
			writeString(connection.conn, "Invalid command")
			return
		} else {
			a := db.TTLGet(input.args[0])
			if a < 0 {
				writeString(connection.conn, "The key has expired")
				return
			}
			writeInteger(connection.conn, a)
		}

	} else {
		writeString(connection.conn, "Invalid key")
		return
	}

}
