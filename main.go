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

func printBCD(printme int, xoffset int, drawHourCorner bool) {
	yoffset := 1
	fillRune := 'x'
	fg := termbox.ColorWhite
	bg := termbox.ColorBlack
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

	// draw box
	drawBox(drawHourCorner, xoffset, yoffset, fg, bg)

	// draw tens digit
	if tens > 7 {
		termbox.SetCell(xoffset + 1, yoffset + 1, fillRune, termbox.ColorWhite, termbox.ColorBlack)
	}
	if tens > 3 && tens < 8 {
		termbox.SetCell(xoffset + 1, yoffset + 3, fillRune, termbox.ColorWhite, termbox.ColorBlack)
	}
	if tens == 2 || tens == 3 || tens == 6 || tens == 7 {
		termbox.SetCell(xoffset + 1, yoffset + 5, fillRune, termbox.ColorWhite, termbox.ColorBlack)
	}
	if tens % 2 == 1 {
		termbox.SetCell(xoffset + 1, yoffset + 7, fillRune, termbox.ColorWhite, termbox.ColorBlack)
	}

	// draw ones digit
	if ones > 7 {
		termbox.SetCell(xoffset + 3, yoffset + 1, fillRune, termbox.ColorWhite, termbox.ColorBlack)
	}
	if ones > 3 && ones < 8 {
		termbox.SetCell(xoffset + 3, yoffset + 3, fillRune, termbox.ColorWhite, termbox.ColorBlack)
	}
	if ones == 2 || ones == 3 || ones == 6 || ones == 7 {
		termbox.SetCell(xoffset + 3, yoffset + 5, fillRune, termbox.ColorWhite, termbox.ColorBlack)
	}
	if ones % 2 == 1 {
		termbox.SetCell(xoffset + 3, yoffset + 7, fillRune, termbox.ColorWhite, termbox.ColorBlack)
	}

	termbox.Flush()
}

func draw() {
	for {
		hour, minute, second := getTime()
		printBCD(hour, 0, true)
		printBCD(minute, 8, false)
		printBCD(second, 16, false)
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

func drawBox(hour bool, xoffset int, yoffset int, fg termbox.Attribute, bg termbox.Attribute) {
	if (hour) {
		termbox.SetCell(xoffset, yoffset + 4, '┌', fg, bg)
		termbox.SetCell(xoffset + 2, yoffset + 2, '├', fg, bg)
	} else {
		termbox.SetCell(xoffset, yoffset + 2, '┌', fg, bg)
		termbox.SetCell(xoffset, yoffset + 3, '│', fg, bg)
		termbox.SetCell(xoffset, yoffset + 4, '├', fg, bg)

		termbox.SetCell(xoffset + 1, yoffset + 2, '─', fg, bg)
		termbox.SetCell(xoffset + 1, yoffset + 3, ' ', fg, bg)

		termbox.SetCell(xoffset + 2, yoffset + 2, '┼', fg, bg)
	}

	termbox.SetCell(xoffset, yoffset + 5, '│', fg, bg)
	termbox.SetCell(xoffset, yoffset + 6, '├', fg, bg)
	termbox.SetCell(xoffset, yoffset + 7, '│', fg, bg)
	termbox.SetCell(xoffset, yoffset + 8, '└', fg, bg)

	termbox.SetCell(xoffset + 1, yoffset + 4, '─', fg, bg)
	termbox.SetCell(xoffset + 1, yoffset + 5, ' ', fg, bg)
	termbox.SetCell(xoffset + 1, yoffset + 6, '─', fg, bg)
	termbox.SetCell(xoffset + 1, yoffset + 7, ' ', fg, bg)
	termbox.SetCell(xoffset + 1, yoffset + 8, '─', fg, bg)

	termbox.SetCell(xoffset + 2, yoffset, '┌', fg, bg)
	termbox.SetCell(xoffset + 2, yoffset + 1, '│', fg, bg)
	termbox.SetCell(xoffset + 2, yoffset + 3, '│', fg, bg)
	termbox.SetCell(xoffset + 2, yoffset + 4, '┼', fg, bg)
	termbox.SetCell(xoffset + 2, yoffset + 5, '│', fg, bg)
	termbox.SetCell(xoffset + 2, yoffset + 6, '┼', fg, bg)
	termbox.SetCell(xoffset + 2, yoffset + 7, '│', fg, bg)
	termbox.SetCell(xoffset + 2, yoffset + 8, '┴', fg, bg)

	termbox.SetCell(xoffset + 3, yoffset, '─', fg, bg)
	termbox.SetCell(xoffset + 3, yoffset + 1, ' ', fg, bg)
	termbox.SetCell(xoffset + 3, yoffset + 2, '─', fg, bg)
	termbox.SetCell(xoffset + 3, yoffset + 3, ' ', fg, bg)
	termbox.SetCell(xoffset + 3, yoffset + 4, '─', fg, bg)
	termbox.SetCell(xoffset + 3, yoffset + 5, ' ', fg, bg)
	termbox.SetCell(xoffset + 3, yoffset + 6, '─', fg, bg)
	termbox.SetCell(xoffset + 3, yoffset + 7, ' ', fg, bg)
	termbox.SetCell(xoffset + 3, yoffset + 8, '─', fg, bg)

	termbox.SetCell(xoffset + 4, yoffset, '┐', fg, bg)
	termbox.SetCell(xoffset + 4, yoffset + 1, '│', fg, bg)
	termbox.SetCell(xoffset + 4, yoffset + 2, '┤', fg, bg)
	termbox.SetCell(xoffset + 4, yoffset + 3, '│', fg, bg)
	termbox.SetCell(xoffset + 4, yoffset + 4, '┤', fg, bg)
	termbox.SetCell(xoffset + 4, yoffset + 5, '│', fg, bg)
	termbox.SetCell(xoffset + 4, yoffset + 6, '┤', fg, bg)
	termbox.SetCell(xoffset + 4, yoffset + 7, '│', fg, bg)
	termbox.SetCell(xoffset + 4, yoffset + 8, '┘', fg, bg)
}

