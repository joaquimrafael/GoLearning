package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

type Player struct {
	Name        string  `json:"name"`
	Position    string  `json:"position"`
	Age         int     `json:"age"`
	Nationality string  `json:"nationality"`
	Price       float64 `json:"price"`
}

func (p Player) String() string {
	return fmt.Sprintf("Player: %s; Position: %s; Age: %d; Nationality: %s; Price: %.2f",
		p.Name, p.Position, p.Age, p.Nationality, p.Price)
}

func main() {

	if len(os.Args) < 2 {
		fmt.Println("No arguments provided")
		return
	}

	fmt.Println("----------------------------------------------------------------------------------------------------------")
	fmt.Println("Go Player (Json) processor")
	fmt.Println("By Joaquim Prieto 2026")
	fmt.Println()

	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	var players []Player

	if err := json.Unmarshal(data, &players); err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	if len(players) == 0 {
		fmt.Println("No players found")
		return
	}

	sort.Slice(players, func(i, j int) bool {
		return players[i].Price > players[j].Price
	})

	var sum float64
	valuable := players[0]
	for _, p := range players {
		sum += p.Price
		fmt.Println(p)
	}

	fmt.Println()
	fmt.Printf("Most valuable player: %s EUR %.2f m\n", valuable.Name, valuable.Price)
	fmt.Printf("Average player price: EUR %.2f m\n", sum/float64(len(players)))

	fmt.Println()
	fmt.Println("----------------------------------------------------------------------------------------------------------")

}
