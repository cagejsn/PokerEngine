package main 

import (
	"testing"
)



func TestFindBestHandForPlayer_FourOfAKind(t *testing.T) {

	communityCards := [5]Card{ {3,"H"}, {3,"S"}, {4,"H"}, {5,"H"}, {5,"D"} }
	hand := Hand{ {3,"D"}, {3,"C"} }

	bestHand := findBestHandForPlayer(communityCards, hand)
  	if bestHand[0] != communityCards[0] {
		t.Error("Expected 3 of H, S, D, C and 5, got ", bestHand)
 	} 
}

