package ws

import (
	"net/http"
	"runtime"
	"strategy_middle/logging"
	//"strategy_middle/models"
	"strategy_middle/rabbitmq"

	//"time"
	//"Goroutine"
	//"os"
	"bytes"
	"encoding/json"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func GetGID() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return n
}

type Senddata struct {
	Id   int    `json:"id"`
	Pong string `json:"pong"`
	Msg  string `json:"msg"`
}

type Recvdata struct {
	Strategy_name string `json:"strategy_name"`
	Op            string `json:"op"`
	Account       string `json:"account"`
}

// Ping webSocket请求Ping 返回pong
func Ping(c *gin.Context) {
	//升级get请求为webSocket协议
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}



	rabbitmq.Allws =append(rabbitmq.Allws,ws)
	//fmt.Println("rabbitmq.Allws:",rabbitmq.Allws)
	defer ws.Close()
	
	for {
		// 读取ws中的数据
		mt, message, err := ws.ReadMessage()
		if err != nil {
			// 客户端关闭连接时也会进入
			logging.Error(err,mt)
			break
		}
		// fmt.Printf("mt : %T", mt)
		// fmt.Println("mt=",mt)
		//发送到rabbitmq
		rabbitmq.Send("test1", string(message))

		msg := &Recvdata{}
		json.Unmarshal(message, msg)

		logging.Info(msg)
		logging.Info(msg.Strategy_name)
		//fmt.Println(msg)
		//fmt.Println(mt)
		//fmt.Println(message)
		//fmt.Printf("%T", message)
		//fmt.Println(GetGID())
		///////////////////////////////////////////////////////
		// var strategy_info models.Strategy_info
		// db := models.DB
		// //db.Where("strategy_name = ?", msg.Strategy_name).First(&strategy_info)
		// //fmt.Println("strategy_info=", strategy_info.Strategy_status)

		// var status string
		// if msg.Op == "start" {
		// 	status = "Proceeding"
		// } else if msg.Op == "stop" {
		// 	status = "Canceled"
		// }

		// if err := db.Model(&strategy_info).Where("strategy_name = ?", msg.Strategy_name).Update("update_time", time.Now()).Update("strategy_status", status).Error; err != nil {
		// 	fmt.Println(err)
		// } else {
		// 	var st models.Strategy_info = models.Strategy_info{Strategy_name: msg.Strategy_name, Update_time: time.Now(), Strategy_status: status}
		// 	s, err := json.Marshal(st)
		// 	err = ws.WriteMessage(mt, []byte(s))
		// 	if err != nil {
		// 		break
		// 	}
		// }
		///////////////////////////////////////////////////////

		//var st Senddata = Senddata{Id: 1, Pong: "pong", Msg: "hello,girl!"}
		//s, err := json.Marshal(st)

		// 如果客户端发送ping就返回pong,否则数据原封不动返还给客户端
		//if string(message) == "ping" {
		//	message = []byte("pong")
		//}
		// 写入ws数据 二进制返回
		//err = ws.WriteMessage(mt, []byte(s))
		// 返回JSON字符串，借助gin的gin.H实现
		// v := gin.H{"message": msg}
		// err = ws.WriteJSON(v)
		//if err != nil {
		//	break
		//}
	}



}
