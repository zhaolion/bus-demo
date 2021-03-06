package biz4

import (
	"fmt"

	"github.com/zhaolion/bus-demo/bus"
	"github.com/zhaolion/bus-demo/model"
)

func init() {
	bus.AddHandler("biz4", Produce)
	bus.AddEventListener(Produce)
}

func Produce(msg *model.Message) error {
	fmt.Printf("[biz4-msg]: content %v\n", msg.Content)
	return nil
}
