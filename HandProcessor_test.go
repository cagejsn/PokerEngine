package main 

import (
	"testing"
)

func TestPairPositive(t *testing.T) {

	fiveCards := [5]Card{ {3,"H"}, {3,"S"}, {4,"H"}, {5,"H"}, {6,"H"} }
	isPair := isPair(fiveCards)
  	if isPair != true {
		t.Error("Expected true, got ", isPair)
 	} 
}

func TestPairNegative(t *testing.T) {
	fiveCards := [5]Card{ {2,"H"}, {3,"S"}, {4,"H"}, {5,"H"}, {6,"H"} }
	isPair := isPair(fiveCards)
	if isPair != false {
	  t.Error("Expected false, got ", isPair)
   } 
}

func TestTwoPairPositive(t *testing.T) {

	fiveCards := [5]Card{ {3,"H"}, {3,"S"}, {4,"H"}, {5,"H"}, {5,"D"} }
	isTwoPair := isTwoPair(fiveCards)
  	if isTwoPair != true {
		t.Error("Expected true, got ", isTwoPair)
 	} 
}

func TestTwoPairNegative(t *testing.T) {
	fiveCards := [5]Card{ {2,"H"}, {3,"S"}, {4,"H"}, {5,"H"}, {6,"H"} }
	isTwoPair := isTwoPair(fiveCards)
	if isTwoPair != false {
	  t.Error("Expected false, got ", isTwoPair)
   } 
}

func TestFlushPositive(t *testing.T) {

	fiveCards := [5]Card{ {2,"H"}, {4,"H"}, {5,"H"}, {6,"H"}, {9,"H"} }
	isFlush := isFlush(fiveCards)
  	if isFlush != true {
		t.Error("Expected true, got ", isFlush)
 	} 
}

func TestFlushNegative(t *testing.T) {

	fiveCards := [5]Card{ {3,"H"}, {3,"S"}, {4,"H"}, {5,"H"}, {6,"H"} }
	isFlush := isFlush(fiveCards)
  	if isFlush != false {
		t.Error("Expected false, got ", isFlush)
 	} 
}


func TestStraightPositive(t *testing.T) {
	fiveCards := [5]Card{ {2,"H"}, {3,"S"}, {4,"H"}, {5,"H"}, {6,"H"} }
	isStraight := isStraight(fiveCards)
  	if isStraight != true {
		t.Error("Expected true, got ", isStraight)
 	} 
}

func TestStraightNegative(t *testing.T) {
	fiveCards := [5]Card{ {3,"S"}, {4,"H"}, {5,"H"}, {6,"H"}, {14,"H"} }
	isStraight := isStraight(fiveCards)
  	if isStraight != false {
		t.Error("Expected false, got ", isStraight)
 	} 
}

func TestStraightFlushPositive(t *testing.T) {
	fiveCards := [5]Card{ {2,"H"}, {3,"H"}, {4,"H"}, {5,"H"}, {6,"H"} }
	isStraightFlush := isStraightFlush(fiveCards)
  	if isStraightFlush != true {
		t.Error("Expected true, got ", isStraightFlush)
 	} 
}

func TestStraightFlushNegative(t *testing.T) {
	fiveCards := [5]Card{ {3,"H"}, {4,"S"}, {5,"H"}, {6,"H"}, {14,"H"} }
	isStraightFlush := isStraightFlush(fiveCards)
  	if isStraightFlush != false {
		t.Error("Expected false, got ", isStraightFlush)
 	} 
}

func TestThreeOfAKindPositive(t *testing.T) {
	fiveCards := [5]Card{ {3,"S"}, {4,"H"}, {4,"S"}, {4,"D"} , {14,"H"}}
	isThreeOfAKind := isThreeOfAKind(fiveCards)
  	if isThreeOfAKind != true {
		t.Error("Expected true, got ", isThreeOfAKind)
 	} 
}

func TestThreeOfAKindNegative(t *testing.T) {
	fiveCards := [5]Card{ {3,"S"}, {4,"H"}, {4,"S"}, {6,"H"}, {14,"H"}}
	isThreeOfAKind := isThreeOfAKind(fiveCards)
  	if isThreeOfAKind != false {
		t.Error("Expected false, got ", isThreeOfAKind)
 	} 
}

func TestFourOfAKindPositive(t *testing.T) {
	fiveCards := [5]Card{ {2,"H"}, {2,"S"}, {2,"D"}, {2,"C"}, {6,"H"} }
	isFourOfAKind := isFourOfAKind(fiveCards)
  	if isFourOfAKind != true {
		t.Error("Expected true, got ", isFourOfAKind)
 	} 
}

