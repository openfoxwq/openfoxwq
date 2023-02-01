// Auto-generated by protoc-gen-go-foxwqrpc plugin. DO NOT EDIT.

package proto

import (
	"context"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"log"

	"google.golang.org/protobuf/proto"
)

type FoxClient interface {
	GetNavInfo(context.Context, *GetNavInfoRequest) (*GetNavInfoResponse, error)
	MACAddress() string
}

var (
	fox_getNavInfoMsgTag = &MessageTag{}
)

func init() {
	if err := proto.Unmarshal([]byte{0x08, 0xE8, 0x07, 0x10, 0xE9, 0x07}, fox_getNavInfoMsgTag); err != nil {
		log.Fatalf("loading message tag for method GetNavInfo: %v", err)
	}
}

type foxClientImpl struct {
	macAddr   string
	reqCh     chan frameWriterReq
	inFrameCh chan []byte
	errCh     chan error

	getNavInfoRequestCh chan internalUnaryReq[*GetNavInfoRequest, *GetNavInfoResponse]
}

func NewFoxClient(addr string, log *log.Logger) (FoxClient, error) {
	rwc, macAddr, err := newFrameChannel(addr)
	if err != nil {
		return nil, fmt.Errorf("connecting to %s: %v", addr, err)
	}

	frw := newLenEncodedFraming(rwc)
	inFrameCh := make(chan []byte, 8)
	reqCh := make(chan frameWriterReq, 8)
	errCh := make(chan error, 8)

	go frameWriter(frw, reqCh, errCh, log)
	go frameReader(frw, inFrameCh, errCh)

	getNavInfoRequestCh := make(chan internalUnaryReq[*GetNavInfoRequest, *GetNavInfoResponse], 16)
	getNavInfoPending := &Queue[internalUnaryReq[*GetNavInfoRequest, *GetNavInfoResponse]]{}
	go func() {
		var routerErr error
		for {
			select {
			case frame := <-inFrameCh:
				tag := uint32(binary.BigEndian.Uint16(frame[:2]))
				switch tag {
				case 0x03E9:
					payload := frame[3 : len(frame)-4]
					resp := &GetNavInfoResponse{}
					if err := proto.Unmarshal(payload, resp); err != nil {
						log.Printf("unmarshalling GetNavInfoResponse response: %v\n\tframe: %04X%s", err, len(frame), hex.EncodeToString(frame))
						routerErr = err
						continue
					}
					if !getNavInfoPending.Empty() {
						rr := getNavInfoPending.Pop()
						go func() { rr.respCh <- resp }()
					} else {
						log.Printf("FoxClient: dead letter: %s { %v }", resp.ProtoReflect().Descriptor().Name(), resp)
					}
				default:
					log.Printf("FoxClient: dead letter: tag=%04X frame=%04X%s", tag, len(frame), hex.EncodeToString(frame))
				}
			case routerErr = <-errCh:
				log.Printf("FoxClient: router error set: %v", routerErr)
			case rr := <-getNavInfoRequestCh:
				if routerErr != nil {
					rr.errCh <- routerErr
				}
				reqCh <- frameWriterReq{
					msgTag: fox_getNavInfoMsgTag,
					msg:    rr.req,
				}
				getNavInfoPending.Push(rr)
			}
		}
	}()

	return &foxClientImpl{
		macAddr:             macAddr,
		inFrameCh:           inFrameCh,
		reqCh:               reqCh,
		errCh:               errCh,
		getNavInfoRequestCh: getNavInfoRequestCh,
	}, nil
}

func (c *foxClientImpl) GetNavInfo(_ context.Context, req *GetNavInfoRequest) (*GetNavInfoResponse, error) {
	rr := internalUnaryReq[*GetNavInfoRequest, *GetNavInfoResponse]{
		req:    req,
		respCh: make(chan *GetNavInfoResponse, 1),
		errCh:  make(chan error),
	}
	defer close(rr.respCh)
	defer close(rr.errCh)
	c.getNavInfoRequestCh <- rr
	select {
	case resp := <-rr.respCh:
		return resp, nil
	case err := <-rr.errCh:
		return nil, err
	}
}

func (c *foxClientImpl) MACAddress() string { return c.macAddr }
