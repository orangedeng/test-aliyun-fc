package main

import (
	"context"
	"fmt"

	"github.com/aliyun/fc-runtime-go-sdk/fc"
)

func HandleRequest(ctx context.Context, event []byte) (string, error) {
	fmt.Printf("event: %s\n", string(event))
	fmt.Println("hello world! 你好，世界!")
	return "hello world! 你好，世界!", nil
}

func main() {
	fc.Start(HandleRequest)
}