func TestFourOfAKindNegative(t *testing.T) {
	fiveCards := [5]Card{  {3,"S"}, {4,"H"}, {5,"H"}, {6,"H"}, {14,"H"}}
	isFourOfAKind := isFourOfAKind(fiveCards)
  	if isFourOfAKind != false {
		t.Error("Expected false, got ", isFourOfAKind)
 	} 
}

func TestFullHouseUpperPositive(t *testing.T) {
	fiveCards := [5]Card{ {3,"H"}, {3,"S"}, {7,"H"}, {7,"D"}, {7,"C"} }
	isFullHouse := isFullHouse(fiveCards)
  	if isFullHouse != true {
		t.Error("Expected true, got ", isFullHouse)
 	} 
}

func TestFullHouseLowerPositive(t *testing.T) {
	fiveCards := [5]Card{ {3,"H"}, {3,"S"}, {3,"H"}, {7,"D"}, {7,"C"} }
	isFullHouse := isFullHouse(fiveCards)
  	if isFullHouse != true {
		t.Error("Expected true, got ", isFullHouse)
 	} 
}

func TestFullHouseNegative(t *testing.T) {
	fiveCards := [5]Card{ {2,"H"}, {3,"S"}, {4,"H"}, {5,"H"}, {6,"H"} }
	isFullHouse := isFullHouse(fiveCards)
  	if isFullHouse != false {
		t.Error("Expected false, got ", isFullHouse)
 	} 
}


func TestRoyalFlushPositive(t *testing.T) {

	fiveCards := [5]Card{ {10,"H"}, {11,"H"}, {12,"H"}, {13,"H"}, {14,"H"}, }
	isRoyalFlush := isRoyalFlush(fiveCards)
  	if isRoyalFlush != true {
		t.Error("Expected true, got ", isRoyalFlush)
 	} 
}

func TestRoyalFlushNegative(t *testing.T) {
	fiveCards := [5]Card{ {2,"H"}, {3,"S"}, {4,"H"}, {5,"H"}, {6,"H"} }
	isRoyalFlush := isRoyalFlush(fiveCards)
	if isRoyalFlush != false {
	  t.Error("Expected false, got ", isRoyalFlush)
   } 
}


func TestTwoFullHouses_FirstIsBetter_HigherThreeCard(t *testing.T){
	firstFiveCards := [5]Card{ {2,"H"}, {2,"S"}, {5,"H"}, {5,"S"}, {5,"D"} }
	secondFiveCards := [5]Card{ {2,"H"}, {2,"S"}, {2,"D"}, {5,"S"}, {5,"D"} }

	results := compareFullHouses(firstFiveCards,secondFiveCards)
	
	if results != 1 {
	  t.Error("Expected first full house to win, got ", results)
   }
}

func TestTwoFullHouses_FirstIsBetter_HigherThreeButInLowerIndex(t *testing.T){
	firstFiveCards := [5]Card{ {7,"H"}, {7,"S"}, {7,"D"}, {8,"S"}, {8,"D"} }
	secondFiveCards := [5]Card{ {2,"H"}, {2,"S"}, {2,"D"}, {7,"S"}, {7,"D"} }

	results := compareFullHouses(firstFiveCards,secondFiveCards)
	
	if results != 1 {
	  t.Error("Expected first full house to win, got ", results)
   }
}

func TestTwoFullHouses_SecondIsBetter_HigherTwoCards(t *testing.T){
	firstFiveCards := [5]Card{ {2,"H"}, {2,"S"}, {4,"H"}, {4,"S"}, {4,"D"} }
	secondFiveCards := [5]Card{ {3,"H"}, {3,"S"}, {4,"H"}, {4,"S"}, {4,"D"} }

	results := compareFullHouses(firstFiveCards,secondFiveCards)
	
	if results != 2 {
	  t.Error("Expected second full house to win, got ", results)
   }
}

func TestTwoFullHouses_SplitPot(t *testing.T){
	firstFiveCards := [5]Card{ {2,"H"}, {2,"S"}, {4,"H"}, {4,"S"}, {4,"D"} }
	secondFiveCards := [5]Card{ {2,"D"}, {2,"C"}, {4,"H"}, {4,"S"}, {4,"D"} }

	results := compareFullHouses(firstFiveCards,secondFiveCards)
	
	if results != 0 {
	  t.Error("Expected tie, got ", results)
   }
}

