package main

import (
	"os"
	"time"

	"github.com/gdamore/tcell"
)

var HEADER_STYLE tcell.Style = tcell.StyleDefault.Background(tcell.ColorDefault).Foreground(tcell.ColorDefault)
var CURSOR_STYLE tcell.Style = tcell.StyleDefault.Background(tcell.ColorYellow)

func main() {
	screen := initScreen()
	drawLayout(screen)

	for true {
		processEvent(screen)
		time.Sleep(1 * time.Millisecond)
	}
}

func processEvent(screen tcell.Screen) {
	ev := screen.PollEvent()

	if ev == nil {
		return
	}

	switch ev := ev.(type) {
	case *tcell.EventMouse:
		processMouseEvent(screen, ev)
	case *tcell.EventKey:
		processKeyEvent(screen, ev)
	}

}

func processMouseEvent(screen tcell.Screen, ev *tcell.EventMouse) {
	btns := ev.Buttons()
	if btns != tcell.Button1 {
		return
	}

	x, y := ev.Position()

	//we should not allow drawing on the first line
	if y == 0 {
		return
	}

	screen.SetContent(x, y, ' ', nil, CURSOR_STYLE)
	screen.Show()

}

func processKeyEvent(screen tcell.Screen, ev *tcell.EventKey) {
	_, _, ch := ev.Modifiers(), ev.Key(), ev.Rune()
	if ch == 'q' {
		screen.Fini()
		os.Exit(0)
	}
	if ch == 'c' {
		clearLayout(screen)
	}
}

func initScreen() tcell.Screen {
	screen, err := tcell.NewScreen()
	screen.Init()

	if err != nil {
		panic(err)
	}

	screen.EnableMouse()
	return screen
}

func drawLayout(screen tcell.Screen) {
	drawText(screen, 0, 0, HEADER_STYLE, "Q - quit, C - clear")
	screen.Show()
}

func clearLayout(screen tcell.Screen) {
	screen.Clear()
	drawLayout(screen)
	screen.Show()
}

func drawText(s tcell.Screen, x, y int, style tcell.Style, text string) {
	offset := 0
	for _, r := range []rune(text) {
		s.SetContent(x+offset, y, r, nil, style)
		offset++
	}
}
