package main

func canPlaceFlowers(flowerbed []int, n int) bool {
	var (
		emptySpace int8 = 0
		plantedFl  int  = 0
	)
	for i, flb := range flowerbed {
		if i == 0 {
			emptySpace++
		}
		if i == len(flowerbed)-1 {
			emptySpace++
		}
		if flb == 0 && emptySpace < 3 {
			emptySpace++
		}
		if flb == 1 {
			emptySpace = 0
		}
		if emptySpace == 3 {
			plantedFl++
			emptySpace = 1
		}
	}
	return plantedFl >= n
}
