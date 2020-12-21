package main

import "fmt"

func main() {
	const MENU = "Menu del giorno:\n" +
		"a. pizza\n" +
		"b. penne al pomodoro\n" +
		"c. cotoletta e patatine\n" +
		"d. crostata e caffe"
	fmt.Println(MENU)

	var scelta, spazio rune
	var ordinazione string

	for {
		fmt.Println("ordinazione?")
		fmt.Scanf("%c", &scelta)
		if scelta == 'f' {
			break
		}
		fmt.Scanf("%c", &spazio)
		fmt.Scanf("%c", &spazio)
		switch scelta {
		case 'a':
			ordinazione += "-pizza- "
		case 'b':
			ordinazione += "-penne al pomodoro- "
		case 'c':
			ordinazione += "-cotoletta e patatine- "
		case 'd':
			ordinazione += "-crostata e caffe- "
		default:
			fmt.Println("-ordinazione non valida-")
		}
	}
	fmt.Println("hai ordinato ", ordinazione)
}
