
package gomarket

import (
	"fmt"
)

type Actor interface {
	Asks() map[*Order]bool
	Bids() map[*Order]bool
	Buy(*Order, float64)
	Sell(*Order, float64)
}

type SimpleActor struct {
	name string
	asks map[*Order]bool
	bids map[*Order]bool
}
func NewSimpleActor(name string) *SimpleActor {
	return &SimpleActor{name, make(map[*Order]bool), make(map[*Order]bool)}
}
func (a *SimpleActor) String() string {
	return a.name
}
func (a *SimpleActor) Ask(units float64, resource interface{}, price float64) {
	a.asks[&Order{units, resource, price, a}] = true
}
func (a *SimpleActor) Bid(units float64, resource interface{}, price float64) {
	a.bids[&Order{units, resource, price, a}] = true
}
func (a *SimpleActor) Asks() map[*Order]bool {
	return a.asks
}
func (a *SimpleActor) Bids() map[*Order]bool {
	return a.bids
}
func (a *SimpleActor) Buy(ask *Order, price float64) {
	fmt.Println(a, " buys ", ask.Units, ask.Resource, " from ", ask.Actor, " á ", price)
}
func (a *SimpleActor) Sell(bid *Order, price float64) {
	fmt.Println(a, " sells ", bid.Units, bid.Resource, " to ", bid.Actor, " á ", price)
}
