package tornado

import (
	"context"
)

func Run() {
	var err error
	ctx0 := context.Background()
	msgCh := make(chan string)

	ctx1, cancel1 := context.WithCancel(ctx0)
	defer cancel1()
	err = engine.Start(ctx1, "engine", msgCh)
	if err != nil {
		return
	}

	// 等待停止信号

	// 停止模块
	cancel1()
	<-msgCh
}
