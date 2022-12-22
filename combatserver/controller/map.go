package controller

import (
	"snake-ladder/combatserver/combat"
	"snake-ladder/controller/pb"

	"google.golang.org/protobuf/proto"
)

func GetMap(user *combat.UserInfo, msg []byte) (proto.Message, error) {
	combatMap := combat.CombatMaps[user.RoomID]
	items := []*pb.Item{}
	for _, v := range combatMap.Items {
		items = append(items, &pb.Item{
			Type:   v.Type,
			Target: v.Target,
			Pos:    v.Pos,
		})
	}
	combatMapPb := &pb.CombatMap{
		Length: combatMap.Length,
		Items:  items,
	}
	combatInfo := &pb.CombatInfo{
		CombatMap: combatMapPb,
		Uid:       user.UserID,
	}
	if user.UserID == combatMap.Player1.Uid {
		combatInfo.Pos = combatMap.Player1.Pos
		combatInfo.Rool = combatMap.Player1.Rool
	} else {
		combatInfo.Pos = combatMap.Player2.Pos
		combatInfo.Rool = combatMap.Player2.Rool
	}
	if combatMap.Player2 != nil {
		combatInfo.Start = true
	}
	return combatInfo, nil
}
