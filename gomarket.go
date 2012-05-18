
package gomarket

import (
	"fmt"
	"sort"
)

type Resource interface {}

type Market struct {
	Actors map[Actor]bool
	Prices map[Resource]float64
}
func NewMarket() *Market {
	return &Market{make(map[Actor]bool), make(map[Resource]float64)}
}
func (m *Market) tradeResource(asks, bids []*Order) float64 {
	satisfied_bids := make(map[*Order]*Order)
	last_ask_price, last_bid_price := 0.0, 0.0
	match_exists := true
	for len(asks) > 0 && len(bids) > 0 && match_exists {
		bid := bids[0]
		ask := asks[len(asks) - 1]
		last_bid_price = bid.Price
		last_ask_price = ask.Price
		if bid.Price >= ask.Price {
			if ask.Units > bid.Units {
				partial_ask := &Order{bid.Units, ask.Resource, ask.Price, ask.Actor}
				satisfied_bids[bid] = partial_ask
				bids = bids[1:]
				ask.Units = ask.Units - bid.Units
			} else if ask.Units < bid.Units {
				partial_bid := &Order{ask.Units, bid.Resource, bid.Price, bid.Actor}
				satisfied_bids[partial_bid] = ask
				asks = asks[:len(asks) - 1]
				bid.Units = bid.Units - ask.Units
			} else {
				satisfied_bids[bid] = ask
				asks = asks[:len(asks) - 1]
				bids = bids[1:]
			}
		} else {
			match_exists = false
		}
	}
	actual_price := 0.0
	if len(satisfied_bids) > 0 {
		if len(asks) == 0 && len(bids) == 0 {
			actual_price = (last_ask_price + last_bid_price) / 2.0
		} else if len(asks) == 0 {
			actual_price = last_bid_price
		} else if len(bids) == 0 {
			actual_price = last_ask_price
		} else {
			actual_price = (last_ask_price + last_bid_price) / 2.0
		}
	} else {
		actual_price = (last_ask_price + last_bid_price) / 2.0
	}
	for bid, ask := range satisfied_bids {
		bid.Actor.Buy(bid.Units, bid.Resource, actual_price, ask.Actor)
	}
	return actual_price
}
func (m *Market) Trade() {
	all_asks, all_bids, ask_sums, bid_sums, resources := m.createSums()
	for resource,_ := range resources {
		bids := all_bids[resource]
		asks := all_asks[resource]
		sort.Sort(Orders(asks))
		sort.Sort(Orders(bids))
		if ask_sums[resource] == 0 {
			m.Prices[resource] = bids[0].Price
		} else if bid_sums[resource] == 0 {
			m.Prices[resource] = asks[len(asks) - 1].Price
		} else {
			m.Prices[resource] = m.tradeResource(asks, bids)
		}
	}
}
func (m *Market) createSums() (
	asks, bids map[Resource][]*Order, 
	ask_sums, bid_sums map[Resource]float64, 
	resources map[Resource]bool) {
	
	asks = make(map[Resource][]*Order)
	bids = make(map[Resource][]*Order)
	resources = make(map[Resource]bool)
	ask_sums = make(map[Resource]float64)
	bid_sums = make(map[Resource]float64)
	for actor,_ := range m.Actors {
		for ask,_ := range actor.Asks() {
			asks[ask.Resource] = append(asks[ask.Resource], ask)
			ask_sums[ask.Resource] += ask.Units
			resources[ask.Resource] = true
		}
		for bid,_ := range actor.Bids() {
			bids[bid.Resource] = append(bids[bid.Resource], bid)
			bid_sums[bid.Resource] += bid.Units
			resources[bid.Resource] = true
		}
	}
	return
}


type Order struct {
	Units float64
	Resource Resource
	Price float64
	Actor Actor
}
func (o *Order) String() string {
	return fmt.Sprint(o.Actor, ":", o.Resource, ":", o.Units, "*", o.Price)
}


type Orders []*Order
func (o Orders) Swap(i, j int) {
	o[i], o[j] = o[j], o[i]
}
func (o Orders) Len() int {
	return len(o)
}
func (o Orders) Less(i,j int) bool {
	return o[i].Price > o[j].Price
}
