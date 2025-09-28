package main

import (
	"fmt"
	"log"
	"net"
	"slices"
	"strings"
	"sync"
)

var AOF *Aof
var SETs = map[string]string{}
var SETsMu = &sync.RWMutex{}

var HSETs = map[string]map[string]string{}
var HSETsMu = &sync.RWMutex{}

var Handlers = map[string]func(args []Value) Value{
	"PING":    PingCommand,
	"SET":     SetCommand,
	"GET":     GetCommand,
	"HSET":    HSetCommand,
	"HGET":    HGetCommand,
	"HGETALL": HGetAllCommand,
	"COMMAND": CommandCommand,
	"FLUSH":   FlushCommand,
}

func main() {
	aof, err := NewAof("database.aof")
	if err != nil {
		fmt.Println(err)
		return
	}
	AOF = aof
	defer AOF.Close()

	err = AOF.Read(func(value Value) {
		command := strings.ToUpper(value.array[0].bulk)
		args := value.array[1:]

		handler, ok := Handlers[command]
		if !ok {
			fmt.Println("Invalid command: ", command)
			return
		}

		handler(args)
	})
	if err != nil {
		fmt.Println("Error reading AOF: ", err)
		return
	}

	l, err := net.Listen("tcp", ":6379")
	if err != nil {
		log.Println("Error starting server: ", err)
		return
	}
	log.Println("Listening on port :6379")

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Println("Error accepting connection: ", err)
			continue
		}

		go handle(conn)
	}
}

var DontStorageCmds = []string{
	"PING",
	"GET",
	"HGET",
	"HGETALL",
	"COMMAND",
	"FLUSH",
}

func handle(conn net.Conn) {
	defer conn.Close()

	for {
		resp := NewResp(conn)
		value, err := resp.Read()
		if err != nil {
			log.Println("Error reading from connection: ", err)
			return
		}

		if value.typ != "array" {
			log.Println("Invalid request, expected array")
			continue
		}

		if len(value.array) == 0 {
			log.Println("Invalid request, expected array length > 0")
			continue
		}

		command := strings.ToUpper(value.array[0].bulk)
		args := value.array[1:]

		writer := NewWriter(conn)

		handler, ok := Handlers[command]
		if !ok {
			log.Println("Invalid command:", command)
			log.Println("Args", args)
			writer.Write(Value{typ: "string", str: ""})
			continue
		}

		if !slices.Contains(DontStorageCmds, command) {
			err = AOF.Write(value)
			if err != nil {
				log.Println("Error writing to AOF: ", err)
			}
		}

		result := handler(args)
		err = writer.Write(result)
		if err != nil {
			log.Println("Error writing to connection: ", err)
		}
	}
}
