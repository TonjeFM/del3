package main

import (
	"fmt"
	"github.com/tonjefm/del3/event"
	"github.com/tonjefm/del3/state"
)

func main() {
	fmt.Println(state.ViewState())
	fmt.Println(event.PutIn("Kylling"))
	fmt.Println(event.CrossRiver("Kylling"))
	fmt.Println(event.TakeOut("Kylling"))
}