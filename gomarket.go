
package gomarket

import (
	"fmt"
)

type Market struct {
	Actors map[Actor]bool
}

func (m *Market) Trade() {
	asks := make(map[Resource][]*Ask)
	bids := make(map[Resource][]*Bid)
	for actor,_ := range m.Actors {
		for ask,_ := range actor.Asks() {
			asks[ask.Resource] = append(asks[ask.Resource], ask)
		}
		for bid,_ := range actor.Bids() {
			bids[bid.Resource] = append(bids[bid.Resource], bid)
		}
	}
	fmt.Println("asks:", asks)
	fmt.Println("bids:", bids)
}

type Resource interface {
	GetName() string
}

type Actor interface {
	Asks() map[*Ask]bool
	Bids() map[*Bid]bool
}

type Order struct {
	Units float32
	Resource Resource
	Actor Actor
}
func (o *Order) String() string {
	return fmt.Sprint(o.Units, "*", o.Resource, "@", o.Actor)
}

type Ask struct {
	*Order
	MinimumPrice float32
}
func (a *Ask) String() string {
	return fmt.Sprint(a.Order.String(), ">", a.MinimumPrice)
}

type Bid struct {
	*Order
	MaximumPrice float32
}
func (b *Bid) String() string {
	return fmt.Sprint(b.Order.String(), "<", b.MaximumPrice)
}

type Consumable struct {
	Name string
}
func (c *Consumable) GetName() string {
	return c.Name;
}
func (c *Consumable) String() string {
	return c.Name
}


type Shoes struct {}
func (s *Shoes) Name() string {
	return "shoes"
}

type Rice struct {}
func (r *Rice) Name() string {
	return "rice"
}

type Human struct {
	name string
	asks map[*Ask]bool
	bids map[*Bid]bool
}
func (h *Human) String() string {
	return h.name
}
func (h *Human) Ask(u float32, r Resource, p float32) {
	h.asks[&Ask{&Order{u, r, h}, p}] = true
}
func (h *Human) Bid(u float32, r Resource, p float32) {
	h.bids[&Bid{&Order{u, r, h}, p}] = true
}
func (h *Human) Asks() map[*Ask]bool {
	return h.asks
}
func (h *Human) Bids() map[*Bid]bool {
	return h.bids
}

