package combat

var CombatMaps map[string]*CombatMap

type CombatMap struct {
	Length  int32
	Items   []*Item
	Player1 *PlayerInfo
	Player2 *PlayerInfo
}

type PlayerInfo struct {
	Uid  string
	Pos  int32
	Rool bool
}

type Item struct {
	Type   int32 //1 梯子 2蛇
	Target int32 //跳转到指定格子
	Pos    int32 //所在位置·
}

const MapLength = 99

func CreatRoom(roomID, uid string) {
	if CombatMaps == nil {
		CombatMaps = map[string]*CombatMap{}
	}
	if CombatMaps[roomID] == nil {
		CombatMaps[roomID] = &CombatMap{
			Length: MapLength,
			Player1: &PlayerInfo{
				Uid:  uid,
				Rool: true,
			},
		}
	} else {
		CombatMaps[roomID].Player2 = &PlayerInfo{
			Uid: uid,
		}
	}
}