func TestComparePairs_SecondWins(t *testing.T){

	firstFiveCards := [5]Card{ {2,"D"}, {2,"C"}, {4,"H"}, {8,"S"}, {8,"D"} }
	secondFiveCards := [5]Card{ {2,"H"}, {3,"S"}, {7,"H"}, {12,"S"}, {12,"D"} }
	thirdFiveCards := [5]Card{ {2,"D"}, {3,"C"}, {4,"H"}, {9,"S"}, {9,"D"} }

	var hands = [][5]Card{firstFiveCards,secondFiveCards,thirdFiveCards}

	winners := make([]bool,3,3)

	comparePairs(hands,&winners)
	
	if winners[0] != false && winners[1] != true && winners[2] != false  {
	  t.Error("Expected second to win, got ", winners)
   }

}

func TestComparePairs_FirstWinsWithKicker(t *testing.T){

	firstFiveCards := [5]Card{ {4,"D"}, {5,"C"}, {8,"H"}, {8,"D"}, {14,"D"} }
	secondFiveCards := [5]Card{ {4,"D"}, {5,"C"}, {7,"H"}, {8,"S"}, {8,"D"} }
	thirdFiveCards := [5]Card{ {4,"D"}, {5,"C"}, {6,"H"}, {6,"S"}, {8,"D"} }

	var hands = [][5]Card{firstFiveCards,secondFiveCards,thirdFiveCards}

	winners := make([]bool,3,3)

	comparePairs(hands,&winners)
	
	if winners[0] != true && winners[1] != false && winners[2] != false {
	  t.Error("Expected first to win, got ", winners)
   }
}

func TestComparePairs_SplitPot_TwoWays(t *testing.T){

	firstFiveCards := [5]Card{ {4,"D"}, {5,"C"}, {8,"S"}, {8,"D"}, {14,"D"} }
	secondFiveCards := [5]Card{ {4,"D"}, {5,"C"}, {8,"C"}, {8,"D"}, {14,"D"}  }
	thirdFiveCards := [5]Card{ {4,"D"}, {5,"C"}, {6,"H"}, {6,"S"}, {8,"D"} }

	var hands = [][5]Card{firstFiveCards,secondFiveCards,thirdFiveCards}
	winners := make([]bool,3,3)

	comparePairs(hands,&winners)
	
	if winners[0] != true && winners[1] != true && winners[2] != false {
	  t.Error("Expected first and second to split pot, got ", winners)
   }
}

func TestComparePairs_SplitPot_ThreeWays(t *testing.T){

	firstFiveCards := [5]Card{ {4,"D"}, {5,"C"}, {8,"S"}, {8,"D"}, {14,"D"} }
	secondFiveCards := [5]Card{ {4,"D"}, {5,"C"}, {8,"C"}, {8,"D"}, {14,"D"}  }
	thirdFiveCards := [5]Card{ {4,"D"}, {5,"C"}, {8,"H"}, {8,"D"}, {14,"D"} }

	var hands = [][5]Card{firstFiveCards,secondFiveCards,thirdFiveCards}
	winners := make([]bool,3,3)

	comparePairs(hands,&winners)
	
	if  winners[0] != true && winners[1] != true && winners[2] != true {
	  t.Error("Expected three way tie, got ", winners)
   }
}

func TestComparePairs_SplitPot_ThreeWaysOutOfFour(t *testing.T){

	firstFiveCards := [5]Card{ {4,"D"}, {5,"C"}, {8,"S"}, {8,"D"}, {14,"D"} }
	secondFiveCards := [5]Card{ {4,"D"}, {5,"C"}, {8,"C"}, {8,"D"}, {14,"D"}  }
	thirdFiveCards := [5]Card{ {4,"D"}, {5,"C"}, {8,"H"}, {8,"D"}, {13,"D"} }
	fourthFiveCards := [5]Card{ {4,"D"}, {5,"C"}, {8,"H"}, {8,"D"}, {14,"D"} }

	var hands = [][5]Card{firstFiveCards,secondFiveCards,thirdFiveCards,fourthFiveCards}
	winners := make([]bool,4,4)

	comparePairs(hands,&winners)
	
	if  winners[0] != true && winners[1] != true && winners[2] != false && winners[3] != true {
	  t.Error("Expected three way tie, got ", winners)
   }
}

func TestFindPairInHand(t *testing.T){
	fiveCards := [5]Card{ {4,"D"}, {4,"C"}, {8,"S"}, {9,"D"}, {14,"D"} }

	rankOfPair, otherCardSlice := findPairFromFiveCard(fiveCards)
	if rankOfPair != 4 {
		t.Error("Expected 4 as pair rank, got", rankOfPair)
	}

	if otherCardSlice[0] != 2 && 
	otherCardSlice[1] != 3 && 
	otherCardSlice[2] != 4  {
		t.Error("Expected 2, 3, 4  as index of non pair cards , got", otherCardSlice)
	}
}

