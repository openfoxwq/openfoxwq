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

type BroadcastClient interface {
	Login(context.Context, *LoginBroadcastServerRequest) (*LoginBroadcastServerResponse, error)
	ListBroadcasts(context.Context, *ListBroadcastsRequest) (ServerStream[*ListBroadcastsResponse], error)
	EnterBroadcast(context.Context, *EnterBroadcastRequest) (*EnterBroadcastResponse, error)
	ListenBroadcastSettingEvents(context.Context) (ServerStream[*BroadcastSettingEvent], error)
	ListenBroadcastStateEvents(context.Context) (ServerStream[*BroadcastStateEvent], error)
	ListenBroadcastMoveEvents(context.Context) (ServerStream[*BroadcastMoveEvent], error)
	ListenBroadcastGameResultEvents(context.Context) (ServerStream[*BroadcastGameResultEvent], error)
	ListenBroadcastTimeControlEvents(context.Context) (ServerStream[*BroadcastTimeControlEvent], error)
	LeaveBroadcast(context.Context, *LeaveBroadcastRequest) (*LeaveBroadcastResponse, error)
	Heartbeat(context.Context, *HeartbeatRequest) error
	UnknownMethod1(context.Context, *UnknownRequest1) error
	UnknownMethod2(context.Context, *UnknownRequest2) error
	UnknownMethod3(context.Context, *UnknownRequest3) error
	UnknownMethod4(context.Context, *UnknownRequest4) error
	UnknownMethod5(context.Context, *UnknownRequest5) error
	UnknownMethod6(context.Context, *UnknownRequest6) (*UnknownResponse6, error)
	UnknownMethod7(context.Context, *UnknownRequest7) (*UnknownResponse7, error)
	UnknownMethod8(context.Context, *UnknownRequest8) (*UnknownResponse8, error)
	MACAddress() string
}

var (
	broadcast_loginMsgTag                            = &MessageTag{}
	broadcast_listBroadcastsMsgTag                   = &MessageTag{}
	broadcast_enterBroadcastMsgTag                   = &MessageTag{}
	broadcast_listenBroadcastSettingEventsMsgTag     = &MessageTag{}
	broadcast_listenBroadcastStateEventsMsgTag       = &MessageTag{}
	broadcast_listenBroadcastMoveEventsMsgTag        = &MessageTag{}
	broadcast_listenBroadcastGameResultEventsMsgTag  = &MessageTag{}
	broadcast_listenBroadcastTimeControlEventsMsgTag = &MessageTag{}
	broadcast_leaveBroadcastMsgTag                   = &MessageTag{}
	broadcast_heartbeatMsgTag                        = &MessageTag{}
	broadcast_unknownMethod1MsgTag                   = &MessageTag{}
	broadcast_unknownMethod2MsgTag                   = &MessageTag{}
	broadcast_unknownMethod3MsgTag                   = &MessageTag{}
	broadcast_unknownMethod4MsgTag                   = &MessageTag{}
	broadcast_unknownMethod5MsgTag                   = &MessageTag{}
	broadcast_unknownMethod6MsgTag                   = &MessageTag{}
	broadcast_unknownMethod7MsgTag                   = &MessageTag{}
	broadcast_unknownMethod8MsgTag                   = &MessageTag{}
)

