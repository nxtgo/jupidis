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

var LISTs = map[string][]string{}
var LISTsMu = &sync.RWMutex{}

type CommandValue struct {
	Run   func(args []Value) Value
	Check func(args []Value) error
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
	"KEYS":        {Run: KeysCommand, Check: KeysCommandCheck},
	"LPUSH":       {Run: LPushCommand, Check: LPushCommandCheck},
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
var port = flag.Int("port", 6379, "Port to listen on")
var debugMode = flag.Bool("debug", false, "Enable debug mode")

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
		if err := handler.Check(args); err != nil {
			log.Println("Error checking command from AOF:", err)
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

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Println("Error starting server:", err)
		return
	}

	log.Println("Listening on port :", *port)

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
	"KEYS",
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
			if *debugMode {
				log.Println("Unknown command:", command)
			}
			writer.Write(Value{typ: "string", str: fmt.Sprintf("ERR unknown command '%s'", command)})
			continue
		}

		if err := handler.Check(args); err != nil {
			if *debugMode {
				log.Println("Error checking command:", err, command, args)
			}
			writer.Write(Value{typ: "string", str: fmt.Sprintf("ERR %v", err)})
			continue
		}

		if !slices.Contains(DontStoreCmds, command) {
			err = AOF.Write(value)
			if err != nil {
				log.Println("Error writing to AOF:", err)
			}
		}

		if *debugMode {
			log.Printf("Executing command: %s with args: %v\n", command, args)
		}

		result := handler.Run(args)
		err = writer.Write(result)
		if err != nil {
			log.Println("Error writing to connection:", err)
		}
	}
}
