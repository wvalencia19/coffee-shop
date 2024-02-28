package models

import (
	"time"

	"github.com/wvalencia19/coffee-shop/models/dtos"
)

type Grinder struct {
	id             int
	gramsPerSecond int
	busy           bool
}

type GrinderInterface interface {
	Grind(beans Beans, tc WorkSimulator) Beans
}

func NewGrinder(gramsPerSecond int) *Grinder {
	return &Grinder{
		gramsPerSecond: gramsPerSecond,
		busy:           false,
	}
}

func (g *Grinder) Grind(beans Beans, tc WorkSimulator) Beans {
	g.busy = true
	defer func() { g.busy = false }()
	tc.SimulateWork(time.Second * time.Duration(beans.weightGrams/g.gramsPerSecond))
	return beans
}

func NewGrinderPool(grinders []*Grinder) chan GrinderInterface {
	grinderPool := make(chan GrinderInterface, len(grinders))
	for i, grinder := range grinders {
		grinder.id = i
		grinder.busy = false
		grinderPool <- grinder
	}
	return grinderPool
}

func JSONGrindersTOGrinders(jg []dtos.Grinder) []*Grinder {
	grinders := make([]*Grinder, len(jg))
	for i, grinder := range jg {
		grinders[i] = NewGrinder(grinder.GramsPerSecond)
	}

	return grinders
}