func init() {
	if err := proto.Unmarshal([]byte{0x08, 0xA0, 0x1F, 0x10, 0xA5, 0x1F}, broadcast_loginMsgTag); err != nil {
		log.Fatalf("loading message tag for method Login: %v", err)
	}
	if err := proto.Unmarshal([]byte{0x08, 0xB4, 0x8D, 0x01, 0x10, 0xCC, 0x1F}, broadcast_listBroadcastsMsgTag); err != nil {
		log.Fatalf("loading message tag for method ListBroadcasts: %v", err)
	}
	if err := proto.Unmarshal([]byte{0x08, 0xC4, 0x27, 0x10, 0xC9, 0x27}, broadcast_enterBroadcastMsgTag); err != nil {
		log.Fatalf("loading message tag for method EnterBroadcast: %v", err)
	}
	if err := proto.Unmarshal([]byte{0x10, 0x99, 0x28}, broadcast_listenBroadcastSettingEventsMsgTag); err != nil {
		log.Fatalf("loading message tag for method ListenBroadcastSettingEvents: %v", err)
	}
	if err := proto.Unmarshal([]byte{0x10, 0x9B, 0x28}, broadcast_listenBroadcastStateEventsMsgTag); err != nil {
		log.Fatalf("loading message tag for method ListenBroadcastStateEvents: %v", err)
	}
	if err := proto.Unmarshal([]byte{0x10, 0xDE, 0x36}, broadcast_listenBroadcastMoveEventsMsgTag); err != nil {
		log.Fatalf("loading message tag for method ListenBroadcastMoveEvents: %v", err)
	}
	if err := proto.Unmarshal([]byte{0x10, 0x92, 0x38}, broadcast_listenBroadcastGameResultEventsMsgTag); err != nil {
		log.Fatalf("loading message tag for method ListenBroadcastGameResultEvents: %v", err)
	}
	if err := proto.Unmarshal([]byte{0x10, 0x9E, 0x38}, broadcast_listenBroadcastTimeControlEventsMsgTag); err != nil {
		log.Fatalf("loading message tag for method ListenBroadcastTimeControlEvents: %v", err)
	}
	if err := proto.Unmarshal([]byte{0x08, 0xD8, 0x27, 0x10, 0xDD, 0x27}, broadcast_leaveBroadcastMsgTag); err != nil {
		log.Fatalf("loading message tag for method LeaveBroadcast: %v", err)
	}
	if err := proto.Unmarshal([]byte{0x08, 0x64}, broadcast_heartbeatMsgTag); err != nil {
		log.Fatalf("loading message tag for method Heartbeat: %v", err)
	}
	if err := proto.Unmarshal([]byte{0x08, 0xE4, 0x5D}, broadcast_unknownMethod1MsgTag); err != nil {
		log.Fatalf("loading message tag for method UnknownMethod1: %v", err)
	}
	if err := proto.Unmarshal([]byte{0x08, 0xB0, 0x6D}, broadcast_unknownMethod2MsgTag); err != nil {
		log.Fatalf("loading message tag for method UnknownMethod2: %v", err)
	}
	if err := proto.Unmarshal([]byte{0x08, 0x9A, 0x47}, broadcast_unknownMethod3MsgTag); err != nil {
		log.Fatalf("loading message tag for method UnknownMethod3: %v", err)
	}
	if err := proto.Unmarshal([]byte{0x08, 0xD1, 0x73}, broadcast_unknownMethod4MsgTag); err != nil {
		log.Fatalf("loading message tag for method UnknownMethod4: %v", err)
	}
	if err := proto.Unmarshal([]byte{0x08, 0xB9, 0x8D, 0x01}, broadcast_unknownMethod5MsgTag); err != nil {
		log.Fatalf("loading message tag for method UnknownMethod5: %v", err)
	}
	if err := proto.Unmarshal([]byte{0x08, 0xA6, 0xEA, 0x01, 0x10, 0xA7, 0xEA, 0x01}, broadcast_unknownMethod6MsgTag); err != nil {
		log.Fatalf("loading message tag for method UnknownMethod6: %v", err)
	}
	if err := proto.Unmarshal([]byte{0x08, 0xA8, 0xEA, 0x01, 0x10, 0xA9, 0xEA, 0x01}, broadcast_unknownMethod7MsgTag); err != nil {
		log.Fatalf("loading message tag for method UnknownMethod7: %v", err)
	}
	if err := proto.Unmarshal([]byte{0x08, 0xAA, 0xEA, 0x01, 0x10, 0xAB, 0xEA, 0x01}, broadcast_unknownMethod8MsgTag); err != nil {
		log.Fatalf("loading message tag for method UnknownMethod8: %v", err)
	}
}

type broadcastClientImpl struct {
	macAddr   string
	reqCh     chan frameWriterReq
	inFrameCh chan []byte
	errCh     chan error

	loginRequestCh                            chan internalUnaryReq[*LoginBroadcastServerRequest, *LoginBroadcastServerResponse]
	listBroadcastsRequestCh                   chan internalServerStreamingReq[*ListBroadcastsRequest, *ListBroadcastsResponse]
	enterBroadcastRequestCh                   chan internalUnaryReq[*EnterBroadcastRequest, *EnterBroadcastResponse]
	listenBroadcastSettingEventsRequestCh     chan internalServerStreamingReq[proto.Message, *BroadcastSettingEvent]
	listenBroadcastStateEventsRequestCh       chan internalServerStreamingReq[proto.Message, *BroadcastStateEvent]
	listenBroadcastMoveEventsRequestCh        chan internalServerStreamingReq[proto.Message, *BroadcastMoveEvent]
	listenBroadcastGameResultEventsRequestCh  chan internalServerStreamingReq[proto.Message, *BroadcastGameResultEvent]
	listenBroadcastTimeControlEventsRequestCh chan internalServerStreamingReq[proto.Message, *BroadcastTimeControlEvent]
	leaveBroadcastRequestCh                   chan internalUnaryReq[*LeaveBroadcastRequest, *LeaveBroadcastResponse]
	heartbeatRequestCh                        chan internalUnaryReq[*HeartbeatRequest, proto.Message]
	unknownMethod1RequestCh                   chan internalUnaryReq[*UnknownRequest1, proto.Message]
	unknownMethod2RequestCh                   chan internalUnaryReq[*UnknownRequest2, proto.Message]
	unknownMethod3RequestCh                   chan internalUnaryReq[*UnknownRequest3, proto.Message]
	unknownMethod4RequestCh                   chan internalUnaryReq[*UnknownRequest4, proto.Message]
	unknownMethod5RequestCh                   chan internalUnaryReq[*UnknownRequest5, proto.Message]
	unknownMethod6RequestCh                   chan internalUnaryReq[*UnknownRequest6, *UnknownResponse6]
	unknownMethod7RequestCh                   chan internalUnaryReq[*UnknownRequest7, *UnknownResponse7]
	unknownMethod8RequestCh                   chan internalUnaryReq[*UnknownRequest8, *UnknownResponse8]
}

