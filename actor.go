
package gomarket

type Actor interface {
	Asks() map[*Order]bool
	Bids() map[*Order]bool
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
func (a *SimpleActor) Ask(units float32, resource interface{}, price float32) {
	a.asks[&Order{units, resource, price, a}] = true
}
func (a *SimpleActor) Bid(units float32, resource interface{}, price float32) {
	a.bids[&Order{units, resource, price, a}] = true
}
func (a *SimpleActor) Asks() map[*Order]bool {
	return a.asks
}
func (a *SimpleActor) Bids() map[*Order]bool {
	return a.bids
}

