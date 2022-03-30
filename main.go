package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type DigimonCard struct {
	Name   string `json:"name"`
	Number string `json:"cardnumber"`
}

type SearchedDigimonCard struct {
	Artist        string      `json:"artist"`
	Attribute     string      `json:"attribute"`
	CardSets      []string    `json:"card_sets"`
	Cardnumber    string      `json:"cardnumber"`
	Cardrarity    string      `json:"cardrarity"`
	Color         string      `json:"color"`
	DigiType      string      `json:"digi_type"`
	Dp            int64       `json:"dp"`
	EvolutionCost int64       `json:"evolution_cost"`
	ImageURL      string      `json:"image_url"`
	Level         int64       `json:"level"`
	Maineffect    string      `json:"maineffect"`
	Name          string      `json:"name"`
	PlayCost      int64       `json:"play_cost"`
	SetName       string      `json:"set_name"`
	Soureeffect   interface{} `json:"soureeffect"`
	Stage         string      `json:"stage"`
	Type          string      `json:"type"`
}


const allCardsCurrentGameEndpoint = "https://digimoncard.io/api-public/getAllCards.php?sort=name&series=Digimon%20Card%20Game&sortdirection=asc"
const searchCardEndpoint = "https://digimoncard.io/api-public/search.php?sort=cardnumber&sortdirection=asc&series=Digimon%20Card%20Game"

func main() {
	searchCard()
}

func getAllCards() {
	resp, err := http.Get(allCardsCurrentGameEndpoint)
	if err != nil {
		panic(err)
	}

	var cards []DigimonCard
	err = json.NewDecoder(resp.Body).Decode(&cards)
	if err != nil {
		panic(err)
	}

	for _, card := range cards {
		fmt.Printf("Name: %v || Number: %v \n", card.Name, card.Number)
	}
}

func searchCard() {
	resp, err := http.Get(searchCardEndpoint)
	if err != nil {
		panic(err)
	}

	var cards []SearchedDigimonCard
	err = json.NewDecoder(resp.Body).Decode(&cards)
	if err != nil {
		panic(err)
	}

	for _, card := range cards {
		fmt.Printf("Name: %v || Number: %v \n", card.Name, card.Cardnumber)
	}
}