func NewBroadcastClient(addr string, log *log.Logger) (BroadcastClient, error) {
	rwc, macAddr, err := newFrameChannel(addr)
	if err != nil {
		return nil, fmt.Errorf("connecting to %s: %v", addr, err)
	}

	frw := NewLenEncodedFraming(rwc)
	inFrameCh := make(chan []byte, 8)
	reqCh := make(chan frameWriterReq, 8)
	errCh := make(chan error, 8)

	go frameWriter(frw, reqCh, errCh, log)
	go frameReader(frw, inFrameCh, errCh)

	loginRequestCh := make(chan internalUnaryReq[*LoginBroadcastServerRequest, *LoginBroadcastServerResponse], 16)
	loginPending := &Queue[internalUnaryReq[*LoginBroadcastServerRequest, *LoginBroadcastServerResponse]]{}
	listBroadcastsRequestCh := make(chan internalServerStreamingReq[*ListBroadcastsRequest, *ListBroadcastsResponse], 16)
	listBroadcastsServerStreams := make(map[*serverStreamImpl[*ListBroadcastsResponse]]struct{})
	enterBroadcastRequestCh := make(chan internalUnaryReq[*EnterBroadcastRequest, *EnterBroadcastResponse], 16)
	enterBroadcastPending := &Queue[internalUnaryReq[*EnterBroadcastRequest, *EnterBroadcastResponse]]{}
	listenBroadcastSettingEventsRequestCh := make(chan internalServerStreamingReq[proto.Message, *BroadcastSettingEvent], 16)
	listenBroadcastSettingEventsServerStreams := make(map[*serverStreamImpl[*BroadcastSettingEvent]]struct{})
	listenBroadcastStateEventsRequestCh := make(chan internalServerStreamingReq[proto.Message, *BroadcastStateEvent], 16)
	listenBroadcastStateEventsServerStreams := make(map[*serverStreamImpl[*BroadcastStateEvent]]struct{})
	listenBroadcastMoveEventsRequestCh := make(chan internalServerStreamingReq[proto.Message, *BroadcastMoveEvent], 16)
	listenBroadcastMoveEventsServerStreams := make(map[*serverStreamImpl[*BroadcastMoveEvent]]struct{})
	listenBroadcastGameResultEventsRequestCh := make(chan internalServerStreamingReq[proto.Message, *BroadcastGameResultEvent], 16)
	listenBroadcastGameResultEventsServerStreams := make(map[*serverStreamImpl[*BroadcastGameResultEvent]]struct{})
	listenBroadcastTimeControlEventsRequestCh := make(chan internalServerStreamingReq[proto.Message, *BroadcastTimeControlEvent], 16)
	listenBroadcastTimeControlEventsServerStreams := make(map[*serverStreamImpl[*BroadcastTimeControlEvent]]struct{})
	leaveBroadcastRequestCh := make(chan internalUnaryReq[*LeaveBroadcastRequest, *LeaveBroadcastResponse], 16)
	leaveBroadcastPending := &Queue[internalUnaryReq[*LeaveBroadcastRequest, *LeaveBroadcastResponse]]{}
	heartbeatRequestCh := make(chan internalUnaryReq[*HeartbeatRequest, proto.Message], 16)
	unknownMethod1RequestCh := make(chan internalUnaryReq[*UnknownRequest1, proto.Message], 16)
	unknownMethod2RequestCh := make(chan internalUnaryReq[*UnknownRequest2, proto.Message], 16)
	unknownMethod3RequestCh := make(chan internalUnaryReq[*UnknownRequest3, proto.Message], 16)
	unknownMethod4RequestCh := make(chan internalUnaryReq[*UnknownRequest4, proto.Message], 16)
	unknownMethod5RequestCh := make(chan internalUnaryReq[*UnknownRequest5, proto.Message], 16)
	unknownMethod6RequestCh := make(chan internalUnaryReq[*UnknownRequest6, *UnknownResponse6], 16)
	unknownMethod6Pending := &Queue[internalUnaryReq[*UnknownRequest6, *UnknownResponse6]]{}
	unknownMethod7RequestCh := make(chan internalUnaryReq[*UnknownRequest7, *UnknownResponse7], 16)
	unknownMethod7Pending := &Queue[internalUnaryReq[*UnknownRequest7, *UnknownResponse7]]{}
	unknownMethod8RequestCh := make(chan internalUnaryReq[*UnknownRequest8, *UnknownResponse8], 16)
	unknownMethod8Pending := &Queue[internalUnaryReq[*UnknownRequest8, *UnknownResponse8]]{}
	go func() {
		var routerErr error
		for {
			select {
			case frame := <-inFrameCh:
				tag := uint32(binary.BigEndian.Uint16(frame[:2]))
				switch tag {
				case 0x0FA5:
					payload := frame[3 : len(frame)-4]
					resp := &LoginBroadcastServerResponse{}
					if err := proto.Unmarshal(payload, resp); err != nil {
						log.Printf("unmarshalling LoginBroadcastServerResponse response: %v\n\tframe: %04X%s", err, len(frame), hex.EncodeToString(frame))
						routerErr = err
						continue
					}
					if !loginPending.Empty() {
						rr := loginPending.Pop()
						go func() { rr.respCh <- resp }()
					} else {
						log.Printf("BroadcastClient: dead letter: %s { %v }", resp.ProtoReflect().Descriptor().Name(), resp)
					}
				case 0x0FCC:
					payload := frame[3 : len(frame)-4]
					resp := &ListBroadcastsResponse{}
					if err := proto.Unmarshal(payload, resp); err != nil {
						log.Printf("unmarshalling ListBroadcastsResponse response: %v\n\tframe: %04X%s", err, len(frame), hex.EncodeToString(frame))
						routerErr = err
						continue
					}
					if len(listBroadcastsServerStreams) == 0 {
						log.Printf("BroadcastClient: dead letter: %s { %v }", resp.ProtoReflect().Descriptor().Name(), resp)
					}
					var closedStreams []*serverStreamImpl[*ListBroadcastsResponse]
					for stream := range listBroadcastsServerStreams {
						select {
						case <-stream.closeCh:
							closedStreams = append(closedStreams, stream)
						case stream.respCh <- resp:
						}
					}
					for _, stream := range closedStreams {
						delete(listBroadcastsServerStreams, stream)
					}
				case 0x13C9:
					payload := frame[3 : len(frame)-4]
					resp := &EnterBroadcastResponse{}
					if err := proto.Unmarshal(payload, resp); err != nil {
						log.Printf("unmarshalling EnterBroadcastResponse response: %v\n\tframe: %04X%s", err, len(frame), hex.EncodeToString(frame))
						routerErr = err
						continue
					}
					if !enterBroadcastPending.Empty() {
						rr := enterBroadcastPending.Pop()
						go func() { rr.respCh <- resp }()
					} else {
						log.Printf("BroadcastClient: dead letter: %s { %v }", resp.ProtoReflect().Descriptor().Name(), resp)
					}
				case 0x1419:
					payload := frame[3 : len(frame)-4]
					resp := &BroadcastSettingEvent{}
					if err := proto.Unmarshal(payload, resp); err != nil {
						log.Printf("unmarshalling BroadcastSettingEvent response: %v\n\tframe: %04X%s", err, len(frame), hex.EncodeToString(frame))
						routerErr = err
						continue
					}
					if len(listenBroadcastSettingEventsServerStreams) == 0 {
						log.Printf("BroadcastClient: dead letter: %s { %v }", resp.ProtoReflect().Descriptor().Name(), resp)
					}
					var closedStreams []*serverStreamImpl[*BroadcastSettingEvent]
					for stream := range listenBroadcastSettingEventsServerStreams {
						select {
						case <-stream.closeCh:
							closedStreams = append(closedStreams, stream)
						case stream.respCh <- resp:
						}
					}
					for _, stream := range closedStreams {
						delete(listenBroadcastSettingEventsServerStreams, stream)
					}
				case 0x141B:
					payload := frame[3 : len(frame)-4]
					resp := &BroadcastStateEvent{}
					if err := proto.Unmarshal(payload, resp); err != nil {
						log.Printf("unmarshalling BroadcastStateEvent response: %v\n\tframe: %04X%s", err, len(frame), hex.EncodeToString(frame))
						routerErr = err
						continue
					}
					if len(listenBroadcastStateEventsServerStreams) == 0 {
						log.Printf("BroadcastClient: dead letter: %s { %v }", resp.ProtoReflect().Descriptor().Name(), resp)
					}
					var closedStreams []*serverStreamImpl[*BroadcastStateEvent]
					for stream := range listenBroadcastStateEventsServerStreams {
						select {
						case <-stream.closeCh:
							closedStreams = append(closedStreams, stream)
						case stream.respCh <- resp:
						}
					}
					for _, stream := range closedStreams {
						delete(listenBroadcastStateEventsServerStreams, stream)
					}
				case 0x1B5E:
					payload := frame[3 : len(frame)-4]
					resp := &BroadcastMoveEvent{}
					if err := proto.Unmarshal(payload, resp); err != nil {
						log.Printf("unmarshalling BroadcastMoveEvent response: %v\n\tframe: %04X%s", err, len(frame), hex.EncodeToString(frame))
						routerErr = err
						continue
					}
					if len(listenBroadcastMoveEventsServerStreams) == 0 {
						log.Printf("BroadcastClient: dead letter: %s { %v }", resp.ProtoReflect().Descriptor().Name(), resp)
					}
					var closedStreams []*serverStreamImpl[*BroadcastMoveEvent]
					for stream := range listenBroadcastMoveEventsServerStreams {
						select {
						case <-stream.closeCh:
							closedStreams = append(closedStreams, stream)
						case stream.respCh <- resp:
						}
					}
					for _, stream := range closedStreams {
						delete(listenBroadcastMoveEventsServerStreams, stream)
					}
				case 0x1C12:
					payload := frame[3 : len(frame)-4]
					resp := &BroadcastGameResultEvent{}
					if err := proto.Unmarshal(payload, resp); err != nil {
						log.Printf("unmarshalling BroadcastGameResultEvent response: %v\n\tframe: %04X%s", err, len(frame), hex.EncodeToString(frame))
						routerErr = err
						continue
					}
					if len(listenBroadcastGameResultEventsServerStreams) == 0 {
						log.Printf("BroadcastClient: dead letter: %s { %v }", resp.ProtoReflect().Descriptor().Name(), resp)
					}
					var closedStreams []*serverStreamImpl[*BroadcastGameResultEvent]
					for stream := range listenBroadcastGameResultEventsServerStreams {
						select {
						case <-stream.closeCh:
							closedStreams = append(closedStreams, stream)
						case stream.respCh <- resp:
						}
					}
					for _, stream := range closedStreams {
						delete(listenBroadcastGameResultEventsServerStreams, stream)
					}
				case 0x1C1E:
					payload := frame[3 : len(frame)-4]
					resp := &BroadcastTimeControlEvent{}
					if err := proto.Unmarshal(payload, resp); err != nil {
						log.Printf("unmarshalling BroadcastTimeControlEvent response: %v\n\tframe: %04X%s", err, len(frame), hex.EncodeToString(frame))
						routerErr = err
						continue
					}
					if len(listenBroadcastTimeControlEventsServerStreams) == 0 {
						log.Printf("BroadcastClient: dead letter: %s { %v }", resp.ProtoReflect().Descriptor().Name(), resp)
					}
					var closedStreams []*serverStreamImpl[*BroadcastTimeControlEvent]
					for stream := range listenBroadcastTimeControlEventsServerStreams {
						select {
						case <-stream.closeCh:
							closedStreams = append(closedStreams, stream)
						case stream.respCh <- resp:
						}
					}
					for _, stream := range closedStreams {
						delete(listenBroadcastTimeControlEventsServerStreams, stream)
					}
				case 0x13DD:
					payload := frame[3 : len(frame)-4]
					resp := &LeaveBroadcastResponse{}
					if err := proto.Unmarshal(payload, resp); err != nil {
						log.Printf("unmarshalling LeaveBroadcastResponse response: %v\n\tframe: %04X%s", err, len(frame), hex.EncodeToString(frame))
						routerErr = err
						continue
					}
					if !leaveBroadcastPending.Empty() {
						rr := leaveBroadcastPending.Pop()
						go func() { rr.respCh <- resp }()
					} else {
						log.Printf("BroadcastClient: dead letter: %s { %v }", resp.ProtoReflect().Descriptor().Name(), resp)
					}
				case 0x7527:
					payload := frame[3 : len(frame)-4]
					resp := &UnknownResponse6{}
					if err := proto.Unmarshal(payload, resp); err != nil {
						log.Printf("unmarshalling UnknownResponse6 response: %v\n\tframe: %04X%s", err, len(frame), hex.EncodeToString(frame))
						routerErr = err
						continue
					}
					if !unknownMethod6Pending.Empty() {
						rr := unknownMethod6Pending.Pop()
						go func() { rr.respCh <- resp }()
					} else {
						log.Printf("BroadcastClient: dead letter: %s { %v }", resp.ProtoReflect().Descriptor().Name(), resp)
					}
				case 0x7529:
					payload := frame[3 : len(frame)-4]
					resp := &UnknownResponse7{}
					if err := proto.Unmarshal(payload, resp); err != nil {
						log.Printf("unmarshalling UnknownResponse7 response: %v\n\tframe: %04X%s", err, len(frame), hex.EncodeToString(frame))
						routerErr = err
						continue
					}
					if !unknownMethod7Pending.Empty() {
						rr := unknownMethod7Pending.Pop()
						go func() { rr.respCh <- resp }()
					} else {
						log.Printf("BroadcastClient: dead letter: %s { %v }", resp.ProtoReflect().Descriptor().Name(), resp)
					}
				case 0x752B:
					payload := frame[3 : len(frame)-4]
					resp := &UnknownResponse8{}
					if err := proto.Unmarshal(payload, resp); err != nil {
						log.Printf("unmarshalling UnknownResponse8 response: %v\n\tframe: %04X%s", err, len(frame), hex.EncodeToString(frame))
						routerErr = err
						continue
					}
					if !unknownMethod8Pending.Empty() {
						rr := unknownMethod8Pending.Pop()
						go func() { rr.respCh <- resp }()
					} else {
						log.Printf("BroadcastClient: dead letter: %s { %v }", resp.ProtoReflect().Descriptor().Name(), resp)
					}
				default:
					log.Printf("BroadcastClient: dead letter: tag=%04X frame=%04X%s", tag, len(frame), hex.EncodeToString(frame))
				}
			case routerErr = <-errCh:
				log.Printf("BroadcastClient: router error set: %v", routerErr)
			case rr := <-loginRequestCh:
				if routerErr != nil {
					rr.errCh <- routerErr
				}
				reqCh <- frameWriterReq{
					msgTag: broadcast_loginMsgTag,
					msg:    rr.req,
				}
				loginPending.Push(rr)
			case rr := <-listBroadcastsRequestCh:
				if routerErr != nil {
					rr.errCh <- routerErr
				}
				reqCh <- frameWriterReq{
					msgTag: broadcast_listBroadcastsMsgTag,
					msg:    rr.req,
				}
				listBroadcastsServerStreams[rr.stream] = struct{}{}
			case rr := <-enterBroadcastRequestCh:
				if routerErr != nil {
					rr.errCh <- routerErr
				}
				reqCh <- frameWriterReq{
					msgTag: broadcast_enterBroadcastMsgTag,
					msg:    rr.req,
				}
				enterBroadcastPending.Push(rr)
			case rr := <-listenBroadcastSettingEventsRequestCh:
				if routerErr != nil {
					rr.errCh <- routerErr
				}
				listenBroadcastSettingEventsServerStreams[rr.stream] = struct{}{}
			case rr := <-listenBroadcastStateEventsRequestCh:
				if routerErr != nil {
					rr.errCh <- routerErr
				}
				listenBroadcastStateEventsServerStreams[rr.stream] = struct{}{}
			case rr := <-listenBroadcastMoveEventsRequestCh:
				if routerErr != nil {
					rr.errCh <- routerErr
				}
				listenBroadcastMoveEventsServerStreams[rr.stream] = struct{}{}
			case rr := <-listenBroadcastGameResultEventsRequestCh:
				if routerErr != nil {
					rr.errCh <- routerErr
				}
				listenBroadcastGameResultEventsServerStreams[rr.stream] = struct{}{}
			case rr := <-listenBroadcastTimeControlEventsRequestCh:
				if routerErr != nil {
					rr.errCh <- routerErr
				}
				listenBroadcastTimeControlEventsServerStreams[rr.stream] = struct{}{}
			case rr := <-leaveBroadcastRequestCh:
				if routerErr != nil {
					rr.errCh <- routerErr
				}
				reqCh <- frameWriterReq{
					msgTag: broadcast_leaveBroadcastMsgTag,
					msg:    rr.req,
				}
				leaveBroadcastPending.Push(rr)
			case rr := <-heartbeatRequestCh:
				if routerErr != nil {
					rr.errCh <- routerErr
				}
				reqCh <- frameWriterReq{
					msgTag: broadcast_heartbeatMsgTag,
					msg:    rr.req,
				}
			case rr := <-unknownMethod1RequestCh:
				if routerErr != nil {
					rr.errCh <- routerErr
				}
				reqCh <- frameWriterReq{
					msgTag: broadcast_unknownMethod1MsgTag,
					msg:    rr.req,
				}
			case rr := <-unknownMethod2RequestCh:
				if routerErr != nil {
					rr.errCh <- routerErr
				}
				reqCh <- frameWriterReq{
					msgTag: broadcast_unknownMethod2MsgTag,
					msg:    rr.req,
				}
			case rr := <-unknownMethod3RequestCh:
				if routerErr != nil {
					rr.errCh <- routerErr
				}
				reqCh <- frameWriterReq{
					msgTag: broadcast_unknownMethod3MsgTag,
					msg:    rr.req,
				}
			case rr := <-unknownMethod4RequestCh:
				if routerErr != nil {
					rr.errCh <- routerErr
				}
				reqCh <- frameWriterReq{
					msgTag: broadcast_unknownMethod4MsgTag,
					msg:    rr.req,
				}
			case rr := <-unknownMethod5RequestCh:
				if routerErr != nil {
					rr.errCh <- routerErr
				}
				reqCh <- frameWriterReq{
					msgTag: broadcast_unknownMethod5MsgTag,
					msg:    rr.req,
				}
			case rr := <-unknownMethod6RequestCh:
				if routerErr != nil {
					rr.errCh <- routerErr
				}
				reqCh <- frameWriterReq{
					msgTag: broadcast_unknownMethod6MsgTag,
					msg:    rr.req,
				}
				unknownMethod6Pending.Push(rr)
			case rr := <-unknownMethod7RequestCh:
				if routerErr != nil {
					rr.errCh <- routerErr
				}
				reqCh <- frameWriterReq{
					msgTag: broadcast_unknownMethod7MsgTag,
					msg:    rr.req,
				}
				unknownMethod7Pending.Push(rr)
			case rr := <-unknownMethod8RequestCh:
				if routerErr != nil {
					rr.errCh <- routerErr
				}
				reqCh <- frameWriterReq{
					msgTag: broadcast_unknownMethod8MsgTag,
					msg:    rr.req,
				}
				unknownMethod8Pending.Push(rr)
			}
		}
	}()

	return &broadcastClientImpl{
		macAddr:                                  macAddr,
		inFrameCh:                                inFrameCh,
		reqCh:                                    reqCh,
		errCh:                                    errCh,
		loginRequestCh:                           loginRequestCh,
		listBroadcastsRequestCh:                  listBroadcastsRequestCh,
		enterBroadcastRequestCh:                  enterBroadcastRequestCh,
		listenBroadcastSettingEventsRequestCh:    listenBroadcastSettingEventsRequestCh,
		listenBroadcastStateEventsRequestCh:      listenBroadcastStateEventsRequestCh,
		listenBroadcastMoveEventsRequestCh:       listenBroadcastMoveEventsRequestCh,
		listenBroadcastGameResultEventsRequestCh: listenBroadcastGameResultEventsRequestCh,
		listenBroadcastTimeControlEventsRequestCh: listenBroadcastTimeControlEventsRequestCh,
		leaveBroadcastRequestCh:                   leaveBroadcastRequestCh,
		heartbeatRequestCh:                        heartbeatRequestCh,
		unknownMethod1RequestCh:                   unknownMethod1RequestCh,
		unknownMethod2RequestCh:                   unknownMethod2RequestCh,
		unknownMethod3RequestCh:                   unknownMethod3RequestCh,
		unknownMethod4RequestCh:                   unknownMethod4RequestCh,
		unknownMethod5RequestCh:                   unknownMethod5RequestCh,
		unknownMethod6RequestCh:                   unknownMethod6RequestCh,
		unknownMethod7RequestCh:                   unknownMethod7RequestCh,
		unknownMethod8RequestCh:                   unknownMethod8RequestCh,
	}, nil
}

