
package gomarket

import (
	"fmt"
)

type Market struct {
	Actors map[Actor]bool
}
func NewMarket() *Market {
	return &Market{make(map[Actor]bool)}
}
func (m *Market) Trade() {
	asks, bids, ask_sums, bid_sums := m.createSums()
	for resource,sum := range ask_sums {
		if sum > bid_sums[resource] {
			fmt.Println("abundance of", resource)
		} else {
			fmt.Println("scarcity of", resource)
		}
	}
}
func (m *Market) createSums() (asks, bids map[interface{}][]*Order, ask_sums, bid_sums map[interface{}]float32) {
	asks = make(map[interface{}][]*Order)
	bids = make(map[interface{}][]*Order)
	ask_sums = make(map[interface{}]float32)
	bid_sums = make(map[interface{}]float32)
	for actor,_ := range m.Actors {
		for ask,_ := range actor.Asks() {
			asks[ask.Resource] = append(asks[ask.Resource], ask)
			ask_sums[ask.Resource] += ask.Units
		}
		for bid,_ := range actor.Bids() {
			bids[bid.Resource] = append(bids[bid.Resource], bid)
			bid_sums[bid.Resource] += bid.Units
		}
	}
	return
}

type Order struct {
	Units float32
	Resource interface{}
	Price float32
	Actor Actor
}
func (o *Order) String() string {
	return fmt.Sprint(o.Units, "*", o.Resource, "á", o.Price, "@", o.Actor)
}


