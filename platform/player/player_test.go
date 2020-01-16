package player

//These are unit tests for the player's cards
import "testing"

func TestDeal(t *testing.T) {
	hand := New()
	hand.Deal("Apple")

	if len(hand.Cards) == 0 {
		t.Errorf("Card was not dealt")
	}
}

func TestGetHand(t *testing.T) {
	hand := New()
	hand.Deal("Apple")
	result := hand.GetHand()
	if len(result) != 1 {
		t.Errorf("Results doesn't equal one")
	}
}