func (c *broadcastClientImpl) Login(_ context.Context, req *LoginBroadcastServerRequest) (*LoginBroadcastServerResponse, error) {
	rr := internalUnaryReq[*LoginBroadcastServerRequest, *LoginBroadcastServerResponse]{
		req:    req,
		respCh: make(chan *LoginBroadcastServerResponse, 1),
		errCh:  make(chan error),
	}
	defer close(rr.respCh)
	defer close(rr.errCh)
	c.loginRequestCh <- rr
	select {
	case resp := <-rr.respCh:
		return resp, nil
	case err := <-rr.errCh:
		return nil, err
	}
}

func (c *broadcastClientImpl) ListBroadcasts(_ context.Context, req *ListBroadcastsRequest) (ServerStream[*ListBroadcastsResponse], error) {
	stream := &serverStreamImpl[*ListBroadcastsResponse]{
		respCh:  make(chan *ListBroadcastsResponse, 8),
		closeCh: make(chan struct{}),
	}
	rr := internalServerStreamingReq[*ListBroadcastsRequest, *ListBroadcastsResponse]{
		req:    req,
		stream: stream,
		errCh:  make(chan error),
	}
	defer close(rr.errCh)
	c.listBroadcastsRequestCh <- rr
	select {
	case err := <-rr.errCh:
		return nil, err
	default:
		return rr.stream, nil
	}
}

