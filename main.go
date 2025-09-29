package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"slices"
	"strings"
	"sync"
)

var AOF *Aof
var KEYs = map[string]ValueType{}
var KEYsMu = &sync.RWMutex{}

var SETs = map[string]string{}
var SETsMu = &sync.RWMutex{}

var HSETs = map[string]map[string]string{}
var HSETsMu = &sync.RWMutex{}

var Handlers = map[string]func(args []Value) Value{
	"COMMAND": CommandCommand,
	"DEL":     DelCommand,
	"EXISTS":  ExistsCommand,
	"FLUSH":   FlushCommand,
	"GET":     GetCommand,
	"HGET":    HGetCommand,
	"HGETALL": HGetAllCommand,
	"HSET":    HSetCommand,
	"PING":    PingCommand,
	"SET":     SetCommand,
}

var aofFilePath = flag.String("aof", "database.aof", "Path to the AOF file")

func init() {
	flag.Parse()

	var err error
	AOF, err = NewAof(*aofFilePath)
	if err != nil {
		log.Fatalln(fmt.Errorf("Error initializing AOF: %v", err))
	}

	err = AOF.Read(func(value Value) {
		command := strings.ToUpper(value.array[0].bulk)
		args := value.array[1:]

		handler, ok := Handlers[command]
		if !ok {
			// this should never happen
			log.Println("Invalid command:", command)
			return
		}
		handler(args)
	})
	if err != nil {
		log.Fatalln("Error reading AOF:", err)
	}
}

func main() {
	defer AOF.Close()

	l, err := net.Listen("tcp", ":6379")
	if err != nil {
		log.Println("Error starting server:", err)
		return
	}

	log.Println("Listening on port :6379")

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Println("Error accepting connection:", err)
			continue
		}

		go handle(conn)
	}
}

var DontStoreCmds = []string{
	"COMMAND",
	// DEL
	"EXISTS",
	"FLUSH",
	"GET",
	"HGET",
	"HGETALL",
	// HSET
	"PING",
	// SET
}

func handle(conn net.Conn) {
	defer conn.Close()

	for {
		resp := NewResp(conn)
		value, err := resp.Read()
		if err != nil {
			log.Println("Error reading from connection:", err)
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
			writer.Write(Value{typ: "string", str: fmt.Sprintf("ERR unknown command '%s'", command)})
			continue
		}

		if !slices.Contains(DontStoreCmds, command) {
			err = AOF.Write(value)
			if err != nil {
				log.Println("Error writing to AOF:", err)
			}
		}

		result := handler(args)
		err = writer.Write(result)
		if err != nil {
			log.Println("Error writing to connection:", err)
		}
	}
}
