
package gomarket

type Actor interface {
	Asks() map[*Order]bool
	Bids() map[*Order]bool
	Buy(float64, Resource, float64, Actor)
	Sell(float64, Resource, float64, Actor)
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
func (a *TestActor) Buy(units float64, resource Resource, price float64, seller Actor) {
	a.BuySums[resource] = a.BuySums[resource] + units
	a.BuyPrices[resource] = price
	seller.Sell(units, resource, price, a)
}
func (a *TestActor) Sell(units float64, resource Resource, price float64, buyer Actor) {
	a.SellSums[resource] = a.SellSums[resource] + units
	a.SellPrices[resource] = price
}