func (c *broadcastClientImpl) EnterBroadcast(_ context.Context, req *EnterBroadcastRequest) (*EnterBroadcastResponse, error) {
	rr := internalUnaryReq[*EnterBroadcastRequest, *EnterBroadcastResponse]{
		req:    req,
		respCh: make(chan *EnterBroadcastResponse, 1),
		errCh:  make(chan error),
	}
	defer close(rr.respCh)
	defer close(rr.errCh)
	c.enterBroadcastRequestCh <- rr
	select {
	case resp := <-rr.respCh:
		return resp, nil
	case err := <-rr.errCh:
		return nil, err
	}
}

func (c *broadcastClientImpl) ListenBroadcastSettingEvents(_ context.Context) (ServerStream[*BroadcastSettingEvent], error) {
	stream := &serverStreamImpl[*BroadcastSettingEvent]{
		respCh:  make(chan *BroadcastSettingEvent, 8),
		closeCh: make(chan struct{}),
	}
	rr := internalServerStreamingReq[proto.Message, *BroadcastSettingEvent]{
		stream: stream,
		errCh:  make(chan error),
	}
	defer close(rr.errCh)
	c.listenBroadcastSettingEventsRequestCh <- rr
	select {
	case err := <-rr.errCh:
		return nil, err
	default:
		return rr.stream, nil
	}
}

