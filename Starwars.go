package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)

	for i := 0; i < 10; i++ {
		go sleepyStormtrooper(i, c)
	}

	darthVaderPatience := time.After(7 * time.Second)

	for i := 0; i < 10; i++ {
		select {
		case stormtrooper := <-c:
			fmt.Printf("El Stormtrooper, numero %d, esta despierto!\n", stormtrooper)
		case <-darthVaderPatience:
			fmt.Println("El imperio no tiene cavidad para gente floja, por tu insubirdinacion moriras *le explota el huevo*!")
			return
		}
	}

}

func sleepyStormtrooper(id int, c chan int) {
	time.Sleep(7 * time.Second)
	fmt.Printf("El StormTrooper, numero %d, esta dormido\n", id)
	c <- id
}