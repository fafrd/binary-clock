package main

import (
	"fmt"
	"time"
	"os"
	"sync"
	"github.com/nsf/termbox-go"
)

func main() {
	err := termbox.Init()
	if err != nil {
			panic(err)
	}
	defer os.Exit(0)
	defer termbox.Close()

	var wg sync.WaitGroup
	wg.Add(1)

	fmt.Println("what's the current time?")

	go draw()
	go poll(&wg)

	wg.Wait()
}

func printBCD(printme int, xoffset int) {
	yoffset := 1
	fillRune := 'x'
	/* representing the number 59 in binary-encoded-decimal:
	8    x
	4  x  
	2    x
	1  x x
	______
	   5 9
	*/

	ones := printme % 10
	tens := (printme - ones) / 10

	if tens > 7 {
		termbox.SetCell(xoffset, yoffset, fillRune, termbox.ColorWhite, termbox.ColorBlack)
	} else {
		termbox.SetCell(xoffset, yoffset, ' ', termbox.ColorWhite, termbox.ColorBlack)
	}
	if tens > 3 && tens < 8 {
		termbox.SetCell(xoffset, yoffset + 1, fillRune, termbox.ColorWhite, termbox.ColorBlack)
	} else {
		termbox.SetCell(xoffset, yoffset + 1, ' ', termbox.ColorWhite, termbox.ColorBlack)
	}
	if tens == 2 || tens == 3 || tens == 6 || tens == 7 {
		termbox.SetCell(xoffset, yoffset + 2, fillRune, termbox.ColorWhite, termbox.ColorBlack)
	} else {
		termbox.SetCell(xoffset, yoffset + 2, ' ', termbox.ColorWhite, termbox.ColorBlack)
	}
	if tens % 2 == 1 {
		termbox.SetCell(xoffset, yoffset + 3, fillRune, termbox.ColorWhite, termbox.ColorBlack)
	} else {
		termbox.SetCell(xoffset, yoffset + 3, ' ', termbox.ColorWhite, termbox.ColorBlack)
	}

	xoffset += 2

	if ones > 7 {
		termbox.SetCell(xoffset, yoffset, fillRune, termbox.ColorWhite, termbox.ColorBlack)
	} else {
		termbox.SetCell(xoffset, yoffset, ' ', termbox.ColorWhite, termbox.ColorBlack)
	}
	if ones > 3 && ones < 8 {
		termbox.SetCell(xoffset, yoffset + 1, fillRune, termbox.ColorWhite, termbox.ColorBlack)
	} else {
		termbox.SetCell(xoffset, yoffset + 1, ' ', termbox.ColorWhite, termbox.ColorBlack)
	}
	if ones == 2 || ones == 3 || ones == 6 || ones == 7 {
		termbox.SetCell(xoffset, yoffset + 2, fillRune, termbox.ColorWhite, termbox.ColorBlack)
	} else {
		termbox.SetCell(xoffset, yoffset + 2, ' ', termbox.ColorWhite, termbox.ColorBlack)
	}
	if ones % 2 == 1 {
		termbox.SetCell(xoffset, yoffset + 3, fillRune, termbox.ColorWhite, termbox.ColorBlack)
	} else {
		termbox.SetCell(xoffset, yoffset + 3, ' ', termbox.ColorWhite, termbox.ColorBlack)
	}

	termbox.Flush()
}

func draw() {
	for {
		hour, minute, second := getTime()
		printBCD(hour % 12, 0)
		printBCD(minute, 6)
		printBCD(second, 12)
	}
}

func poll(wg *sync.WaitGroup) {
	for {
		ev := termbox.PollEvent()
		switch ev.Type {
			case termbox.EventKey:
				switch ev.Key {
					case termbox.KeyCtrlQ:
					case termbox.KeyCtrlC:
						wg.Done()
				}
		}
	}
}

func getTime() (int, int, int) {
	t := time.Now()
	hour := t.Hour()
	minute := t.Minute()
	second := t.Second()

	return hour, minute, second
}