func (c *broadcastClientImpl) ListenBroadcastStateEvents(_ context.Context) (ServerStream[*BroadcastStateEvent], error) {
	stream := &serverStreamImpl[*BroadcastStateEvent]{
		respCh:  make(chan *BroadcastStateEvent, 8),
		closeCh: make(chan struct{}),
	}
	rr := internalServerStreamingReq[proto.Message, *BroadcastStateEvent]{
		stream: stream,
		errCh:  make(chan error),
	}
	defer close(rr.errCh)
	c.listenBroadcastStateEventsRequestCh <- rr
	select {
	case err := <-rr.errCh:
		return nil, err
	default:
		return rr.stream, nil
	}
}

func (c *broadcastClientImpl) ListenBroadcastMoveEvents(_ context.Context) (ServerStream[*BroadcastMoveEvent], error) {
	stream := &serverStreamImpl[*BroadcastMoveEvent]{
		respCh:  make(chan *BroadcastMoveEvent, 8),
		closeCh: make(chan struct{}),
	}
	rr := internalServerStreamingReq[proto.Message, *BroadcastMoveEvent]{
		stream: stream,
		errCh:  make(chan error),
	}
	defer close(rr.errCh)
	c.listenBroadcastMoveEventsRequestCh <- rr
	select {
	case err := <-rr.errCh:
		return nil, err
	default:
		return rr.stream, nil
	}
}

