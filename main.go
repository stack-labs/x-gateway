package main

import (
	"github.com/micro/go-micro"
	"github.com/micro-in-cn/x-gateway/cmd"
)

func main() {
	cmd.Init(
		micro.AfterStop(cleanWork),
	)
}
