
package main

import (
	"github.com/davyxu/cellnet"
	"github.com/davyxu/cellnet/peer"
	"github.com/davyxu/cellnet/proc"

	_ "github.com/davyxu/cellnet/peer/tcp"
	_ "github.com/davyxu/cellnet/proc/tcp"

	"gos_server/proto/s2s_proto"
	"gos_server/config"
	"gos_server/center_server/db"
)


func TcpService() {

	// 创建一个事件处理队列，整个服务器只有这一个队列处理事件，服务器属于单线程服务器
	queue := cellnet.NewEventQueue()

	// 创建一个tcp的侦听器，名称为server，连接地址为127.0.0.1:8801，所有连接将事件投递到queue队列,单线程的处理（收发封包过程是多线程）
	p := peer.NewGenericPeer("tcp.Acceptor", "CenterServerTcpService"+config.CenterServerInner, config.CenterServerInner, queue)

	// 设定封包收发处理的模式为tcp的ltv(Length-Type-Value), Length为封包大小，Type为消息ID，Value为消息内容
	// 每一个连接收到的所有消息事件(cellnet.Event)都被派发到用户回调, 用户使用switch判断消息类型，并做出不同的处理
	proc.BindProcessorHandler(p, "tcp.ltv", func(ev cellnet.Event) {

		switch msg := ev.Message().(type) {
		// 有新的连接
		case *cellnet.SessionAccepted:
			config.LogCenterServer.Debugln("CenterServerTcpService accepted")
		// 有连接断开
		case *cellnet.SessionClosed:
			config.LogCenterServer.Debugln("CenterServerTcpService session closed: ", ev.Session().ID())
		// 收到某个连接的ChatREQ消息
		case *s2s_proto.ChatREQ:

			// 写db
			var c = db.Gmdb.DB("test").C("Chat")
			if c == nil {
				config.LogCenterServer.Errorln("Collection does not exist in db test!")
			}

			// 准备回应的消息
			ack := s2s_proto.ChatACK{
				Content: msg.Content,       // 聊天内容
				Id:      ev.Session().ID(), // 使用会话ID作为发送内容的ID
			}

			// 在Peer上查询SessionAccessor接口，并遍历Peer上的所有连接，并发送回应消息（即广播消息）
			p.(cellnet.SessionAccessor).VisitSession(func(ses cellnet.Session) bool {

				ses.Send(&ack)

				return true
			})

		}

	})

	// 开始侦听
	p.Start()

	// 事件队列开始循环
	queue.StartLoop()

	// 阻塞等待事件队列结束退出( 在另外的goroutine调用queue.StopLoop() )
	queue.Wait()

}
