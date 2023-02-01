package main

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/websocket"
	"google.golang.org/protobuf/encoding/prototext"
	"google.golang.org/protobuf/proto"

	pb "github.com/ale64bit/openfoxwq/proto"

	_ "net/http/pprof"
)

var (
	port = flag.Int("port", 8999, "")
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  4 << 10,
	WriteBufferSize: 4 << 10,
}

func roomIdStr(id *pb.RoomId) string {
	return fmt.Sprintf("%d|%d|%d|%d", id.GetId_1(), id.GetId_2(), id.GetId_3(), id.GetId_4())
}

func broadcastEventHandler(ctx context.Context, bcClient pb.BroadcastClient, respCh chan<- *pb.WsResponse) error {
	broadcastSettingStream, err := bcClient.ListenBroadcastSettingEvents(ctx)
	if err != nil {
		return fmt.Errorf("listening to broadcast config events: %v", err)
	}
	broadcastStateStream, err := bcClient.ListenBroadcastStateEvents(ctx)
	if err != nil {
		return fmt.Errorf("listening to broadcast state events: %v", err)
	}
	broadcastMoveStream, err := bcClient.ListenBroadcastMoveEvents(ctx)
	if err != nil {
		return fmt.Errorf("listening to broadcast move events: %v", err)
	}
	broadcastGameResultStream, err := bcClient.ListenBroadcastGameResultEvents(ctx)
	if err != nil {
		return fmt.Errorf("listening to broadcast game result events: %v", err)
	}
	broadcastTimeControlStream, err := bcClient.ListenBroadcastTimeControlEvents(ctx)
	if err != nil {
		return fmt.Errorf("listening to broadcast time control events: %v", err)
	}

	log.Printf("listening to broadcast events...")
	for {
		select {
		case evt := <-broadcastSettingStream.Ch():
			// log.Printf("broadcastroom config event: %v", evt)
			respCh <- &pb.WsResponse{Resp: &pb.WsResponse_BroadcastSettingEvent{BroadcastSettingEvent: evt}}
		case evt := <-broadcastStateStream.Ch():
			// log.Printf("broadcast state event: %v", evt)
			respCh <- &pb.WsResponse{Resp: &pb.WsResponse_BroadcastStateEvent{BroadcastStateEvent: evt}}
		case evt := <-broadcastMoveStream.Ch():
			// log.Printf("broadcast move event: %v", evt)
			respCh <- &pb.WsResponse{Resp: &pb.WsResponse_BroadcastMoveEvent{BroadcastMoveEvent: evt}}
		case evt := <-broadcastGameResultStream.Ch():
			// log.Printf("broadcast game result event: %v", evt)
			respCh <- &pb.WsResponse{Resp: &pb.WsResponse_BroadcastGameResultEvent{BroadcastGameResultEvent: evt}}
		case evt := <-broadcastTimeControlStream.Ch():
			// log.Printf("broadcast time control event: %v", evt)
			respCh <- &pb.WsResponse{Resp: &pb.WsResponse_BroadcastTimeControlEvent{BroadcastTimeControlEvent: evt}}
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

func playerEventHandler(ctx context.Context, playClient pb.PlayClient, respCh chan<- *pb.WsResponse) error {
	playerOnlineCountStream, err := playClient.ListenPlayerOnlineCountEvents(ctx, &pb.MessageHeader{})
	if err != nil {
		return fmt.Errorf("listening to player online count events: %v", err)
	}
	playerOnlineStream, err := playClient.ListenPlayerOnlineEvents(ctx, &pb.MessageHeader{})
	if err != nil {
		return fmt.Errorf("listening to player online events: %v", err)
	}
	playerOfflineStream, err := playClient.ListenPlayerOfflineEvents(ctx, &pb.MessageHeader{})
	if err != nil {
		return fmt.Errorf("listening to player offline events: %v", err)
	}
	playerStateStream, err := playClient.ListenPlayerStateEvents(ctx, &pb.MessageHeader{})
	if err != nil {
		return fmt.Errorf("listening to player state events: %v", err)
	}

	log.Printf("listening to player events...")
	for {
		select {
		case evt := <-playerOnlineCountStream.Ch():
			// log.Printf("player online count event: %v", evt)
			respCh <- &pb.WsResponse{Resp: &pb.WsResponse_PlayerOnlineCountEvent{PlayerOnlineCountEvent: evt.Resp}}
		case evt := <-playerOnlineStream.Ch():
			// log.Printf("player online event: %v", evt)
			respCh <- &pb.WsResponse{Resp: &pb.WsResponse_PlayerOnlineEvent{PlayerOnlineEvent: evt.Resp}}
		case evt := <-playerOfflineStream.Ch():
			// log.Printf("player offline event: %v", evt)
			respCh <- &pb.WsResponse{Resp: &pb.WsResponse_PlayerOfflineEvent{PlayerOfflineEvent: evt.Resp}}
		case evt := <-playerStateStream.Ch():
			// log.Printf("player state event: %v", evt)
			respCh <- &pb.WsResponse{Resp: &pb.WsResponse_PlayerStateEvent{PlayerStateEvent: evt.Resp}}
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

func gameEventHandler(ctx context.Context, playerID int64, playClient pb.PlayClient, respCh chan<- *pb.WsResponse) error {
	// Automatch event streams
	automatchFoundStream, err := playClient.ListenAutomatchFoundEvents(ctx, &pb.MessageHeader{})
	if err != nil {
		return fmt.Errorf("listening to automatch found events: %v", err)
	}

	// Game event streams
	matchStartStream, err := playClient.ListenMatchStartEvents(ctx, &pb.MessageHeader{})
	if err != nil {
		return fmt.Errorf("listening to match start events: %v", err)
	}
	nextMoveStream, err := playClient.ListenNextMoveEvents(ctx, &pb.MessageHeader{})
	if err != nil {
		return fmt.Errorf("listening to next move events: %v", err)
	}
	passStream, err := playClient.ListenPassEvents(ctx, &pb.MessageHeader{})
	if err != nil {
		return fmt.Errorf("listening to pass events: %v", err)
	}
	countdownStream, err := playClient.ListenCountdownEvents(ctx, &pb.MessageHeader{})
	if err != nil {
		return fmt.Errorf("listening to countdown events: %v", err)
	}
	resumeCountdownStream, err := playClient.ListenResumeCountdownEvents(ctx, &pb.MessageHeader{})
	if err != nil {
		return fmt.Errorf("listening to resume countdown events: %v", err)
	}
	countingDecisionStream, err := playClient.ListenCountingDecisions(ctx, &pb.MessageHeader{})
	if err != nil {
		return fmt.Errorf("listening to counting decision events: %v", err)
	}
	countingResultStream, err := playClient.ListenCountingResultEvents(ctx, &pb.MessageHeader{})
	if err != nil {
		return fmt.Errorf("listening to counting result events: %v", err)
	}
	gameResultStream, err := playClient.ListenGameResultEvents(ctx, &pb.MessageHeader{})
	if err != nil {
		return fmt.Errorf("listening to game result events: %v", err)
	}

	log.Printf("listening to game events...")
	for {
		select {
		case evt := <-automatchFoundStream.Ch():
			log.Printf("automatch found event: %v", evt.Resp)
			acceptHdr := &pb.MessageHeader{
				PlayerId:       proto.Int64(playerID),
				UnknownField_8: evt.Resp.RoomId_2,
			}
			acceptResp, err := playClient.AcceptMatch(ctx, acceptHdr, &pb.AcceptMatchRequest{})
			if err != nil {
				log.Printf("accepting match: %v", err)
				continue
			}
			if acceptResp.Resp.GetErrorCode() != 0 {
				log.Printf("accepting match: error code = %d", acceptResp.Resp.GetErrorCode())
				continue
			}
			log.Printf("match %d|%d|%d|_ accepted", evt.Resp.GetRoomId_1(), evt.Resp.GetRoomId_2(), evt.Resp.GetRoomId_3())
		case evt := <-matchStartStream.Ch():
			log.Printf("match start event: %v", evt.Resp)
			respCh <- &pb.WsResponse{Resp: &pb.WsResponse_MatchStartEvent{MatchStartEvent: evt.Resp}}
		case evt := <-nextMoveStream.Ch():
			log.Printf("next move event: %v", evt.Resp)
			respCh <- &pb.WsResponse{Resp: &pb.WsResponse_NextMove{
				NextMove: &pb.WsNextMoveResponse{
					RoomId_2: evt.Header.UnknownField_8,
					Event:    evt.Resp,
				},
			}}
		case evt := <-passStream.Ch():
			log.Printf("pass event: %v", evt)
		case evt := <-countdownStream.Ch():
			log.Printf("countdown event: %v", evt)
		case evt := <-resumeCountdownStream.Ch():
			log.Printf("resume countdown event: %v", evt)
		case evt := <-countingDecisionStream.Ch():
			log.Printf("counting decision: %v", evt)
		case evt := <-countingResultStream.Ch():
			log.Printf("counting result event: %v", evt)
		case evt := <-gameResultStream.Ch():
			log.Printf("game result event: %v", evt)
			respCh <- &pb.WsResponse{Resp: &pb.WsResponse_GameResult{
				GameResult: &pb.WsGameResultResponse{
					RoomId_2: evt.Header.UnknownField_8,
					Result:   evt.Resp,
				},
			}}
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	log.Printf("new connection")
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("upgrading connection: %v", err)
		return
	}
	defer conn.Close()

	ctx := context.Background()

	// Websocket writes cannot be concurrent
	wsRespCh := make(chan *pb.WsResponse, 16)

	go func() {
		for msg := range wsRespCh {
			b, err := proto.Marshal(msg)
			if err != nil {
				log.Fatalf("marshalling proto: %v", err)
			}
			if err := conn.WriteMessage(websocket.BinaryMessage, b); err != nil {
				log.Fatalf("writing to websocket: %v", err)
			}
		}
	}()

	// Create log files
	foxClientLogFile, err := os.Create("log/FoxClient.log")
	if err != nil {
		log.Fatalf("creating FoxClient log file: %v", err)
	}
	defer foxClientLogFile.Close()
	navClientLogFile, err := os.Create("log/NavigationClient.log")
	if err != nil {
		log.Fatalf("creating NavigationClient log file: %v", err)
	}
	defer navClientLogFile.Close()
	bcClientLogFile, err := os.Create("log/BroadcastClient.log")
	if err != nil {
		log.Fatalf("creating BroadcastClient log file: %v", err)
	}
	defer bcClientLogFile.Close()
	playClientLogFile, err := os.Create("log/PlayClient.log")
	if err != nil {
		log.Fatalf("creating PlayClient log file: %v", err)
	}
	defer playClientLogFile.Close()

	// Get navigation and server info
	log.Printf("creating fox client...")
	foxClient, err := pb.NewFoxClient("hknavi.foxwq.com:8002", log.New(foxClientLogFile, "", log.LstdFlags))
	if err != nil {
		log.Printf("creating fox client: %v", err)
		return
	}

	macAddr := foxClient.MACAddress()

	log.Printf("getting navigation info...")
	navInfo, err := foxClient.GetNavInfo(ctx, &pb.GetNavInfoRequest{
		UnknownField_1: proto.Int64(0),
		UnknownField_3: proto.Int64(0),
		UnknownField_4: proto.Int64(0),
		UnknownField_5: proto.Int64(0),
		UnknownField_6: proto.Int64(0),
		UnknownField_7: proto.Int64(1),
	})
	if err != nil {
		log.Printf("getting navigation info: %v", err)
		return
	}

	log.Printf("creating navigation client...")
	navClient, err := pb.NewNavigationClient(fmt.Sprintf("%s:%d", navInfo.GetNavHost(), navInfo.GetNavPort()), log.New(navClientLogFile, "", log.LstdFlags))
	if err != nil {
		log.Printf("creating navigation client: %v", err)
		return
	}

	log.Printf("getting server list...")
	serverList, err := navClient.ListServers(ctx, &pb.ListServersRequest{
		PlayerId: proto.Int64(0),
	})
	if err != nil {
		log.Printf("getting server list: %v", err)
		return
	}

	// Send initial navigation and server info
	wsRespCh <- &pb.WsResponse{Resp: &pb.WsResponse_NavInfo{NavInfo: navInfo}}
	wsRespCh <- &pb.WsResponse{Resp: &pb.WsResponse_ServerInfo{ServerInfo: serverList.GetServerInfo()}}

	// Session state
	serverInfo := serverList.GetServerInfo()
	var loginInfo *pb.LoginResponse
	var bcClient pb.BroadcastClient
	var playClient pb.PlayClient

	// Main loop
	for {
		msgType, data, err := conn.ReadMessage()
		if err != nil {
			log.Fatalf("reading message: %v", err)
		}
		if msgType != websocket.BinaryMessage {
			continue
		}
		msg := &pb.WsRequest{}
		if err := proto.Unmarshal(data, msg); err != nil {
			log.Fatalf("unmarshalling message: %v", err)
		}

		log.Printf("[ws] req: %v", msg)
		switch req := msg.Req.(type) {
		case *pb.WsRequest_Login:
			passwordHash := md5.Sum([]byte(req.Login.GetPassword()))
			loginResp, err := navClient.Login(ctx, &pb.LoginRequest{
				User:           proto.String(req.Login.GetUsername()),
				App:            proto.String("foxwq"),
				PasswordHash:   proto.String(hex.EncodeToString(passwordHash[:])),
				ClientVersion:  navInfo.GetVersionInfo().Version1,
				MacAddress:     proto.String(macAddr),
				UnknownField_4: proto.Int64(8),
			})
			if err != nil {
				log.Printf("logging in: %v", err)
				return
			}

			wsRespCh <- &pb.WsResponse{Resp: &pb.WsResponse_Login{Login: loginResp}}

			loginInfo = loginResp
			serverList, err := navClient.ListServers(ctx, &pb.ListServersRequest{
				PlayerId: proto.Int64(loginResp.GetPlayerId()),
			})
			if err != nil {
				log.Printf("getting updated server list: %v", err)
				return
			}
			serverInfo = serverList.GetServerInfo()

		case *pb.WsRequest_GetInitData:
			// Login to broadcast service and list rooms
			{
				log.Printf("creating broadcast client...")
				if bcClient, err = pb.NewBroadcastClient(fmt.Sprintf("%s:%d", serverInfo.GetBroadcastHost(), serverInfo.GetBroadcastPort()), log.New(bcClientLogFile, "", log.LstdFlags)); err != nil {
					log.Printf("creating broadcast client: %v", err)
					return
				}

				if resp, err := bcClient.Login(ctx, &pb.LoginBroadcastServerRequest{
					PlayerId:       loginInfo.PlayerId,
					Version:        navInfo.GetVersionInfo().Version1,
					Token:          loginInfo.Token1,
					UnknownField_3: proto.Int64(0),
					MacAddress:     proto.String(macAddr),
					UnknownField_6: proto.Int64(1),
					UnknownField_7: proto.Int64(1),
				}); err != nil || resp.GetPlayerId() == 0 {
					log.Printf("logging into broadcast service: %v", err)
					return
				}

				// Start heartbeat
				go func() {
					for range time.Tick(20 * time.Second) {
						if err := bcClient.Heartbeat(ctx, &pb.HeartbeatRequest{}); err != nil {
							log.Printf("sending heartbeat to broadcast: %v", err)
						}
					}
				}()

				log.Printf("getting broadcast list...")
				broadcastListStream, err := bcClient.ListBroadcasts(ctx, &pb.ListBroadcastsRequest{
					PlayerId:       loginInfo.PlayerId,
					UnknownField_2: proto.Int64(1),
					UnknownField_3: proto.Int64(0),
				})
				if err != nil {
					log.Printf("getting broadcast list: %v", err)
					return
				}

				for {
					resp, err := broadcastListStream.Recv()
					if err != nil {
						log.Printf("receiving broadcast list stream response: %v", err)
						return
					}

					log.Printf("broadcast list resp %d/%d: %d", resp.GetPageIndex()+1, resp.GetPageCount(), len(resp.GetBroadcasts()))

					wsRespCh <- &pb.WsResponse{Resp: &pb.WsResponse_ListBroadcasts{ListBroadcasts: resp}}

					if resp.GetPageIndex()+1 == resp.GetPageCount() {
						broadcastListStream.Close()
						break
					}
				}

				// Listen to broadcast events
				go func() {
					if err := broadcastEventHandler(ctx, bcClient, wsRespCh); err != nil {
						log.Fatalf("handling broadcast events: %v", err)
					}
				}()
			}

			// Login to game service and list players
			{
				log.Printf("creating play client...")
				if playClient, err = pb.NewPlayClient(fmt.Sprintf("%s:%d", serverInfo.GetPlayHost(), serverInfo.GetPlayPort()), log.New(playClientLogFile, "", log.LstdFlags)); err != nil {
					log.Printf("creating play client: %v", err)
					return
				}

				log.Printf("logging into play service...")
				if resp, err := pb.DiscardHeader(playClient.Login(ctx, &pb.MessageHeader{PlayerId: loginInfo.PlayerId}, &pb.LoginPlayServerRequest{Token: loginInfo.Token1})); err != nil || resp.GetUnknownField_1() != 0 {
					log.Printf("logging in to play client: %v", err)
					return
				}

				// Listen to player events
				go func() {
					if err := playerEventHandler(ctx, playClient, wsRespCh); err != nil {
						log.Fatalf("handling player events: %v", err)
					}
				}()

				// Listen to game events
				go func() {
					if err := gameEventHandler(ctx, loginInfo.GetPlayerId(), playClient, wsRespCh); err != nil {
						log.Fatalf("handling game events: %v", err)
					}
				}()

				log.Printf("listing players...")
				playerListStream, err := playClient.ListPlayers(ctx, &pb.MessageHeader{PlayerId: loginInfo.PlayerId}, &pb.ListPlayersRequest{UnknownField_1: proto.Int64(1)})
				if err != nil {
					log.Printf("listing players: %v", err)
					return
				}

				for {
					resp, err := pb.DiscardHeader(playerListStream.Recv())
					if err != nil {
						log.Printf("receiving player list response: %v", err)
						return
					}
					log.Printf("player list %d/%d: %d", resp.GetPageIndex()+1, resp.GetPageCount(), len(resp.GetPlayers()))

					wsRespCh <- &pb.WsResponse{Resp: &pb.WsResponse_ListPlayers{ListPlayers: resp}}

					if resp.GetPageIndex()+1 == resp.GetPageCount() {
						log.Printf("online count: %d", resp.GetOnlineCount())
						playerListStream.Close()
						break
					}
				}

				// Start time sync
				log.Printf("starting time sync...")
				go func() {
					for range time.Tick(20 * time.Second) {
						if resp, err := pb.DiscardHeader(playClient.SyncTime(ctx, &pb.MessageHeader{PlayerId: loginInfo.PlayerId}, &pb.SyncTimeRequest{UnixTs: proto.Int64(time.Now().Unix())})); err != nil {
							log.Printf("sending sync time to play: %v", err)
						} else {
							log.Printf("time sync resp: %v", resp)
						}
					}
				}()

				// Start player sync
				log.Printf("starting player info sync...")
				go func() {
					req := &pb.SyncPlayersRequest{
						UnknownField_1: proto.Int64(1),
					}
					for range time.Tick(10 * time.Second) {
						if resp, err := pb.DiscardHeader(playClient.SyncPlayers(ctx, &pb.MessageHeader{PlayerId: loginInfo.PlayerId}, req)); err != nil {
							log.Printf("sending sync players to play: %v", err)
						} else {
							log.Printf("sync players resp: %v", resp)
						}
					}
				}()

				// Start unknown request 1 stream
				go func() {
					req := &pb.UnknownPlayRequest1{
						UnknownField_1: &pb.UnknownPlayRequest1_UnknownPlayRequest1Nested1{
							MacAddress: proto.String(macAddr),
						},
						UnknownField_2: proto.Int64(2),
					}
					for range time.Tick(1 * time.Minute) {
						if resp, err := pb.DiscardHeader(playClient.Unknown1(ctx, &pb.MessageHeader{PlayerId: loginInfo.PlayerId}, req)); err != nil {
							log.Printf("sending unknown req 1 to play: %v", err)
						} else {
							log.Printf("unknown resp 1: %v", resp)
						}
					}
				}()

				// Get automatch population
				log.Printf("getting automatch stats...")
				go func() {
					{
						resp, err := pb.DiscardHeader(playClient.GetAutomatchStats(ctx, &pb.MessageHeader{PlayerId: loginInfo.PlayerId}, &pb.GetAutomatchStatsRequest{}))
						if err != nil {
							log.Printf("getting automatch stats: %v", err)
							return
						}
						wsRespCh <- &pb.WsResponse{Resp: &pb.WsResponse_GetAutomatchStats{GetAutomatchStats: resp}}
						log.Printf("automatch stats: %v", resp.GetPopulation())
					}
				}()
			}

		case *pb.WsRequest_EnterRoom:
			switch room := req.EnterRoom.Room.(type) {
			case *pb.WsEnterRoomRequest_BroadcastId:
				resp, err := bcClient.EnterBroadcast(ctx, &pb.EnterBroadcastRequest{
					PlayerId:    loginInfo.PlayerId,
					BroadcastId: proto.Int64(int64(room.BroadcastId)),
				})
				if err != nil {
					log.Printf("entering broadcast room: %v", err)
				}
				log.Printf("enter broadcast room resp: %v", resp)
			case *pb.WsEnterRoomRequest_RoomId:
				enterRoomHdr := &pb.MessageHeader{
					PlayerId: loginInfo.PlayerId,
				}
				resp, err := pb.DiscardHeader(playClient.EnterRoom(ctx, enterRoomHdr, &pb.EnterRoomRequest{Id: room.RoomId}))
				if err != nil {
					log.Printf("entering match room %s: %v", roomIdStr(room.RoomId), err)
					continue
				}
				if resp.GetErrorCode() != 0 {
					log.Printf("entering match room %s: error code = %d", roomIdStr(room.RoomId), resp.GetErrorCode())
					continue
				}
				log.Printf("entered match room %s: heartbeat_info = %v", roomIdStr(room.RoomId), resp.GetHeartbeatInfo())
				wsRespCh <- &pb.WsResponse{Resp: &pb.WsResponse_EnterRoom{EnterRoom: resp}}
			}

		case *pb.WsRequest_LeaveRoom:
			switch room := req.LeaveRoom.Room.(type) {
			case *pb.WsLeaveRoomRequest_BroadcastId:
				resp, err := bcClient.LeaveBroadcast(ctx, &pb.LeaveBroadcastRequest{
					PlayerId:    loginInfo.PlayerId,
					BroadcastId: proto.Int64(int64(room.BroadcastId)),
				})
				if err != nil {
					log.Printf("leaving broadcast: %v", err)
				}
				log.Printf("leave broadcast resp: %v", resp)
			case *pb.WsLeaveRoomRequest_RoomId:
				hdr := &pb.MessageHeader{
					PlayerId: loginInfo.PlayerId,
				}
				_, err := playClient.LeaveRoom(ctx, hdr, &pb.LeaveRoomRequest{Id: req.LeaveRoom.GetRoomId()})
				if err != nil {
					log.Printf("[%s] leaving room: %v", roomIdStr(req.LeaveRoom.GetRoomId()), err)
				} else {
					log.Printf("[%s] left room", roomIdStr(req.LeaveRoom.GetRoomId()))
				}
			}

		case *pb.WsRequest_SyncMatchTime:
			hdr := &pb.MessageHeader{
				PlayerId:       loginInfo.PlayerId,
				UnknownField_8: req.SyncMatchTime.RoomId_2,
			}
			if resp, err := pb.DiscardHeader(playClient.SyncMatchTime(ctx, hdr, &pb.SyncMatchTimeRequest{Ts: req.SyncMatchTime.Ts})); err != nil {
				log.Printf("[_|%d|_|_] sending match time sync: %v", req.SyncMatchTime.GetRoomId_2(), err)
			} else {
				log.Printf("[_|%d|_|_] sync match time response: %v", req.SyncMatchTime.GetRoomId_2(), resp)
			}

		case *pb.WsRequest_GetPlayerInfo:
			// TODO hangs sometimes
			go func() {
				getPlayerInfoReq := &pb.GetPlayerInfoRequest{
					InfoOptions: &pb.InfoOptions{
						UnknownField_1: proto.Int64(1),
						UnknownField_2: proto.Int64(0),
					},
				}
				switch info := req.GetPlayerInfo.Info.(type) {
				case *pb.WsGetPlayerInfoRequest_Id:
					getPlayerInfoReq.PlayerId = proto.Int64(info.Id)
				case *pb.WsGetPlayerInfoRequest_Name:
					getPlayerInfoReq.PlayerName = proto.String(info.Name)
				}
				resp, err := pb.DiscardHeader(playClient.GetPlayerInfo(ctx, &pb.MessageHeader{PlayerId: loginInfo.PlayerId}, getPlayerInfoReq))
				if err != nil {
					log.Printf("getting player info: %v", err)
				}
				log.Printf("player info:\n%s", prototext.Format(resp))
				wsRespCh <- &pb.WsResponse{Resp: &pb.WsResponse_GetPlayerInfo{GetPlayerInfo: resp}}
			}()

		case *pb.WsRequest_StartAutomatch:
			resp, err := pb.DiscardHeader(playClient.StartAutomatch(ctx, &pb.MessageHeader{PlayerId: loginInfo.PlayerId}, &pb.StartAutomatchRequest{
				PresetId: req.StartAutomatch.PresetId,
			}))
			if err != nil {
				log.Printf("starting automatch: %v", err)
			}
			if resp.GetErrorCode() == 0 {
				log.Printf("automatch started: preset=%d", req.StartAutomatch.GetPresetId())
			} else {
				log.Printf("starting automatch: %v", resp)
			}

		case *pb.WsRequest_StopAutomatch:
			resp, err := pb.DiscardHeader(playClient.StopAutomatch(ctx, &pb.MessageHeader{PlayerId: loginInfo.PlayerId}, &pb.StopAutomatchRequest{}))
			if err != nil {
				log.Printf("stopping automatch: %v", err)
			}
			if resp.GetErrorCode() == 0 {
				log.Printf("automatch stopped")
			} else {
				log.Printf("stopping automatch: %v", resp)
			}

		case *pb.WsRequest_Move:
			hdr := &pb.MessageHeader{
				PlayerId:       loginInfo.PlayerId,
				UnknownField_8: req.Move.RoomId_2,
			}
			extResp, err := playClient.Move(ctx, hdr, req.Move.Move)
			if err != nil {
				log.Printf("sending move: %v", err)
				return
			}
			log.Printf("move response: %v", extResp.Resp)
		}
	}
}

func main() {
	flag.Parse()

	http.HandleFunc("/", handler)

	log.Printf("serving...")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), nil))
}
