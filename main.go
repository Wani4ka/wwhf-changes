package main

import (
	"fmt"
	"math"
)

type Item struct {
	Id   int
	Name string
}

func mi(id int, name string) Item {
	return Item{Id: id, Name: name}
}

func (i Item) String() string {
	return i.Name
}

func (i Item) Mask() int {
	return 1 << i.Id
}

type Swap struct {
	fromMe  Item
	fromHim Item
}

func (s Swap) String() string {
	return fmt.Sprintf("Wani4ka gives %s, akurse gives %s", s.fromMe, s.fromHim)
}

var (
	SHIELD = mi(0, "shield")
	SWORD  = mi(1, "sword")
	CLOCK  = mi(2, "clock")
	JESTER = mi(3, "jester")
	KING   = mi(4, "king")
	CHAIN  = mi(5, "chain")
	BOOK   = mi(6, "book")
	FLASK  = mi(7, "flask")
)

var items = []Item{SHIELD, SWORD, CLOCK, JESTER, KING, CHAIN, BOOK, FLASK}

var required = SHIELD.Mask() | SWORD.Mask() | CLOCK.Mask() | KING.Mask()

var known map[int]bool

var mostOptimal []Swap
var mostOptimalLength = math.MaxInt32

func search(current int, steps *[]Swap) {
	if current == required {
		if len(*steps) < mostOptimalLength {
			mostOptimal = *steps
			mostOptimalLength = len(mostOptimal)
		}
		return
	}
	for i := 0; i < len(items); i++ {
		for j := i - 1; j < i+2; j++ {
			if i == j || j < 0 || j >= len(items) || (current&items[i].Mask()) == (current&items[j].Mask()) {
				continue
			}
			nxt := (current ^ items[i].Mask()) ^ items[j].Mask()
			_, was := known[nxt]
			if was {
				continue
			}
			known[nxt] = true
			if (current & items[i].Mask()) > 0 {
				*steps = append(*steps, Swap{fromMe: items[i], fromHim: items[j]})
			} else {
				*steps = append(*steps, Swap{fromMe: items[j], fromHim: items[i]})
			}
			search(nxt, steps)
			*steps = (*steps)[:len(*steps)-1]
		}
	}
}

func main() {
	known = make(map[int]bool)
	steps := make([]Swap, 0)
	search(KING.Mask()|BOOK.Mask()|JESTER.Mask()|CLOCK.Mask(), &steps)
	for i, step := range mostOptimal {
		fmt.Printf("%d. %s\n", i+1, step)
	}
}