func (c *broadcastClientImpl) ListenBroadcastGameResultEvents(_ context.Context) (ServerStream[*BroadcastGameResultEvent], error) {
	stream := &serverStreamImpl[*BroadcastGameResultEvent]{
		respCh:  make(chan *BroadcastGameResultEvent, 8),
		closeCh: make(chan struct{}),
	}
	rr := internalServerStreamingReq[proto.Message, *BroadcastGameResultEvent]{
		stream: stream,
		errCh:  make(chan error),
	}
	defer close(rr.errCh)
	c.listenBroadcastGameResultEventsRequestCh <- rr
	select {
	case err := <-rr.errCh:
		return nil, err
	default:
		return rr.stream, nil
	}
}

func (c *broadcastClientImpl) ListenBroadcastTimeControlEvents(_ context.Context) (ServerStream[*BroadcastTimeControlEvent], error) {
	stream := &serverStreamImpl[*BroadcastTimeControlEvent]{
		respCh:  make(chan *BroadcastTimeControlEvent, 8),
		closeCh: make(chan struct{}),
	}
	rr := internalServerStreamingReq[proto.Message, *BroadcastTimeControlEvent]{
		stream: stream,
		errCh:  make(chan error),
	}
	defer close(rr.errCh)
	c.listenBroadcastTimeControlEventsRequestCh <- rr
	select {
	case err := <-rr.errCh:
		return nil, err
	default:
		return rr.stream, nil
	}
}

