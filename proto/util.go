package proto

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
	"hash/adler32"
	"io"
	"log"
	"net"
	"strings"

	"google.golang.org/protobuf/proto"
)

type ExtendedResponse[H, T proto.Message] struct {
	Header H
	Resp   T
}

func DiscardHeader[H, T proto.Message](extResp *ExtendedResponse[H, T], err error) (T, error) {
	var t T
	if err != nil {
		return t, err
	}
	return extResp.Resp, nil
}

type internalUnaryReq[REQ proto.Message, RESP any] struct {
	req    REQ
	hdr    *MessageHeader
	respCh chan RESP
	errCh  chan error
}

type internalServerStreamingReq[REQ proto.Message, RESP any] struct {
	req    REQ
	hdr    *MessageHeader
	stream *serverStreamImpl[RESP]
	errCh  chan error
}

type frameWriterReq struct {
	msgTag *MessageTag
	msg    proto.Message
}

func newFrameChannel(addr string) (io.ReadWriteCloser, string, error) {
	// config := &tls.Config{
	// }
	// conn, err := tls.Dial("tcp", addr, config)
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return nil, "", fmt.Errorf("connecting to %s: %v", addr, err)
	}
	macAddr, err := getMACAddr(conn.(*net.TCPConn))
	if err != nil {
		return nil, "", fmt.Errorf("resolving mac address: %v", err)
	}
	return conn, macAddr, nil
}

func frameWriter(frw FrameReadWriter, reqCh <-chan frameWriterReq, errCh chan<- error, log *log.Logger) {
	for req := range reqCh {
		switch {
		case req.msgTag.Req == nil:
			errCh <- fmt.Errorf("failed precondition: request message is missing request tag: %v", req.msgTag)
		default:
			payload, err := proto.Marshal(req.msg)
			if err != nil {
				errCh <- fmt.Errorf("marshalling req message: %v", err)
				continue
			}

			var frame []byte

			if req.msgTag.Header == nil {
				frame = binary.BigEndian.AppendUint16(frame, uint16(req.msgTag.GetReq()))
				frame = append(frame, 0)
			} else {
				hdr := req.msgTag.GetHeader()
				hdrPayload, err := proto.Marshal(hdr)
				if err != nil {
					errCh <- fmt.Errorf("marshalling extended message: %v", err)
					continue
				}
				frame = binary.BigEndian.AppendUint16(frame, uint16(req.msgTag.GetReq()))
				frame = append(frame, 0)
				frame = binary.BigEndian.AppendUint16(frame, uint16(len(hdrPayload)))
				frame = append(frame, hdrPayload...)
			}

			frame = append(frame, payload...)
			frame = binary.BigEndian.AppendUint32(frame, adler32.Checksum(frame))

			if err := frw.WriteFrame(frame); err != nil {
				errCh <- fmt.Errorf("writing frame: %v", err)
			}
			log.Printf("dbg: frameWriter: tag={%v} msg={%v}", req.msgTag, req.msg)
		}
	}
}

func frameReader(frw FrameReadWriter, frameCh chan<- []byte, errCh chan<- error) {
	for {
		frame, err := frw.ReadFrame()
		if err != nil {
			errCh <- err
			return
		}
		frameCh <- frame
	}
}

func getMACAddr(conn *net.TCPConn) (string, error) {
	localIP := conn.LocalAddr().(*net.TCPAddr).IP

	ifs, err := net.Interfaces()
	if err != nil {
		return "", fmt.Errorf("getting network interfaces: %v", err)
	}

	for _, i := range ifs {
		if i.Flags&net.FlagUp != 0 && !bytes.Equal(i.HardwareAddr, nil) {
			addrs, err := i.Addrs()
			if err != nil {
				log.Printf("getting addresses for interface %v: %v", i.Name, err)
				continue
			}
			for _, addr := range addrs {
				if addr.Network() != "ip+net" {
					continue
				}
				if net.IP.Equal(addr.(*net.IPNet).IP, localIP) {
					return strings.ToUpper(hex.EncodeToString(i.HardwareAddr)), nil
				}
			}
		}
	}

	return "", errors.New("MAC address not found")
}
