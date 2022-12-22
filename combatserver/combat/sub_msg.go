package combat

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"snake-ladder/controller/pb"
	"snake-ladder/kissnet"

	"google.golang.org/protobuf/proto"
)

type FuncMsg func(*UserInfo, []byte) (proto.Message, error)

var (
	MsgMap map[pb.Combat_MsgID_MsgID]FuncMsg = make(map[pb.Combat_MsgID_MsgID]FuncMsg)
)

func CombatClientCB(conn kissnet.IConnection, msg []byte) error {
	if msg == nil {
		//退出
		UserMgrInfo.DelUserInfo(conn)
		return nil
	}
	if len(msg) < 2 {
		UserMgrInfo.DelUserInfo(conn)
		return fmt.Errorf("msg len error")
	}
	msgID := binary.LittleEndian.Uint16(msg)
	if msgID == uint16(pb.Combat_MsgID_JoinRoom) {
		joinRoom := &pb.JoinRoom{}
		err := proto.Unmarshal(msg[2:], joinRoom)
		if err != nil {
			return err
		}

		u := UserMgrInfo.GetUserInfoByUserID(joinRoom.Uid)
		if u != nil {
			// 踢掉之前连接
			UserMgrInfo.DelUserInfo(u.Conn)
		}
		UserMgrInfo.AddUserInfo(joinRoom.RoomID, joinRoom.Uid, conn)
		// roolEventChan <- &PlayerInfoEvent{
		// 	RoomID: joinRoom.RoomID,
		// 	Uid:    joinRoom.Uid,
		// }
		CreatRoom(joinRoom.RoomID, joinRoom.Uid)
		BroadcastRoom(joinRoom.RoomID, msg)
	} else {
		ProcMsg(conn, msgID, msg)
	}
	return nil
}

func BroadcastPbRoom(roomID string, message proto.Message, msgID uint16) error {
	m, err := proto.Marshal(message)
	if err != nil {
		return err
	}
	sendMsg := new(bytes.Buffer)
	binary.Write(sendMsg, binary.LittleEndian, msgID)
	sendMsg.Write(m)
	for _, uid := range UserMgrInfo.roomMap[roomID] {
		userInfo := UserMgrInfo.GetUserInfoByUserID(uid)
		err = userInfo.Conn.SendMsg(sendMsg)
		if err != nil {
			continue
		}
	}
	return nil
}

func BroadcastRoom(roomID string, msg []byte) error {
	for _, uid := range UserMgrInfo.roomMap[roomID] {
		SendResponseByte(uid, msg)
	}
	return nil
}

func SendResponseByte(uid string, msg []byte) error {
	userInfo := UserMgrInfo.GetUserInfoByUserID(uid)
	if userInfo == nil {
		return fmt.Errorf("uid:%v conn nil", uid)
	}
	sendMsg := new(bytes.Buffer)
	sendMsg.Write(msg)
	err := userInfo.Conn.SendMsg(sendMsg)
	if err != nil {
		return err
	}
	return nil
}

func SendResponse(conn kissnet.IConnection, message proto.Message, msgID uint16) error {
	m, err := proto.Marshal(message)
	if err != nil {
		return err
	}
	sendMsg := new(bytes.Buffer)
	binary.Write(sendMsg, binary.LittleEndian, msgID)
	sendMsg.Write(m)
	err = conn.SendMsg(sendMsg)
	if err != nil {
		return err
	}
	return nil
}

func ProcMsg(conn kissnet.IConnection, msgID uint16, msg []byte) error {
	f, ok := MsgMap[pb.Combat_MsgID_MsgID(msgID)]
	if !ok {
		u := UserMgrInfo.GetUserInfoByConn(conn)
		if u == nil {
			return fmt.Errorf("(%d) ProcMsg GetUserChat nil", msgID)
		}
		BroadcastRoom(u.RoomID, msg)
		return nil
	}
	u := UserMgrInfo.GetUserInfoByConn(conn)
	if u == nil {
		return fmt.Errorf("(%d) ProcMsg GetUserChat nil", msgID)
	}
	message, err := f(u, msg[2:])
	if err != nil {
		return err
	}
	return SendResponse(conn, message, msgID)
}
