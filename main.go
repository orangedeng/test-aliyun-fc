package main

import (
	"context"
	"fmt"

	"github.com/aliyun/fc-runtime-go-sdk/fc"
)

type StructEvent struct {
	Key string `json:"key"`
}

func HandleRequest(ctx context.Context, event StructEvent) (string, error) {
	fmt.Printf("hello, %s!\n", event.Key)
	return fmt.Sprintf("hello, %s!", event.Key), nil
}

func main() {
	fc.Start(HandleRequest)
}
