package main

import (
	"fmt"
	"snake-ladder/combatserver/combat"
	combat_hander "snake-ladder/combatserver/hander"
	"snake-ladder/kissnet"

	"github.com/sirupsen/logrus"
)

var gRPCConnectorMgr = kissnet.RpcConnectorMgr{
	RpcConnMap: make(map[string]*kissnet.RpcConnector),
}
var gAcceptor kissnet.IAcceptor

func main() {

	port := 80

	combat.InitRoom()
	event := kissnet.NewNetEvent()
	logrus.Info("acceptor start")
	combat_hander.RegisterMsgMapFunc()
	var err error
	gAcceptor, err = kissnet.AcceptorFactory(
		"ws",
		port,
		combat.CombatClientCB,
	)
	if err != nil {
		fmt.Println(err)
		logrus.WithFields(logrus.Fields{
			"error": err,
		}).Fatal("AcceptorFactory error")
		return
	}

	gAcceptor.Run()
	event.EventLoop()
	gAcceptor.Close()
	logrus.Info("acceptor end")
}
