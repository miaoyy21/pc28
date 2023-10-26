package server

func Run(portGold, portBetting string) {
	go gGold(portGold)

	gBetting(portBetting)
}
