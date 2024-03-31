package handlers

import "github.com/gin-gonic/gin"

func Chats(c *gin.Context, clients map[string]chan []byte) {
	c.Header("Content-Type", "text/event-stream")
	c.Header("Cache-Control", "no-cache")
	c.Header("Connection", "keep-alive")

	//clients["aaa"] = clientChannel
	clientChannel := clients["aaa"]
	if clientChannel == nil {
		// 新的数据通道
		clientChannel := make(chan []byte)
		clients["aaa"] = clientChannel
	}

	defer func() {
		// 关闭
		close(clientChannel)
		delete(clients, "aaa")
	}()

	for {
		select {
		case data := <-clientChannel:
			// 将数据写入响应
			_, err := c.Writer.Write([]byte("data: " + string(data) + "\n\n"))
			if err != nil {
				return
			}
			c.Writer.Flush()
		case <-c.Request.Context().Done():
			return
		}
	}
}
