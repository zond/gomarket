
package gomarket

import (
	"fmt"
)

type Market struct {
	Actors map[Actor]bool
}

func (m *Market) Trade() {
	asks := make(map[interface{}][]*Ask)
	bids := make(map[interface{}][]*Bid)
	for actor,_ := range m.Actors {
		for ask,_ := range actor.Asks() {
			asks[ask.Resource] = append(asks[ask.Resource], ask)
		}
		for bid,_ := range actor.Bids() {
			bids[bid.Resource] = append(bids[bid.Resource], bid)
		}
	}
	fmt.Println("asks:", asks)
	fmt.Println("bids:", bids)
}

type Order struct {
	Units float32
	Resource interface{}
	Actor Actor
}
func (o *Order) String() string {
	return fmt.Sprint(o.Units, "*", o.Resource, "@", o.Actor)
}

type Ask struct {
	*Order
	MinimumPrice float32
}
func (a *Ask) String() string {
	return fmt.Sprint(a.Order.String(), ">", a.MinimumPrice)
}

type Bid struct {
	*Order
	MaximumPrice float32
}
func (b *Bid) String() string {
	return fmt.Sprint(b.Order.String(), "<", b.MaximumPrice)
}

