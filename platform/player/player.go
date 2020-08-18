package player

import (
	"math/rand"
	"time"
)

//player's hand
type Player struct {
	Cards      []string
	Job        string
	PlayerName string
	IsJudge    bool
	Cookie     string
}

//returns a pointer to a player
func New() *Player {
	return &Player{}
}

//Deals cards to players hand
func (player *Player) Deal(words []string) {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(words), func(i, j int) { words[i], words[j] = words[j], words[i] })
	for i := 0; i < 6; i++ {
		player.Cards[i] = words[i]
	}
}

//Returns a player's hand
func (player *Player) GetHand() []string {
	return player.Cards
}
