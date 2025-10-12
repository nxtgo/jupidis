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

var VALUEs = map[string]string{}
var VALUEsMu = &sync.RWMutex{}

var HSETs = map[string]map[string]string{}
var HSETsMu = &sync.RWMutex{}

var SETs = map[string][]string{}
var SETsMu = &sync.RWMutex{}

type CommandValue struct {
	Run   func(args []Value) Value
	Check func(args []Value) bool
}

var Handlers = map[string]CommandValue{
	"APPEND":      {Run: AppendCommand, Check: AppendCommandCheck},
	"COMMAND":     {Run: CommandCommand, Check: CommandCommandCheck},
	"DECR":        {Run: DecrCommand, Check: DecrCommandCheck},
	"DECRBY":      {Run: DecrByCommand, Check: DecrByCommandCheck},
	"DEL":         {Run: DelCommand, Check: DelCommandCheck},
	"EXISTS":      {Run: ExistsCommand, Check: ExistsCommandCheck},
	"FLUSH":       {Run: FlushCommand, Check: FlushCommandCheck},
	"GET":         {Run: GetCommand, Check: GetCommandCheck},
	"HGET":        {Run: HGetCommand, Check: HGetCommandCheck},
	"HGETALL":     {Run: HGetAllCommand, Check: HGetAllCommandCheck},
	"HSET":        {Run: HSetCommand, Check: HSetCommandCheck},
	"INCR":        {Run: IncrCommand, Check: IncrCommandCheck},
	"INCRBY":      {Run: IncrByCommand, Check: IncrByCommandCheck},
	"PING":        {Run: PingCommand, Check: PingCommandCheck},
	"SADD":        {Run: SAddCommand, Check: SAddCommandCheck},
	"SCARD":       {Run: SCardCommand, Check: SCardCommandCheck},
	"SDIFF":       {Run: SDiffCommand, Check: SDiffCommandCheck},
	"SDIFFSTORE":  {Run: SDiffStoreCommand, Check: SDiffStoreCommandCheck},
	"SET":         {Run: SetCommand, Check: SetCommandCheck},
	"SINTER":      {Run: SInterCommand, Check: SInterCommandCheck},
	"SINTERSTORE": {Run: SInterStoreCommand, Check: SInterStoreCommandCheck},
	"SISMEMBER":   {Run: SIsMemberCommand, Check: SIsMemberCommandCheck},
	"SMEMBERS":    {Run: SMembersCommand, Check: SMembersCommandCheck},
	"SMISMEMBER":  {Run: SMIsMemberCommand, Check: SMIsMemberCommandCheck},
	"SMOVE":       {Run: SMoveCommand, Check: SMoveCommandCheck},
	"SREM":        {Run: SRemCommand, Check: SRemCommandCheck},
	"SUNION":      {Run: SUnionCommand, Check: SUnionCommandCheck},
	"SUNIONSTORE": {Run: SUnionStoreCommand, Check: SUnionStoreCommandCheck},
	"TYPE":        {Run: TypeCommand, Check: TypeCommandCheck},
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
		command := strings.ToUpper(value.array[0].str)
		args := value.array[1:]

		handler, ok := Handlers[command]
		if !ok {
			// this should never happen
			log.Println("Invalid command:", command)
			return
		}
		if !handler.Check(args) {
			// this should never happen
			log.Println("Invalid arguments for command:", command)
			return
		}
		handler.Run(args)
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
	// "APPEND",
	"COMMAND",
	// "DECR",
	// "DECRBY",
	// "DEL",
	"EXISTS",
	"FLUSH",
	"GET",
	"HGET",
	"HGETALL",
	// "HSET",
	// "INCR",
	// "INCRBY",
	"PING",
	// "SADD",
	"SCARD",
	// "SET",
	"SISMEMBER",
	"SMEMBERS",
	"SMISMEMBER",
	// "SMOVE",
	// "SREM",
	"TYPE",
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

		command := strings.ToUpper(value.array[0].str)
		args := value.array[1:]

		writer := NewWriter(conn)

		handler, ok := Handlers[command]
		if !ok {
			writer.Write(Value{typ: "string", str: fmt.Sprintf("ERR unknown command '%s'", command)})
			continue
		}

		if !handler.Check(args) {
			writer.Write(Value{typ: "string", str: fmt.Sprintf("ERR wrong number of arguments for '%s' command", command)})
			continue
		}

		if !slices.Contains(DontStoreCmds, command) {
			err = AOF.Write(value)
			if err != nil {
				log.Println("Error writing to AOF:", err)
			}
		}

		result := handler.Run(args)
		err = writer.Write(result)
		if err != nil {
			log.Println("Error writing to connection:", err)
		}
	}
}
