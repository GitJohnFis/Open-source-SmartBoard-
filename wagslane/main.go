package main

import (
	"time"

	"github.com/GitJohnFis/Open-source-SmartBoard-/internal/pokeapi"
)


func main() {
pokeClient := pokeapi.NewClient(5 * time.Second)
cfg := &config{
	pokeapiClient: pokeClient,
}


	startRepl(cfg);
}