package combat

import (
	"snake-ladder/kissnet"
	"sync"
)

type UserInfo struct {
	Conn      kissnet.IConnection
	UserID    string
	RoomID    string
	Heartbeat int64 //心跳时间
}

type UserMgr struct {
	roomMap   map[string][]string
	userIDMap map[string]*UserInfo
	connMap   map[kissnet.IConnection]*UserInfo
	num       int64
	mutex     sync.RWMutex
}

var UserMgrInfo *UserMgr = &UserMgr{
	roomMap:   map[string][]string{},
	userIDMap: make(map[string]*UserInfo),
	connMap:   make(map[kissnet.IConnection]*UserInfo),
	num:       int64(0),
}

func (u *UserMgr) GetUserInfoByUserID(UserID string) *UserInfo {
	u.mutex.RLock()
	defer u.mutex.RUnlock()
	if v, ok := u.userIDMap[UserID]; ok {
		return v
	}
	return nil
}
func (u *UserMgr) GetUserInfoByConn(conn kissnet.IConnection) *UserInfo {
	u.mutex.RLock()
	defer u.mutex.RUnlock()
	if v, ok := u.connMap[conn]; ok {
		return v
	}
	return nil
}

func (u *UserMgr) AddUserInfo(roomID, UserID string, conn kissnet.IConnection) *UserInfo {
	u.mutex.Lock()
	defer u.mutex.Unlock()
	userInfo := &UserInfo{
		Conn:   conn,
		UserID: UserID,
		RoomID: roomID,
	}
	if u.roomMap[roomID] == nil {
		u.roomMap[roomID] = []string{}
	}
	exit := false
	for _, v := range u.roomMap[roomID] {
		if v == UserID {
			exit = true
		}
	}
	if !exit {
		u.roomMap[roomID] = append(u.roomMap[roomID], UserID)
	}
	u.connMap[userInfo.Conn] = userInfo
	u.userIDMap[userInfo.UserID] = userInfo
	u.num++
	return userInfo
}

func (u *UserMgr) DelUserInfo(conn kissnet.IConnection) {
	u.mutex.Lock()
	defer u.mutex.Unlock()
	userInfo, ok := u.connMap[conn]
	if !ok {
		return
	}
	for i, v := range u.roomMap[userInfo.RoomID] {
		if v == userInfo.UserID {
			u.roomMap[userInfo.RoomID] = append(u.roomMap[userInfo.RoomID][:i], u.roomMap[userInfo.RoomID][i+1:]...)
			break
		}
	}

	u.num--

	delete(u.connMap, conn)
	delete(u.userIDMap, userInfo.UserID)
	delete(u.roomMap, userInfo.RoomID)
}
func (u *UserMgr) Close() {
	for k := range u.connMap {
		k.Close()
	}
}

func (u *UserMgr) GetUserNum() int64 {
	u.mutex.RLock()
	defer u.mutex.RUnlock()
	return u.num
}
