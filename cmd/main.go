/**
 * @author   Liang
 * @create   2022/8/19 11:07
 * @version  1.0
 */

package main

import (
	"fmt"
	"go-template/config"
	"go-template/internal/serve"
)

func main() {

	exitc := make(chan error)

	go func() {
		exitc <- serve.StartHttp(config.HttpPort)
	}()

	go func() {
		exitc <- serve.StartGrpc(config.GrpcPort)
	}()

	fmt.Println("exit", <-exitc)
}
