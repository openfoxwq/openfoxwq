package main

import (
	"encoding/binary"
	"encoding/hex"
	"flag"
	"fmt"
	"log"

	"github.com/openfoxwq/openfoxwq/util/msgdbg"
)

var (
	msg = flag.String("msg", "", "")
)

func main() {
	flag.Parse()

	data, err := hex.DecodeString(*msg)
	if err != nil {
		log.Fatalf("decoding message bytes: %v", err)
	}

	expectedLen := binary.BigEndian.Uint16(data[0:2])
	if len(data) >= int(expectedLen)+2 {
		s, err := msgdbg.Process(data[2:2+expectedLen], false)
		if err != nil {
			log.Fatalf("processing message: %v", err)
		}
		fmt.Println(s)
		rem := data[2+expectedLen:]
		if len(rem) > 0 {
			fmt.Printf("remainder: %s\n", hex.EncodeToString(rem))
		}
	} else {
		log.Fatalf("len mismatch: want %d, got %d", expectedLen+2, len(data))
	}
}
