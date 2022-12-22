package main

import (
	"bytes"
	"encoding/binary"
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
	"os/signal"
	"snake-ladder/controller/pb"
	"time"

	"github.com/gorilla/websocket"
	"google.golang.org/protobuf/proto"
)

var addr = flag.String("addr", "127.0.0.1:80", "http service address")

func clientSendRecv(c *websocket.Conn, msgID pb.Combat_MsgID_MsgID, msg []byte) ([]byte, error) {
	sendMsg := new(bytes.Buffer)
	binary.Write(sendMsg, binary.LittleEndian, uint16(msgID))
	sendMsg.Write(msg)

	err := c.WriteMessage(websocket.BinaryMessage, sendMsg.Bytes())
	if err != nil {
		return nil, err
	}
	_, recvMsg, err := c.ReadMessage()
	return recvMsg, nil
}

type PlayerInfo struct {
	Uid  string
	Pos  int32
	Rool bool
}

func main() {
	flag.Parse()
	log.SetFlags(0)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "ws", Host: *addr, Path: "/ws"}
	log.Printf("connecting to %s", u.String())

	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer conn.Close()
	playerInfo := &PlayerInfo{
		Uid: "2",
	}

	joinRoom(playerInfo, conn)
	time.Sleep(time.Second)
	for i := 0; i < 10000; i++ {
		time.Sleep(time.Second)
		start := getMap(playerInfo, conn)
		if start {
			break
		}
	}
	for i := 0; i < 10000; i++ {
		time.Sleep(time.Second)
		getMap(playerInfo, conn)
		end := rool(playerInfo, conn)
		if end {
			break
		}
	}
}

func joinRoom(playerInfo *PlayerInfo, conn *websocket.Conn) {
	pbMsg := &pb.JoinRoom{
		RoomID: "1",
		Uid:    playerInfo.Uid,
	}
	msg, _ := proto.Marshal(pbMsg)
	recvMsg, err := clientSendRecv(conn, pb.Combat_MsgID_JoinRoom, msg)
	if err != nil {
		fmt.Println(err)
	}

	recvMsgID := binary.LittleEndian.Uint16(recvMsg)
	fmt.Println(recvMsgID)
}

func getMap(playerInfo *PlayerInfo, conn *websocket.Conn) bool {
	pbMsg := &pb.JoinRoom{
		RoomID: "1",
		Uid:    playerInfo.Uid,
	}
	msg, _ := proto.Marshal(pbMsg)
	recvMsg, err := clientSendRecv(conn, pb.Combat_MsgID_GetMap, msg)
	if err != nil {
		fmt.Println(err)
	}
	combatInfo := &pb.CombatInfo{}
	err = proto.Unmarshal(recvMsg[2:], combatInfo)
	if err != nil {
		fmt.Println(err)
	}
	recvMsgID := binary.LittleEndian.Uint16(recvMsg)
	fmt.Println(recvMsgID)
	fmt.Println(combatInfo)
	if combatInfo.Uid == playerInfo.Uid {
		playerInfo.Pos = combatInfo.Pos
		playerInfo.Rool = combatInfo.Rool
	}
	return combatInfo.Start
}

func rool(playerInfo *PlayerInfo, conn *websocket.Conn) bool {
	if !playerInfo.Rool {
		return false
	}
	pbMsg := &pb.JoinRoom{
		RoomID: "1",
		Uid:    playerInfo.Uid,
	}
	msg, _ := proto.Marshal(pbMsg)
	recvMsg, err := clientSendRecv(conn, pb.Combat_MsgID_Rool, msg)
	if err != nil {
		fmt.Println(err)
	}
	roolRsp := &pb.RoolRsp{}
	err = proto.Unmarshal(recvMsg[2:], roolRsp)
	if err != nil {
		fmt.Println(err)
	}
	recvMsgID := binary.LittleEndian.Uint16(recvMsg)
	fmt.Println(recvMsgID)
	fmt.Println(roolRsp)
	if roolRsp.Uid == playerInfo.Uid {
		playerInfo.Pos = roolRsp.Pos
		playerInfo.Rool = roolRsp.Rool
	}

	if roolRsp.Pos == 99 {
		if roolRsp.Uid == playerInfo.Uid {
			fmt.Println("胜利")
		} else {
			fmt.Println("失败")
		}
		return true
	}
	return false
}
