package main

import "fmt"

func main() {
	playingDomino([][]int{{6, 5}, {3, 4}, {2, 1}, {3, 3}}, []int{4, 3})
	playingDomino([][]int{{6, 5}, {3, 3}, {3, 4}, {2, 1}}, []int{3, 6})
	playingDomino([][]int{{6, 6}, {2, 4}, {3, 6}}, []int{5, 1})
}

func playingDomino(cards [][]int, deck []int) interface{} {
	indxBiggestPair := -1
	biggestPair := 0
	for i := 0; i < len(cards); i++ {
		if deck[0] == cards[i][0] && deck[0] > biggestPair ||
			deck[0] == cards[i][1] && deck[0] > biggestPair {
			biggestPair = deck[0]
			indxBiggestPair = i
		}
		if deck[1] == cards[i][0] && deck[1] > biggestPair ||
			deck[1] == cards[i][1] && deck[1] > biggestPair {
			biggestPair = deck[1]
			indxBiggestPair = i
		}
	}
	if indxBiggestPair == -1 {
		fmt.Println("tutup kartu")
		return []int{0, 0}
	}
	fmt.Println(cards[indxBiggestPair])
	return cards[indxBiggestPair]
}
