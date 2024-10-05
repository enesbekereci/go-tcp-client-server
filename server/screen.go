package main

import (
	"fmt"
	"time"
)

var Update chan string = make(chan string)

func UpdateScreen() {
	for {

		PrintScreen()
		time.Sleep(200 * time.Millisecond)
	}

}

func PrintScreen() {
	PrintHeader()
	mu_cl.Lock()
	PrintTopLine(len(cl))

	fmt.Printf("00  ┃")
	for _, user := range cl {
		if user.state {
			fmt.Printf("\x1B[32m") //green
		} else {
			fmt.Printf("\x1B[31m") //red
		}
		fmt.Print("▓▓▓▓▓")
		fmt.Printf("\x1B[37m")
		fmt.Printf("%20s┃", user.name) //25+1 width
	}

	PrintMidLine(len(cl))

	for _, user := range cl {
		user.logqueue.ResetCurrent()
	}
	for i := 0; i < queuelimit; i++ {
		fmt.Printf("%2d  ┃", i)
		for _, user := range cl {
			if user.logqueue.current != nil {
				fmt.Printf("%25s┃", user.logqueue.Get().message) //  25+1 width
			} else {
				fmt.Print("                         ┃") //25+1 space
			}
		}
		ClearRow()
	}
	PrintBottomLine(len(cl))
	mu_cl.Unlock()
}

func ClearRow() {
	fmt.Print("\033[K") //delete rest of the row
	fmt.Print("\n")
}

func PrintBottomLine(size int) {
	fmt.Print("━━━━┻")
	for i := 1; i < size; i++ {
		fmt.Print("━━━━━━━━━━━━━━━━━━━━━━━━━┻")
	}
	fmt.Print("━━━━━━━━━━━━━━━━━━━━━━━━━┛")
	ClearRow()
}

func PrintTopLine(size int) {
	fmt.Print("━━━━┳")
	for i := 1; i < size; i++ {
		fmt.Print("━━━━━━━━━━━━━━━━━━━━━━━━━┳")
	}
	fmt.Print("━━━━━━━━━━━━━━━━━━━━━━━━━┓\n")
}

func PrintHeader() {
	//fmt.Print("\033[?25l")
	fmt.Print("\033[H") //Move Cursor to the upper left
	fmt.Print("Connection Count : ")
	fmt.Print(GetTotalConn())
	h, m, s := time.Now().Clock()
	//store
	fmt.Printf("          %d:%d:%d", h, m, s)
	ClearRow()
}

func PrintMidLine(size int) {
	fmt.Print("\n━━━━╋")
	for i := 1; i < size; i++ {
		fmt.Print("━━━━━━━━━━━━━━━━━━━━━━━━━╋")
	}
	fmt.Print("━━━━━━━━━━━━━━━━━━━━━━━━━┫")
	ClearRow()
}
