package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

const limit = 50

func main() {
	rand.Seed(time.Now().UnixNano())
	for {
		for i := 0; i <= limit; i++ {
			num := rand.Intn(10)
			fmt.Print(strconv.Itoa(num))
			switch {
			case num < 2:
				if i < limit {
					fmt.Print(strconv.Itoa(num))
					i++
				}
			case num >= 9:
				for ii := 0; ii < 3; ii++ {
					if i >= limit {
						break
					}
					fmt.Print(".")
					time.Sleep(time.Duration(num*10) * time.Millisecond)
					i++
				}
			}
			if i == limit {
				fmt.Println("")
			}
			time.Sleep(time.Duration(num*3) * time.Millisecond)
		}
	}
}
