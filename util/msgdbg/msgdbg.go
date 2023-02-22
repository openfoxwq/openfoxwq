package msgdbg

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"hash/adler32"
	"strings"

	"github.com/protocolbuffers/protoscope"
)

func indent(s string) string {
	return "  " + strings.ReplaceAll(s, "\n", "\n  ")
}

func singleLine(s string) string {
	return strings.TrimSuffix(strings.ReplaceAll(s, "\n", ", "), ", ")
}

func Process(data []byte, compact bool) (string, error) {
	if len(data) < 6 {
		return "", fmt.Errorf("frame is too short")
	}

	checksum := binary.BigEndian.Uint32(data[len(data)-4:])
	if want := adler32.Checksum(data[:len(data)-4]); checksum != want {
		return "", fmt.Errorf("checksum doesn't match: got=%08X, want=%08X", checksum, want)
	}

	psOpts := protoscope.WriterOptions{
		NoQuotedStrings:        false,
		AllFieldsAreMessages:   false,
		ExplicitWireTypes:      false,
		NoGroups:               true,
		ExplicitLengthPrefixes: false,
	}

	ret := new(bytes.Buffer)

	tag := binary.BigEndian.Uint16(data[0:2])
	if compact {
		fmt.Fprintf(ret, "[t:%04X] ", tag)
	} else {
		fmt.Fprintf(ret, "tag: %04X\n", tag)
	}

	extendedTags := map[uint16]bool{
		0x6658: true,
	}
	if extendedTags[tag] {
		extMsgLen := binary.BigEndian.Uint16(data[3:5])
		extPayload := data[5 : 5+extMsgLen]
		mainPayload := data[5+extMsgLen : len(data)-4]
		if compact {
			fmt.Fprintf(ret, "hdr:{%s} payload:{%s}", singleLine(protoscope.Write(extPayload, psOpts)), singleLine(protoscope.Write(mainPayload, psOpts)))
		} else {
			fmt.Fprintln(ret, "ext payload:")
			fmt.Fprintln(ret, indent(protoscope.Write(extPayload, psOpts)))
			fmt.Fprintln(ret, "main payload:")
			fmt.Fprintln(ret, indent(protoscope.Write(mainPayload, psOpts)))
		}
	} else {
		payload := data[3 : len(data)-4]
		if compact {
			fmt.Fprintf(ret, "payload:{%s}", singleLine(protoscope.Write(payload, psOpts)))
		} else {
			fmt.Fprintln(ret, "payload:")
			fmt.Fprintln(ret, indent(protoscope.Write(payload, psOpts)))
		}
	}

	if compact {
		fmt.Fprintf(ret, " sum:%08X", checksum)
	} else {
		fmt.Fprintf(ret, "checksum: %08X\n", checksum)
	}

	return ret.String(), nil
}
