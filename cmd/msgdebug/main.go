package main

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/protocolbuffers/protoscope"
)

var (
	msg = flag.String("msg", "", "")
)

func indent(s string) string {
	return "  " + strings.ReplaceAll(s, "\n", "\n  ")
}

func singleLine(s string) string {
	return strings.TrimSuffix(strings.ReplaceAll(s, "\n", ", "), ", ")
}

func processOne(data []byte, postprocFn func(string) string) string {
	psOpts := protoscope.WriterOptions{
		NoQuotedStrings:        false,
		AllFieldsAreMessages:   false,
		ExplicitWireTypes:      false,
		NoGroups:               true,
		ExplicitLengthPrefixes: true,
	}

	expectedLen := binary.BigEndian.Uint16(data[0:2])
	if len(data) != int(expectedLen)+2 {
		log.Fatalf("len mismatch: want %d, got %d", expectedLen+2, len(data))
	}

	ret := new(bytes.Buffer)

	tag := binary.BigEndian.Uint16(data[2:4])
	fmt.Fprintf(ret, "tag: %04X\n", tag)

	extendedTags := map[uint16]bool{
		0x6658: true,
	}
	if extendedTags[tag] {
		extMsgLen := binary.BigEndian.Uint16(data[5:7])
		extPayload := data[7 : 7+extMsgLen]
		mainPayload := data[7+extMsgLen : len(data)-4]
		fmt.Fprintln(ret, "ext payload:")
		fmt.Fprintln(ret, postprocFn(protoscope.Write(extPayload, psOpts)))
		fmt.Fprintln(ret, "main payload:")
		fmt.Fprintln(ret, postprocFn(protoscope.Write(mainPayload, psOpts)))
	} else {
		payload := data[5 : len(data)-4]
		fmt.Fprintln(ret, "payload:")
		fmt.Fprintln(ret, postprocFn(protoscope.Write(payload, psOpts)))
	}

	checksum := hex.EncodeToString(data[len(data)-4:])
	fmt.Fprintf(ret, "checksum: %s\n", checksum)

	return ret.String()
}

// func processBuf(ref string, buf *bytes.Buffer) {
// 	for {
// 		if buf.Len() < 2 {
// 			return
// 		}
// 		msgLen := int(binary.BigEndian.Uint16(buf.Bytes()[0:2]))
// 		if buf.Len() < msgLen+2 {
// 			return
// 		}
//
// 		data := make([]byte, msgLen+2)
// 		buf.Read(data)
//
// 		fmt.Printf("%s: %s\n", ref, processOne(data))
// 	}
// }

func main() {
	flag.Parse()

	// clientBuf, serverBuf := new(bytes.Buffer), new(bytes.Buffer)

	// for _, p := range packets {
	// 	if p.side == 0 {
	// 		clientBuf.Write(p.data)
	// 		processBuf("CLI", clientBuf)
	// 	} else {
	// 		serverBuf.Write(p.data)
	// 		processBuf("SRV", serverBuf)
	// 	}
	// }

	data, err := hex.DecodeString(*msg)
	if err != nil {
		log.Fatalf("decoding message bytes: %v", err)
	}

	expectedLen := binary.BigEndian.Uint16(data[0:2])
	if len(data) >= int(expectedLen)+2 {
		fmt.Println(processOne(data[:2+expectedLen], indent))
		rem := data[2+expectedLen:]
		if len(rem) > 0 {
			fmt.Printf("remainder: %s\n", hex.EncodeToString(rem))
		}
	} else {
		log.Fatalf("len mismatch: want %d, got %d", expectedLen+2, len(data))
	}
}