func (c *broadcastClientImpl) LeaveBroadcast(_ context.Context, req *LeaveBroadcastRequest) (*LeaveBroadcastResponse, error) {
	rr := internalUnaryReq[*LeaveBroadcastRequest, *LeaveBroadcastResponse]{
		req:    req,
		respCh: make(chan *LeaveBroadcastResponse, 1),
		errCh:  make(chan error),
	}
	defer close(rr.respCh)
	defer close(rr.errCh)
	c.leaveBroadcastRequestCh <- rr
	select {
	case resp := <-rr.respCh:
		return resp, nil
	case err := <-rr.errCh:
		return nil, err
	}
}

func (c *broadcastClientImpl) Heartbeat(_ context.Context, req *HeartbeatRequest) error {
	rr := internalUnaryReq[*HeartbeatRequest, proto.Message]{
		req:   req,
		errCh: make(chan error),
	}
	defer close(rr.errCh)
	c.heartbeatRequestCh <- rr
	select {
	case err := <-rr.errCh:
		return err
	default:
		return nil
	}
}

func (c *broadcastClientImpl) UnknownMethod1(_ context.Context, req *UnknownRequest1) error {
	rr := internalUnaryReq[*UnknownRequest1, proto.Message]{
		req:   req,
		errCh: make(chan error),
	}
	defer close(rr.errCh)
	c.unknownMethod1RequestCh <- rr
	select {
	case err := <-rr.errCh:
		return err
	default:
		return nil
	}
}

func (c *broadcastClientImpl) UnknownMethod2(_ context.Context, req *UnknownRequest2) error {
	rr := internalUnaryReq[*UnknownRequest2, proto.Message]{
		req:   req,
		errCh: make(chan error),
	}
	defer close(rr.errCh)
	c.unknownMethod2RequestCh <- rr
	select {
	case err := <-rr.errCh:
		return err
	default:
		return nil
	}
}

func (c *broadcastClientImpl) UnknownMethod3(_ context.Context, req *UnknownRequest3) error {
	rr := internalUnaryReq[*UnknownRequest3, proto.Message]{
		req:   req,
		errCh: make(chan error),
	}
	defer close(rr.errCh)
	c.unknownMethod3RequestCh <- rr
	select {
	case err := <-rr.errCh:
		return err
	default:
		return nil
	}
}

func (c *broadcastClientImpl) UnknownMethod4(_ context.Context, req *UnknownRequest4) error {
	rr := internalUnaryReq[*UnknownRequest4, proto.Message]{
		req:   req,
		errCh: make(chan error),
	}
	defer close(rr.errCh)
	c.unknownMethod4RequestCh <- rr
	select {
	case err := <-rr.errCh:
		return err
	default:
		return nil
	}
}

func (c *broadcastClientImpl) UnknownMethod5(_ context.Context, req *UnknownRequest5) error {
	rr := internalUnaryReq[*UnknownRequest5, proto.Message]{
		req:   req,
		errCh: make(chan error),
	}
	defer close(rr.errCh)
	c.unknownMethod5RequestCh <- rr
	select {
	case err := <-rr.errCh:
		return err
	default:
		return nil
	}
}

func (c *broadcastClientImpl) UnknownMethod6(_ context.Context, req *UnknownRequest6) (*UnknownResponse6, error) {
	rr := internalUnaryReq[*UnknownRequest6, *UnknownResponse6]{
		req:    req,
		respCh: make(chan *UnknownResponse6, 1),
		errCh:  make(chan error),
	}
	defer close(rr.respCh)
	defer close(rr.errCh)
	c.unknownMethod6RequestCh <- rr
	select {
	case resp := <-rr.respCh:
		return resp, nil
	case err := <-rr.errCh:
		return nil, err
	}
}

func (c *broadcastClientImpl) UnknownMethod7(_ context.Context, req *UnknownRequest7) (*UnknownResponse7, error) {
	rr := internalUnaryReq[*UnknownRequest7, *UnknownResponse7]{
		req:    req,
		respCh: make(chan *UnknownResponse7, 1),
		errCh:  make(chan error),
	}
	defer close(rr.respCh)
	defer close(rr.errCh)
	c.unknownMethod7RequestCh <- rr
	select {
	case resp := <-rr.respCh:
		return resp, nil
	case err := <-rr.errCh:
		return nil, err
	}
}

func (c *broadcastClientImpl) UnknownMethod8(_ context.Context, req *UnknownRequest8) (*UnknownResponse8, error) {
	rr := internalUnaryReq[*UnknownRequest8, *UnknownResponse8]{
		req:    req,
		respCh: make(chan *UnknownResponse8, 1),
		errCh:  make(chan error),
	}
	defer close(rr.respCh)
	defer close(rr.errCh)
	c.unknownMethod8RequestCh <- rr
	select {
	case resp := <-rr.respCh:
		return resp, nil
	case err := <-rr.errCh:
		return nil, err
	}
}

func (c *broadcastClientImpl) MACAddress() string { return c.macAddr }
