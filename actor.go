
package gomarket

type Actor interface {
	Asks() map[*Order]bool
	Bids() map[*Order]bool
	Buy(*Order, float64)
	Sell(*Order, float64)
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
func (a *TestActor) Buy(ask *Order, price float64) {
	a.BuySums[ask.Resource] = a.BuySums[ask.Resource] + ask.Units
	a.BuyPrices[ask.Resource] = price
}
func (a *TestActor) Sell(bid *Order, price float64) {
	a.SellSums[bid.Resource] = a.SellSums[bid.Resource] + bid.Units
	a.SellPrices[bid.Resource] = price
}

