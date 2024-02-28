package models

import (
	"time"

	"github.com/wvalencia19/coffee-shop/models/dtos"
)

const (
	// gramsToOuncesRatio represents the conversion ratio from grams of coffee to ounces of water needed.
	gramsToOuncesRatio = 12

	// OuncesMultiplier is the factor used to calculate the final ounces of coffee produced.
	ouncesMultiplier = 6
)

type BrewerInterface interface {
	Brew(beans Beans, tc WorkSimulator) Coffee
}

type Brewer struct {
	id                   int
	ouncesWaterPerSecond int
	busy                 bool
}

func NewBrewer(ouncesWaterPerSecond int) *Brewer {
	return &Brewer{
		ouncesWaterPerSecond: ouncesWaterPerSecond,
		busy:                 false,
	}
}

func (b *Brewer) Brew(beans Beans, tc WorkSimulator) Coffee {
	b.busy = true
	defer func() { b.busy = false }()
	ouncesNeeded := (beans.weightGrams / gramsToOuncesRatio) * ouncesMultiplier
	tc.SimulateWork(time.Second * time.Duration(ouncesNeeded/b.ouncesWaterPerSecond))
	return Coffee{sizeOunces: ouncesNeeded}
}

func NewBrewerPool(brewers []*Brewer) chan BrewerInterface {
	brewerPool := make(chan BrewerInterface, len(brewers))
	for i, brewer := range brewers {
		brewer.id = i
		brewer.busy = false
		brewerPool <- brewer
	}
	return brewerPool
}

func JSONBrewersTOGBrewers(jb []dtos.Brewer) []*Brewer {
	grinders := make([]*Brewer, len(jb))
	for i, brewer := range jb {
		grinders[i] = NewBrewer(brewer.OuncesWaterPerSecond)
	}

	return grinders
}
