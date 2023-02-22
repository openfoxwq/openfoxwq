// capture captures all traffic in the given interface and filters what is exchanged by the fox client and the server.
package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"net"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"github.com/openfoxwq/openfoxwq/proto"
	"github.com/openfoxwq/openfoxwq/util/msgdbg"
)

var (
	in      = flag.String("in", "en0", "")
	snapLen = flag.Int("snaplen", 1600, "")
)

type frameReadWriter interface {
	proto.FrameReadWriter
	io.Writer
}

type stream struct {
	srcIP, srcPort string
	dstIP, dstPort string
	frw            frameReadWriter
	buf            bytes.Buffer
}

func main() {
	flag.Parse()

	ifs, err := net.Interfaces()
	if err != nil {
		log.Fatalf("listing interfaces: %v", err)
	}

	var names []string
	for _, i := range ifs {
		names = append(names, i.Name)
	}
	log.Printf("interfaces: %v", names)

	log.Printf("starting capture on %s...", *in)

	handle, err := pcap.OpenLive(*in, int32(*snapLen), true, pcap.BlockForever)
	if err != nil {
		log.Fatalf("creating capture handle: %v", err)
	}

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	streams := map[string]*stream{}
	ignoredStreams := map[string]bool{}
	for packet := range packetSource.Packets() {
		if tl := packet.TransportLayer(); tl == nil || tl.LayerType() != layers.LayerTypeTCP || packet.ApplicationLayer() == nil {
			continue
		}

		srcIP := packet.NetworkLayer().NetworkFlow().Src().String()
		dstIP := packet.NetworkLayer().NetworkFlow().Dst().String()
		srcPort := packet.TransportLayer().TransportFlow().Src().String()
		dstPort := packet.TransportLayer().TransportFlow().Dst().String()
		data := packet.ApplicationLayer().Payload()

		if len(data) == 0 {
			continue
		}

		streamKey := fmt.Sprintf("%s:%s-%s:%s", srcIP, srcPort, dstIP, dstPort)

		if ignoredStreams[streamKey] {
			continue
		}

		s := streams[streamKey]
		if s == nil {
			s = &stream{
				srcIP:   srcIP,
				srcPort: srcPort,
				dstIP:   dstIP,
				dstPort: dstPort,
			}
			s.frw = proto.NewLenEncodedFraming(&s.buf)
			streams[streamKey] = s
		}

		s.frw.Write(data)
		for s.frw.HasFrame() {
			frame, err := s.frw.ReadFrame()
			if err != nil {
				log.Fatalf("reading frame from stream %s: %v", streamKey, err)
			}
			if s, err := msgdbg.Process(frame, true); err != nil {
				ignoredStreams[streamKey] = true
				delete(streams, streamKey)
			} else {
				log.Printf("[%s:%s -> %s:%s]: %s", srcIP, srcPort, dstIP, dstPort, s)
			}
		}
	}
}
