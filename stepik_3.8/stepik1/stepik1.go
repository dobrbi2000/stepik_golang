package main

import (
	"fmt"
	"time"
)

// func finder1(mine [5]string) {
// 	for _, item := range mine {
// 		if item == "ore" {
// 			fmt.Println("The Finder1: Found ore!")
// 		}
// 	}
// }

// func finder2(mine [5]string) {
// 	for _, item := range mine {
// 		if item == "ore" {
// 			fmt.Println("The Finder2: Found ore!")
// 		}
// 	}
// }

func main() {

	theMine := [5]string{"rock", "ore", "ore", "rock", "ore"}

	oreChannel := make(chan string)
	minedOreChan := make(chan string)

	//Разведчик
	go func(mine [5]string) {
		for _, item := range mine {
			if item == "ore" {
				oreChannel <- item //передаем в oreChannel
			}
		}
	}(theMine)

	//Добытчик
	go func() {
		for i := 0; i < 3; i++ {
			foundOre := <-oreChannel // чтение из канала oreChannel
			fmt.Println("From Finder: ", foundOre)
			minedOreChan <- "minedOre" // передаем данные в minedOreChan
		}
	}()

	//Переработчик
	go func() {
		for i := 0; i < 3; i++ {
			minedOre := <-minedOreChan // чтение данных из minedOreChan
			fmt.Println("From Miner ", minedOre)
			fmt.Println("From Smelter: Ore is smelted")
		}
	}()

	<-time.After(5 * time.Second)

}
