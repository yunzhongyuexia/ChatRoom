package webSocket

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"server/model"
)

// 变量定义初始化
var (
	ug = websocket.Upgrader{
		//配置读写缓冲大小
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	clientMsg     = msg{}
	sMsg          = make(chan msg)
	offline       = make(chan *websocket.Conn)
	clientMsgData = clientMsg // 临时存储 clientMsg 数据
)

// client & serve 的消息体
type msg struct {
	Status int64           `json:"status"`
	Data   model.Message   `json:"data"`
	Conn   *websocket.Conn `json:"conn"`
}

// WcDemoV0 标准demo
func WcDemoV0(ctx *gin.Context) {

	// 这里可以设置允许所有的来源
	ug.CheckOrigin = func(r *http.Request) bool { return true }

	conn, err := ug.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		log.Println(err)
		return
	}

	defer func() {
		_ = conn.Close()
	}()

	for {
		messageType, p, err1 := conn.ReadMessage()
		if err1 != nil {
			log.Println(err1)
			return
		}

		//fmt.Println(string(p))
		//if clientMsg.Data.name != "" {
		//	if clientMsg.Status == msgTypeOnline { // 进入房间，建立连接
		//		roomId, _ := getRoomId()
		//
		//		enterRooms <- wsClients{
		//			Conn:       c,
		//			RemoteAddr: c.RemoteAddr().String(),
		//			Uid:        clientMsg.Data.Uid,
		//			Username:   clientMsg.Data.Username,
		//			RoomId:     roomId,
		//			AvatarId:   clientMsg.Data.AvatarId,
		//		}
		//	}
		//
		//	_, serveMsg := formatServeMsgStr(clientMsg.Status, c)
		//	sMsg <- serveMsg
		//}

		if err1 = conn.WriteMessage(messageType, p); err1 != nil {
			log.Println(err1)
			return
		}
	}
}
