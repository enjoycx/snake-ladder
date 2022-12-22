package combat_hander

import (
	"snake-ladder/combatserver/combat"
	"snake-ladder/combatserver/controller"
	"snake-ladder/controller/pb"
)

func RegisterMsgMapFunc() {
	combat.MsgMap[pb.Combat_MsgID_GetMap] = controller.GetMap
	combat.MsgMap[pb.Combat_MsgID_Rool] = controller.Rool
}
