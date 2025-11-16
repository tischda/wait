package main

import "fmt"

func hide_cursor() {
	fmt.Print("\033[?25l")
}

func show_cursor() {
	fmt.Print("\033[?25h")
}
