package main

import (
	"time"

	tm "github.com/buger/goterm"
)

func main() {

	text := `

      ██████╗ ██████╗ ███████╗███████╗███████╗███████╗    ██████╗ ██████╗ ███████╗ █████╗ ██╗  ██╗
     ██╔════╝██╔═══██╗██╔════╝██╔════╝██╔════╝██╔════╝    ██╔══██╗██╔══██╗██╔════╝██╔══██╗██║ ██╔╝
     ██║     ██║   ██║█████╗  █████╗  █████╗  █████╗      ██████╔╝██████╔╝█████╗  ███████║█████╔╝
     ██║     ██║   ██║██╔══╝  ██╔══╝  ██╔══╝  ██╔══╝      ██╔══██╗██╔══██╗██╔══╝  ██╔══██║██╔═██╗
     ╚██████╗╚██████╔╝██║     ██║     ███████╗███████╗    ██████╔╝██║  ██║███████╗██║  ██║██║  ██╗
      ╚═════╝ ╚═════╝ ╚═╝     ╚═╝     ╚══════╝╚══════╝    ╚═════╝ ╚═╝  ╚═╝╚══════╝╚═╝  ╚═╝╚═╝  ╚═╝


`
	brew := `
                                                   _______________________
                                                  (___________           |
                                                    [XXXXX]   |          |
                                               __  /~~~~~~~\  |          |
                                              /  \|@@@@@@@@@\ |          |
                                              \   |@@@@@@@@@@||          |
                              ) (                 \@@@@@@@@@@||  ______  |
                             _((__                 \@@@@@@@@/ | |on|off| |
                          C\|     \               __\@@@@@@/__|  ~~~~~~  |
                            \     /              (____________|__________|
                             \___/               |_______________________|

                        __________________________________________________________
  `
	startTime := time.Now()
	tm.Clear()
	tm.MoveCursor(1, 1)
	tm.Println(text)
	tm.Println(brew)

	for {
		//	tm.Clear() // Clear current screen
		tm.MoveCursor(27, 35)

		tm.Printf("Coffee is ready since  %d Seconds", int(time.Since(startTime).Seconds()))
		tm.Flush() // Call it every time at the end of rendering
		time.Sleep(time.Second)
	}

}
