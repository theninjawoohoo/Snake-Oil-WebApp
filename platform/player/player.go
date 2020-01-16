package player

//player's hand
type Player struct {
	Cards []string 
	Job string
	Playername string
}

//returns a pointer to a player
func New() *Player {
	return &Player{}
}

//Deals a card to players hand
func (pHand *Player) Deal(word string) {
	pHand.Cards = append(pHand.Cards, word)
}

//Returns a player's hand
func (pHand *Player) GetHand() []string {
	return pHand.Cards
}