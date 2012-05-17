
package gomarket

import (
	"testing"
)

type TestActor struct {
	asks map[*Order]bool
	bids map[*Order]bool
	BuySums map[interface{}]float64
	SellSums map[interface{}]float64
	BuyPrices map[interface{}]float64
	SellPrices map[interface{}]float64
}
func NewTestActor() *TestActor {
	return &TestActor{
		make(map[*Order]bool), 
		make(map[*Order]bool), 
		make(map[interface{}]float64), 
		make(map[interface{}]float64),
		make(map[interface{}]float64), 
		make(map[interface{}]float64)}
}
func (a *TestActor) Ask(units float64, resource interface{}, price float64) {
	a.asks[&Order{units, resource, price, a}] = true
}
func (a *TestActor) Bid(units float64, resource interface{}, price float64) {
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

func TestOneSellerNoBuyers(t *testing.T) {
	m := NewMarket()
	seller := NewTestActor()
	shoes := "shoes"
	seller.Ask(10.0, shoes, 5.0)
	m.Actors[seller] = true
	m.Trade()
	if m.Prices[shoes] != 5.0 {
		t.Fail()
	}
	if len(seller.BuySums) != 0 {
		t.Fail()
	}
	if len(seller.SellSums) != 0 {
		t.Fail()
	}
}

func TestOneSellerOneBuyerNoDeal(t *testing.T) {
	m := NewMarket()
	seller := NewTestActor()
	buyer := NewTestActor()
	shoes := "shoes"
	seller.Ask(10.0, shoes, 15.0)
	buyer.Bid(10.0, shoes, 10.0)
	m.Actors[seller] = true
	m.Actors[buyer] = true
	m.Trade()
	if m.Prices[shoes] != 12.5 {
		t.Fail()
	}
	if len(seller.BuySums) != 0 {
		t.Fail()
	}
	if len(seller.SellSums) != 0 {
		t.Fail()
	}
	if len(buyer.BuySums) != 0 {
		t.Fail()
	}
	if len(buyer.SellSums) != 0 {
		t.Fail()
	}
}

func TestOneSellerOneBuyerDeal(t *testing.T) {
	m := NewMarket()
	seller := NewTestActor()
	buyer := NewTestActor()
	shoes := "shoes"
	seller.Ask(10.0, shoes, 5.0)
	buyer.Bid(10.0, shoes, 10.0)
	m.Actors[seller] = true
	m.Actors[buyer] = true
	m.Trade()
	if m.Prices[shoes] != 7.5 {
		t.Fail()
	}
	if len(seller.BuySums) != 0 {
		t.Fail()
	}
	if len(seller.SellSums) != 1 {
		t.Fail()
	}
	if seller.SellSums[shoes] != 10.0 {
		t.Fail()
	}
	if seller.SellPrices[shoes] != 7.5 {
		t.Fail()
	}
	if len(buyer.BuySums) != 1 {
		t.Fail()
	}
	if len(buyer.SellSums) != 0 {
		t.Fail()
	}
	if buyer.BuySums[shoes] != 10.0 {
		t.Fail()
	}
	if buyer.BuyPrices[shoes] != 7.5 {
		t.Fail()
	}
}