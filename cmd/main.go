package main

import (
	"fmt"
	"log"

	"github.com/MarioCarrion/loteria"
)

func main() {
	caller := loteria.NewCaller()
	p1, err := caller.AddPlayer("mario")
	if err != nil {
		log.Fatalf("adding p1 failed %s", err)
	}
	p2, err := caller.AddPlayer("bing")
	if err != nil {
		log.Fatalf("adding p2 failed %s", err)
	}

	players := [2]*loteria.Player{&p1, &p2}

AnnoucingLoop:
	for {
		card, err := caller.Announce()
		if err != nil {
			fmt.Printf("Got error %s!\n", err)
			break
		}
		fmt.Print(".")
		// fmt.Printf("> Caller: %s!\n", card)
		for _, player := range players {
			if err := player.Mark(card); err != nil {
				continue
			}
			// fmt.Printf("\t> %s got it!\n", player.Name())

			if !player.IsWinner() {
				continue
			}

			if err := caller.Loteria(player.Name()); err == nil {
				fmt.Printf("\n%s WON!!\n", player.Name())
				break AnnoucingLoop
			}
			// fmt.Printf("\n")
		}
	}
}
