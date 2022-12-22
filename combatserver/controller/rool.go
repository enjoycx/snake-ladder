package controller

import (
	"math/rand"
	"snake-ladder/combatserver/combat"
	"snake-ladder/controller/pb"

	"google.golang.org/protobuf/proto"
)

func Rool(user *combat.UserInfo, msg []byte) (proto.Message, error) {
	combatMap := combat.CombatMaps[user.RoomID]
	var player *combat.PlayerInfo
	var player2 *combat.PlayerInfo
	if combatMap.Player1.Uid == user.UserID {
		player = combatMap.Player1
		player2 = combatMap.Player2
	} else {
		player = combatMap.Player2
		player2 = combatMap.Player1
	}
	if !player.Rool {
		return nil, nil
	}
	roolNum := rand.Int31n(6) + 1
	player.Pos += roolNum
	overflow := player.Pos - combat.MapLength
	if overflow > 0 {
		player.Pos -= 2 * overflow
	}
	player.Rool = false
	player2.Rool = true

	roolRsp := &pb.RoolRsp{
		Pos:  player.Pos,
		Rool: player.Rool,
		Uid:  user.UserID,
	}
	return roolRsp, nil
}
