/*
undergroundの配列ver
昔、ゲームPGでこんなことやったな版
*/
package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

const limit = 50

var out string
var clear string

func main() {
	for i := 0; i <= limit+10; i++ {
		clear = clear + " "
	}

	rand.Seed(time.Now().UnixNano())
	for {
		for i := 0; i <= limit; i++ {
			num := rand.Intn(10)
			out = out + strconv.Itoa(num)
			switch {
			case num < 1:
				dotReprint()
				fmt.Printf("\r%s", clear)
			case num < 2:
				if i < limit {
					out = out + strconv.Itoa(num)
					i++
				}
			case num >= 9:
				for ii := 0; ii < 3; ii++ {
					if i >= limit {
						break
					}
					out = out + "."
					time.Sleep(time.Duration(num*10) * time.Millisecond)
					i++
				}
			}
			fmt.Printf("\r%s", out) // for sandbox => "fmt.Printf("\x0c%s", out)"
			if i == limit {
				out = ""
				fmt.Println("")
			}
			time.Sleep(time.Duration(num*3) * time.Millisecond)
		}
	}
}

func dotReprint() {
	num := rand.Intn(4)
	for ii := 1; ii <= num; ii++ {
		ticker := time.Tick(time.Second / 2)
		for i := 1; i <= 3; i++ {
			<-ticker
			hoge := ""
			for x := 0; x < i; x++ {
				hoge = hoge + "."
			}
			fmt.Printf("\r%s", clear)
			fmt.Printf("\r%s%s", out, hoge)
		}
	}
}
