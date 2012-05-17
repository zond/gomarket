
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

func Check(t *testing.T,
	ask_units, ask_prices, bid_units, bid_prices []float64, 
	expected_price float64, 
	expected_sells, expected_buys []float64) {
	m := NewMarket()
	sellers := make([]*TestActor, 0)
	buyers := make([]*TestActor, 0)
	product := "product"
	for i := 0; i < len(ask_units); i++ {
		seller := NewTestActor()
		seller.Ask(ask_units[i], product, ask_prices[i])
		m.Actors[seller] = true
		sellers = append(sellers, seller)
	}
	for i := 0; i < len(bid_units); i++ {
		buyer := NewTestActor()
		buyer.Bid(bid_units[i], product, bid_prices[i])
		m.Actors[buyer] = true
		buyers = append(buyers, buyer)
	}
	m.Trade()
	if m.Prices[product] != expected_price {
		t.Error("When selling",ask_units,"for",ask_prices,"and buying",bid_units,"for",bid_prices,"expected price to be",expected_price,"but was",m.Prices[product])
	}
	for i := 0; i < len(expected_sells); i++ {
		if sellers[i].SellSums[product] != expected_sells[i] {
			t.Error("When selling",ask_units,"for",ask_prices,"and buying",bid_units,"for",bid_prices,"expected seller",i,"to sell",expected_sells[i],"units, but sold",sellers[i].SellSums[product],"units.")
		}
		if sellers[i].SellSums[product] > 0 && sellers[i].SellPrices[product] != expected_price {
			t.Error("When selling",ask_units,"for",ask_prices,"and buying",bid_units,"for",bid_prices,"expected seller",i,"to sell for",expected_price,"but sold for",sellers[i].SellPrices[product])
		}
	}
	for i := 0; i < len(expected_buys); i++ {
		if buyers[i].BuySums[product] != expected_buys[i] {
			t.Error("When selling",ask_units,"for",ask_prices,"and buying",bid_units,"for",bid_prices,"expected buyer",i,"to buy",expected_buys[i],"units, but bought",buyers[i].BuySums[product],"units.")
		}
		if buyers[i].BuySums[product] > 0 && buyers[i].BuyPrices[product] != expected_price {
			t.Error("When selling",ask_units,"for",ask_prices,"and buying",bid_units,"for",bid_prices,"expected buyer",i,"to buy for",expected_price,"but bought for",buyers[i].BuyPrices[product])
		}
	}
}

func TestOneSellerNoBuyers(t *testing.T) {
	Check(t,
		[]float64{10.0}, []float64{5.0}, []float64{}, []float64{},
		5.0,
		[]float64{0.0}, []float64{})
}

func TestNoSellersOneBuyer(t *testing.T) {
	Check(t,
		[]float64{}, []float64{}, []float64{10.0}, []float64{10.0},
		10.0,
		[]float64{}, []float64{0.0})
}

func TestOneSellerOneBuyerNoDeal(t *testing.T) {
	Check(t,
		[]float64{10.0}, []float64{5.0}, []float64{10.0}, []float64{2.0},
		3.5,
		[]float64{0.0}, []float64{0.0})
}

func TestOneSellerOneBuyerDeal(t *testing.T) {
	Check(t,
		[]float64{10.0}, []float64{5.0}, []float64{10.0}, []float64{10.0},
		7.5,
		[]float64{10.0}, []float64{10.0})
}

func TestOneSellerManyBuyersDeal(t *testing.T) {
	Check(t,
		[]float64{10.0}, []float64{5.0}, []float64{5.0, 5.0, 5.0}, []float64{10.0, 15.0, 12.0},
		12.0,
		[]float64{10.0}, []float64{0.0, 5.0, 5.0})
}

func TestOneSellerManyBuyersNoDeal(t *testing.T) {
	Check(t,
		[]float64{10.0}, []float64{25.0}, []float64{5.0, 5.0, 5.0}, []float64{10.0, 15.0, 12.0},
		20.0,
		[]float64{0.0}, []float64{0.0, 0.0, 0.0})
}

func TestManySellersOneBuyerDeal(t *testing.T) {
	Check(t,
		[]float64{7.0, 5.0, 15.0}, []float64{5.0, 4.0, 6.0}, []float64{10.0}, []float64{10.0},
		5.0,
		[]float64{5.0, 5.0, 0.0}, []float64{10.0})
}

func TestManySellersOneBuyerNoDeal(t *testing.T) {
	Check(t,
		[]float64{7.0, 5.0, 15.0}, []float64{5.0, 4.0, 6.0}, []float64{10.0}, []float64{1.0},
		2.5,
		[]float64{0.0, 0.0, 0.0}, []float64{0.0})
}

func TestManySellersManyBuyersDealScarcity(t *testing.T) {
	Check(t,
		[]float64{7.0, 5.0, 15.0}, []float64{5.0, 4.0, 6.0}, []float64{10.0, 20.0}, []float64{20.0, 10.0},
		10.0,
		[]float64{7.0, 5.0, 15.0}, []float64{10.0, 17.0})
}
