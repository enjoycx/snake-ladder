package kissnet

import (
	"fmt"
	"net"
	"sync/atomic"
)

type TcpAcceptor struct {
	Acceptor
	listener *net.TCPListener
	Addr     *net.TCPAddr
	running  int32
}

func NewTcpAcceptor(port int, cb ClientCB) (*TcpAcceptor, error) {
	ep := fmt.Sprintf("0.0.0.0:%d", port)
	tcpAddr, err := net.ResolveTCPAddr("tcp", ep)
	if err != nil {
		return nil, err
	}
	ln, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		return nil, err
	}
	return &TcpAcceptor{Addr: tcpAddr, listener: ln, Acceptor: Acceptor{clientCB: cb}}, nil

}
func (this *TcpAcceptor) Run() error {
	go this.accept()
	return nil
}

func (this *TcpAcceptor) accept() {
	atomic.StoreInt32(&this.running, 1)
	for {
		tcpConn, err := this.listener.AcceptTCP()
		if err != nil {
			return
		}
		conn := newConnection(tcpConn)
		go conn.start()
	}
}

func (this *TcpAcceptor) Close() error {
	this.listener.Close()
	return nil
}

func (this *TcpAcceptor) IsRunning() bool {
	return atomic.LoadInt32(&this.running) > 0
}
