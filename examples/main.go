package main

import (
	"context"

	"github.com/stemstr/blastr"
)

func main() {
	client, _ := blastr.New("nsec1xxx")

	_ = client.SendText(context.Background(), "hello, world!")
}
