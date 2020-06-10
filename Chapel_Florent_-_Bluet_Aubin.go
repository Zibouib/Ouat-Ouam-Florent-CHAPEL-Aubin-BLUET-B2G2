package main

import "github.com/tadvi/winc"

func main() {

	var joueur bool = true
	var jouer [10]bool
	var jouerG [10]int

	mainWindow := winc.NewForm(nil)
	mainWindow.SetSize(400, 300)
	mainWindow.SetText("TicTacToe")

	/*WinWindow := winc.NewForm(nil) // VOUS N ETES PAS PRET
	WinWindow.SetSize(400, 300)
	WinWindow.SetText("TicTacToe") */

	btn := winc.NewPushButton(mainWindow) // BOUTON FERMER
	btn.SetText("Fermer")
	btn.SetPos(150, 200)
	btn.SetSize(100, 40)
	btn.OnClick().Bind(wndOnClose)

	btn1 := winc.NewPushButton(mainWindow) // BTN 1
	btn1.SetText("Rien")
	btn1.SetPos(0, 0)
	btn1.SetSize(40, 40)
	btn1.OnClick().Bind(func(e *winc.Event) {
		if jouer[1] == false {
			if joueur == true {
				btn1.SetText("X")
				joueur = false
				jouerG[1] = 1
			} else {
				btn1.SetText("O")
				joueur = true
				jouerG[1] = 2
			}
			jouer[1] = true
		}
		if (jouerG[1] == 1 && jouerG[2] == 1 && jouerG[3] == 1) || (jouerG[1] == 1 && jouerG[5] == 1 && jouerG[9] == 1) || (jouerG[1] == 1 && jouerG[4] == 1 && jouerG[7] == 1) {
			mainWindow.Close()
			WinJ1()
		} else if (jouerG[1] == 2 && jouerG[2] == 2 && jouerG[3] == 2) || (jouerG[1] == 2 && jouerG[5] == 2 && jouerG[9] == 2) || (jouerG[1] == 2 && jouerG[4] == 2 && jouerG[7] == 2) {
			mainWindow.Close()
			WinJ2()
		}
	})

	btn2 := winc.NewPushButton(mainWindow) // BTN 2
	btn2.SetText("Rien")
	btn2.SetPos(40, 0)
	btn2.SetSize(40, 40)
	btn2.OnClick().Bind(func(e *winc.Event) {
		if jouer[2] == false {
			if joueur == true {
				btn2.SetText("X")
				joueur = false
				jouerG[2] = 1
			} else {
				btn2.SetText("O")
				joueur = true
				jouerG[2] = 2
			}
			jouer[2] = true
		}
		if (jouerG[1] == 1 && jouerG[2] == 1 && jouerG[3] == 1) || (jouerG[2] == 1 && jouerG[5] == 1 && jouerG[8] == 1) {
			mainWindow.Close()
			WinJ1()
		} else if (jouerG[1] == 2 && jouerG[2] == 2 && jouerG[3] == 2) || (jouerG[2] == 2 && jouerG[5] == 2 && jouerG[8] == 2) {
			mainWindow.Close()
			WinJ2()
		}
	})

	btn3 := winc.NewPushButton(mainWindow) // BTN 3
	btn3.SetText("Rien")
	btn3.SetPos(80, 0)
	btn3.SetSize(40, 40)
	btn3.OnClick().Bind(func(e *winc.Event) {
		if jouer[3] == false {
			if joueur == true {
				btn3.SetText("X")
				joueur = false
				jouerG[3] = 1
			} else {
				btn3.SetText("O")
				joueur = true
				jouerG[3] = 2
			}
			jouer[3] = true
		}
		if (jouerG[1] == 1 && jouerG[2] == 1 && jouerG[3] == 1) || (jouerG[3] == 1 && jouerG[5] == 1 && jouerG[7] == 1) || (jouerG[3] == 1 && jouerG[6] == 1 && jouerG[9] == 1) {
			mainWindow.Close()
			WinJ1()
		} else if (jouerG[1] == 2 && jouerG[2] == 2 && jouerG[3] == 2) || (jouerG[3] == 2 && jouerG[5] == 2 && jouerG[7] == 2) || (jouerG[3] == 2 && jouerG[6] == 2 && jouerG[9] == 2) {
			mainWindow.Close()
			WinJ2()
		}
	})

	btn4 := winc.NewPushButton(mainWindow) // BTN 4
	btn4.SetText("Rien")
	btn4.SetPos(0, 40)
	btn4.SetSize(40, 40)
	btn4.OnClick().Bind(func(e *winc.Event) {
		if jouer[4] == false {
			if joueur == true {
				btn4.SetText("X")
				joueur = false
				jouerG[4] = 1
			} else {
				btn4.SetText("O")
				joueur = true
				jouerG[4] = 2
			}
			jouer[4] = true
		}
		if (jouerG[4] == 1 && jouerG[5] == 1 && jouerG[6] == 1) || (jouerG[1] == 1 && jouerG[4] == 1 && jouerG[7] == 1) {
			mainWindow.Close()
			WinJ1()
		} else if (jouerG[4] == 2 && jouerG[5] == 2 && jouerG[6] == 2) || (jouerG[1] == 2 && jouerG[4] == 2 && jouerG[7] == 2) {
			mainWindow.Close()
			WinJ2()
		}
	})

	btn5 := winc.NewPushButton(mainWindow) // BTN 5
	btn5.SetText("Rien")
	btn5.SetPos(40, 40)
	btn5.SetSize(40, 40)
	btn5.OnClick().Bind(func(e *winc.Event) {
		if jouer[5] == false {
			if joueur == true {
				btn5.SetText("X")
				joueur = false
				jouerG[5] = 1
			} else {
				btn5.SetText("O")
				joueur = true
				jouerG[5] = 2
			}
			jouer[5] = true
		}
		if (jouerG[4] == 1 && jouerG[5] == 1 && jouerG[6] == 1) || (jouerG[1] == 1 && jouerG[5] == 1 && jouerG[9] == 1) || (jouerG[2] == 1 && jouerG[5] == 1 && jouerG[8] == 1) || (jouerG[3] == 1 && jouerG[5] == 1 && jouerG[7] == 1) {
			mainWindow.Close()
			WinJ1()
		} else if (jouerG[4] == 2 && jouerG[5] == 2 && jouerG[6] == 2) || (jouerG[1] == 2 && jouerG[5] == 2 && jouerG[9] == 2) || (jouerG[2] == 2 && jouerG[5] == 2 && jouerG[8] == 2) || (jouerG[3] == 1 && jouerG[5] == 1 && jouerG[7] == 1) {
			mainWindow.Close()
			WinJ2()
		}
	})

	btn6 := winc.NewPushButton(mainWindow) // BTN 6
	btn6.SetText("Rien")
	btn6.SetPos(80, 40)
	btn6.SetSize(40, 40)
	btn6.OnClick().Bind(func(e *winc.Event) {
		if jouer[6] == false {
			if joueur == true {
				btn6.SetText("X")
				joueur = false
				jouerG[6] = 1
			} else {
				btn6.SetText("O")
				joueur = true
				jouerG[6] = 2
			}
			jouer[6] = true
		}
		if (jouerG[4] == 1 && jouerG[5] == 1 && jouerG[6] == 1) || (jouerG[3] == 1 && jouerG[6] == 1 && jouerG[9] == 1) {
			mainWindow.Close()
			WinJ1()
		} else if (jouerG[4] == 2 && jouerG[5] == 2 && jouerG[6] == 2) || (jouerG[3] == 2 && jouerG[6] == 2 && jouerG[9] == 2) {
			mainWindow.Close()
			WinJ2()
		}
	})

	btn7 := winc.NewPushButton(mainWindow) // BTN 7
	btn7.SetText("Rien")
	btn7.SetPos(0, 80)
	btn7.SetSize(40, 40)
	btn7.OnClick().Bind(func(e *winc.Event) {
		if jouer[7] == false {
			if joueur == true {
				btn7.SetText("X")
				joueur = false
				jouerG[7] = 1
			} else {
				btn7.SetText("O")
				joueur = true
				jouerG[7] = 2
			}
			jouer[7] = true
		}
		if (jouerG[7] == 1 && jouerG[8] == 1 && jouerG[9] == 1) || (jouerG[7] == 1 && jouerG[5] == 1 && jouerG[3] == 1) || (jouerG[1] == 1 && jouerG[4] == 1 && jouerG[7] == 1) {
			mainWindow.Close()
			WinJ1()
		} else if (jouerG[7] == 2 && jouerG[8] == 2 && jouerG[9] == 2) || (jouerG[3] == 2 && jouerG[5] == 2 && jouerG[7] == 2) || (jouerG[1] == 2 && jouerG[4] == 2 && jouerG[7] == 2) {
			mainWindow.Close()
			WinJ2()
		}
	})

	btn8 := winc.NewPushButton(mainWindow) // BTN 8
	btn8.SetText("Rien")
	btn8.SetPos(40, 80)
	btn8.SetSize(40, 40)
	btn8.OnClick().Bind(func(e *winc.Event) {
		if jouer[8] == false {
			if joueur == true {
				btn8.SetText("X")
				joueur = false
				jouerG[8] = 1
			} else {
				btn8.SetText("O")
				joueur = true
				jouerG[8] = 2
			}
			jouer[8] = true
		}
		if (jouerG[7] == 1 && jouerG[8] == 1 && jouerG[9] == 1) || (jouerG[2] == 1 && jouerG[5] == 1 && jouerG[8] == 1) {
			mainWindow.Close()
			WinJ1()
		} else if (jouerG[7] == 2 && jouerG[8] == 2 && jouerG[9] == 2) || (jouerG[2] == 2 && jouerG[5] == 2 && jouerG[8] == 2) {
			mainWindow.Close()
			WinJ2()
		}
	})

	btn9 := winc.NewPushButton(mainWindow) // BTN 9
	btn9.SetText("Rien")
	btn9.SetPos(80, 80)
	btn9.SetSize(40, 40)
	btn9.OnClick().Bind(func(e *winc.Event) {
		if jouer[9] == false {
			if joueur == true {
				btn9.SetText("X")
				joueur = false
				jouerG[9] = 1
			} else {
				btn9.SetText("O")
				joueur = true
				jouerG[9] = 2
			}
			jouer[9] = true
		}
		if (jouerG[7] == 1 && jouerG[8] == 1 && jouerG[9] == 1) || (jouerG[1] == 1 && jouerG[5] == 1 && jouerG[9] == 1) || (jouerG[3] == 1 && jouerG[6] == 1 && jouerG[9] == 1) {
			mainWindow.Close()
			WinJ1()
		} else if (jouerG[7] == 2 && jouerG[8] == 2 && jouerG[9] == 2) || (jouerG[1] == 2 && jouerG[5] == 2 && jouerG[9] == 2) || (jouerG[3] == 2 && jouerG[6] == 2 && jouerG[9] == 2) {
			mainWindow.Close()
			WinJ2()
		}
	})

	mainWindow.Center()
	mainWindow.Show()
	mainWindow.OnClose().Bind(wndOnClose)
	winc.RunMainLoop()
}
func wndOnClose(arg *winc.Event) {
	winc.Exit()
}

func WinJ1() {
	WinWindow := winc.NewForm(nil) // VOUS N ETES PAS PRET
	WinWindow.SetSize(400, 300)
	WinWindow.SetText("Joueur 1 WIN")
	WinWindow.Show()
	WinWindow.OnClose().Bind(wndOnClose)
	winc.RunMainLoop()

	btn := winc.NewPushButton(WinWindow) // BOUTON FERMER
	btn.SetText("Fermer")
	btn.SetPos(150, 200)
	btn.SetSize(100, 40)
	btn.OnClick().Bind(wndOnClose)
}

func WinJ2() {
	WinWindow := winc.NewForm(nil) // VOUS N ETES PAS PRET
	WinWindow.SetSize(400, 300)
	WinWindow.SetText("Joueur 2 WIN")
	WinWindow.Show()
	WinWindow.OnClose().Bind(wndOnClose)
	winc.RunMainLoop()

	btn := winc.NewPushButton(WinWindow) // BOUTON FERMER
	btn.SetText("Fermer")
	btn.SetPos(150, 200)
	btn.SetSize(100, 40)
	btn.OnClick().Bind(wndOnClose)
}
