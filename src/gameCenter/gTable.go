package gameCenter

import (
	"sync"
	"time"

	"github.com/go-vgo/robotgo"
)

type clickModel struct {
	x, y int
	fch  chan interface{}
}

func newClickModel(x, y int) *clickModel {
	return &clickModel{
		x:   x,
		y:   y,
		fch: make(chan interface{}),
	}
}

type gt struct {
	ch        chan *clickModel
	clickOnce *sync.Once
}

func newGTable() *gt {
	return &gt{
		ch:        make(chan *clickModel),
		clickOnce: new(sync.Once),
	}
}

func (g *gt) Click(x, y int) {
	g.clickOnce.Do(func() {
		go func() {
			for m := range g.ch {
				robotgo.MoveClick(m.x, m.y)
				time.Sleep(time.Millisecond*20)
				close(m.fch)
			}
		}()
	})

	cm := newClickModel(x, y)
	g.ch <- cm

	<-cm.fch
}
