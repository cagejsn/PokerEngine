package main

func isRoyalFlush(sortedFiveCards [5]Card) bool {

	if isFlush(sortedFiveCards) &&
	isStraight(sortedFiveCards) &&
	sortedFiveCards[0].Rank == 10 {
		return true
	}

	return false
}

func isStraightFlush(sortedFiveCards [5]Card) bool {
	return isFlush(sortedFiveCards) && isStraight(sortedFiveCards)

}

func isFourOfAKind(sortedFiveCards [5]Card) bool {
	for i := 0; i < len(sortedFiveCards) - 3 ; i++ {
		benchmarkRank := sortedFiveCards[i].Rank 
		if benchmarkRank == sortedFiveCards[i+1].Rank && benchmarkRank == sortedFiveCards[i+2].Rank && benchmarkRank == sortedFiveCards[i+3].Rank{
			return true	
		}
	}
	
	return false

}

func isFullHouse(sortedFiveCards [5]Card) bool {

	if 	sortedFiveCards[0].Rank == sortedFiveCards[1].Rank &&
		sortedFiveCards[3].Rank == sortedFiveCards[4].Rank &&
		sortedFiveCards[2].Rank == sortedFiveCards[3].Rank || sortedFiveCards[2].Rank == sortedFiveCards[1].Rank {
			return true
		}
	return false

}

func isFlush(sortedFiveCards [5]Card) bool {

	var suit = sortedFiveCards[0].Suit

	for _,v := range sortedFiveCards {
		if v.Suit != suit {
			return false
		}
	}

	return true

}

func isStraight(sortedFiveCards [5]Card) bool {

	var benchmarkRank = sortedFiveCards[0].Rank

	for i := 1; i < len(sortedFiveCards) ; i++ {
		
		if sortedFiveCards[i].Rank != (benchmarkRank + 1) {
			return false;
		}
		benchmarkRank++
	}

	return true
}

func isThreeOfAKind(sortedFiveCards [5]Card) bool {
	

	for i := 0; i < len(sortedFiveCards) - 2 ; i++ {
		benchmarkRank := sortedFiveCards[i].Rank 
		if benchmarkRank == sortedFiveCards[i+1].Rank && benchmarkRank == sortedFiveCards[i+2].Rank{
			return true	
		}
	}
	
	return false

}

func isTwoPair(sortedFiveCards [5]Card) bool {
	
	pairsFound := 0

	for i := 0 ; i < 4 ; {

		if sortedFiveCards[i].Rank == sortedFiveCards[i+1].Rank {
			pairsFound++
			i += 2
		} else {
			i += 1
		}
	}
		
	if pairsFound == 2 {
		return true
	}
	
	return false
	
}

func isPair(sortedFiveCards [5]Card) bool {

	
	for i := 0; i < 4 ; i++ {
		if sortedFiveCards[i].Rank == sortedFiveCards[i+1].Rank {
			return true	
		}
	}
	
	return false
}

func compareFullHouses(firstHand [5]Card, secondHand [5]Card) int {

	// full houses are evaluated by the rank of threeOfAKind and then the pair
	// so the pair only comes into play when the threeOfAKind is the same for both hands
	// the cards are sorted so the middle card will always be a part of the threeOfAKind

	if firstHand[2].Rank > secondHand[2].Rank {
		return 1
	}

	if firstHand[2].Rank < secondHand[2].Rank {
		return 2
	}

	// they are equal, compare the first & last index

	if firstHand[4].Rank > secondHand[4].Rank || 
	firstHand[0].Rank > secondHand[0].Rank {
		return 1
	}

	if firstHand[4].Rank < secondHand[4].Rank || 
	firstHand[0].Rank < secondHand[0].Rank {
		return 2
	} 

	// both are equal
	return 0
}

func comparePairs(pairHands [][5]Card, winningTable *[]bool) {

	tableOfOtherCards := make([][]Card,0,0)
	tableOfPairValues := make([]int,0,0)
	
	for _, fiveCards := range pairHands {

		pairValue, otherCards := findPairFromFiveCard(fiveCards)

		//insert pair value
		tableOfPairValues = append(tableOfPairValues,pairValue)

		
		highCardsFromHand := make([]Card,0,0)
		//other cards is a slice of ints which are the indicies of the
		//card that do not participate in the pair, the indicies are used
		//to grab the actual cards from the fiveCard and then they are passed into a function which 
		//will be used for many situations like this one, where the winner should be resolved by high cards

		for _,v := range otherCards {
			highCardsFromHand = append(highCardsFromHand,fiveCards[v])
		}
		tableOfOtherCards = append(tableOfOtherCards,highCardsFromHand)

	}

	evaluateOtherCards(tableOfOtherCards, winningTable)

}

func evaluateOtherCards(otherCards [][]Card, winningTable *[]bool) {

var numberOfOtherCardsToEvaluate = len(otherCards)
var numberOfCardsInEachOtherCards = len(otherCards[0])
var bestRank int = 0
var workingWinningTable = make([]bool,len(*winningTable))
for currentIndex := numberOfCardsInEachOtherCards - 1 ; currentIndex >= 0 ; currentIndex-- {

	for i := 0; i < numberOfOtherCardsToEvaluate ; i++ {
		if bestRank == 0 {
			bestRank = otherCards[i][currentIndex].Rank
			continue
		}

		if otherCards[i][currentIndex].Rank > bestRank {
			workingWinningTable = make([]bool,len(*winningTable))
			workingWinningTable[i] = true
		}

		if otherCards[i][currentIndex].Rank == bestRank {
			workingWinningTable[i] = true
		}
	}

	//we've made it through a index of high cards, starting with the best one
	//if the winningTable has one entry, then return with that person as the winner, if more than one
	//then keep going through all the cards
	if numberOfWinnersInWinningTable(&workingWinningTable) == 1 {
		return
	}

}

//we have iterated through all the indicies of high cards, if there is still a tie, then 

*winningTable = workingWinningTable

}

func numberOfWinnersInWinningTable(winningTable *[]bool) int {

	var numberOfWinners int

	for _,v := range *winningTable {
		if v == true {
			numberOfWinners++
		}
	}
	return numberOfWinners
}

func findPairFromFiveCard(fiveCard [5]Card) (int , []int) {

	var pairRank int
	var found = false
	indexOfNonPairCards := make([]int,0,0)
	for i := 0 ; i < 5 ; i++ {
		
		if i < 4 && fiveCard[i].Rank == fiveCard[i + 1].Rank && !found {
			found = true
			pairRank = fiveCard[i].Rank
			i = i + 1
		} else {
			indexOfNonPairCards = append(indexOfNonPairCards, i )
		}
	}

	return pairRank, indexOfNonPairCards

}