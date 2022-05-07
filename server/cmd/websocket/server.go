package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/websocket"
)

func main() {
	http.HandleFunc("/ws", handleWebSocket)
	log.Fatal(http.ListenAndServe(":9090", nil))
}
func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	// 将http连接升级为websocket连接
	u := &websocket.Upgrader{
		// 来自什么的请求都可以
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	c, err := u.Upgrade(w, r, nil)
	defer c.Close()
	if err != nil {
		fmt.Printf("cannot upgrade %v\n", err)
		return
	}

	// 这个channl来记录是否有err出现
	done := make(chan struct{})
	//收消息
	go func() {
		for {
			m := make(map[string]interface{})
			err := c.ReadJSON(&m)
			if err != nil {
				if !websocket.IsCloseError(err,
					websocket.CloseGoingAway,
					websocket.CloseNormalClosure) {
					fmt.Printf("unexpected read error: %v\n", err)
				}
				// 向done写入一个错误标记
				done <- struct{}{}
				break
			}
			fmt.Printf("message received : %v\n", m)
		}

	}()
	i := 0
	for {
		i++
		err = c.WriteJSON(map[string]string{
			"hello":  "websocket",
			"msg_id": strconv.Itoa(i),
		})
		if err != nil {
			fmt.Printf("cannot WriteJSON: %v\n", err)
		}
		select {
		case <-time.After(200 * time.Millisecond):
		//有错误就退出
		case <-done:
			return
		}
		time.Sleep(200 * time.Millisecond)
	}
}
