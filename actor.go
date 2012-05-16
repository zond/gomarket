
package gomarket

type Actor interface {
	Asks() map[*Ask]bool
	Bids() map[*Bid]bool
}

type SimpleActor struct {
	name string
	asks map[*Ask]bool
	bids map[*Bid]bool
}
func NewSimpleActor(name string) *SimpleActor {
	return &SimpleActor{name, make(map[*Ask]bool), make(map[*Bid]bool)}
}
func (a *SimpleActor) String() string {
	return a.name
}
func (a *SimpleActor) Ask(units float32, resource interface{}, price float32) {
	a.asks[&Ask{&Order{units, resource, a}, price}] = true
}
func (a *SimpleActor) Bid(units float32, resource interface{}, price float32) {
	a.bids[&Bid{&Order{units, resource, a}, price}] = true
}
func (a *SimpleActor) Asks() map[*Ask]bool {
	return a.asks
}
func (a *SimpleActor) Bids() map[*Bid]bool {
	return a.bids
}

