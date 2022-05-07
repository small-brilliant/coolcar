package ws

import (
	"context"
	"coolcar/car/mq"
	"net/http"

	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

// Handler creates a websocket http handler
func Handler(u *websocket.Upgrader, sub mq.Subscriber, logger *zap.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c, err := u.Upgrade(w, r, nil)
		defer c.Close()
		if err != nil {
			logger.Warn("cannot upgrade", zap.Error(err))
			return
		}
		msgs, cleanUp, err := sub.Subscribe(context.Background())
		defer cleanUp()
		if err != nil {
			logger.Error("cannot subscribe", zap.Error(err))
			//返回给用户错误
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		// 这个channl来记录是否有err出现
		done := make(chan struct{})
		//收消息
		go func() {
			for {
				_, _, err := c.ReadMessage()
				if err != nil {
					if !websocket.IsCloseError(err,
						websocket.CloseGoingAway,
						websocket.CloseNormalClosure) {
						logger.Warn("unexpected read error", zap.Error(err))
					}
					// 向done写入一个错误标记
					done <- struct{}{}
					break
				}
			}

		}()
		// 发消息：将从RabbitMQ中受到消息转出去。
		for {
			select {
			case msg := <-msgs:
				//收到了信息，就发给websocket客户端
				err = c.WriteJSON(msg)
				if err != nil {
					logger.Warn("cannot write JSON", zap.Error(err))
				}
			//有错误就退出
			case <-done:
				return
			}
		}
	}

}
