
package gomarket

type Actor interface {
	Asks() map[*Order]bool
	Bids() map[*Order]bool
	Buy(*Order, *Order, float64)
	Deliver(*Order, *Order, float64)
}

type TestActor struct {
	asks map[*Order]bool
	bids map[*Order]bool
	BuySums map[Resource]float64
	SellSums map[Resource]float64
	BuyPrices map[Resource]float64
	SellPrices map[Resource]float64
}
func NewTestActor() *TestActor {
	return &TestActor{
		make(map[*Order]bool), 
		make(map[*Order]bool), 
		make(map[Resource]float64), 
		make(map[Resource]float64),
		make(map[Resource]float64), 
		make(map[Resource]float64)}
}
func (a *TestActor) Ask(units float64, resource Resource, price float64) {
	a.asks[&Order{units, resource, price, a}] = true
}
func (a *TestActor) Bid(units float64, resource Resource, price float64) {
	a.bids[&Order{units, resource, price, a}] = true
}
func (a *TestActor) Asks() map[*Order]bool {
	return a.asks
}
func (a *TestActor) Bids() map[*Order]bool {
	return a.bids
}
func (a *TestActor) Buy(bid, ask *Order, price float64) {
	a.BuySums[bid.Resource] = a.BuySums[bid.Resource] + bid.Units
	a.BuyPrices[bid.Resource] = price
	ask.Actor.Deliver(bid, ask, price)
}
func (a *TestActor) Deliver(bid, ask *Order, price float64) {
	a.SellSums[ask.Resource] = a.SellSums[ask.Resource] + ask.Units
	a.SellPrices[ask.Resource] = price
}

