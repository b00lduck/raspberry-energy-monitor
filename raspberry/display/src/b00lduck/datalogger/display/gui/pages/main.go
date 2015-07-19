package pages

import (
	"b00lduck/datalogger/display/gui/elems"
	"fmt"
)

type MainPage struct {
	BasePage
	ButtonGas *elems.Button
	ButtonElc *elems.Button
	ButtonWat *elems.Button
	ButtonSys *elems.Button
}

func CreateMainPage() Page {

	mainPage := new(MainPage)
	mainPage.BasePage = NewBasePage()

	mainPage.BasePage.SetBackground(LoadImage("bg.png"))
	mainPage.ButtonGas = mainPage.BasePage.AddButton(LoadImage("button_gas.png"), 0, 199, func() {
		fmt.Println("GAS")
	})
	mainPage.ButtonElc = mainPage.BasePage.AddButton(LoadImage("button_elc.png"), 80, 199, func() {
		fmt.Println("ELC")
	})
	mainPage.ButtonWat = mainPage.BasePage.AddButton(LoadImage("button_wat.png"), 160, 199, func() {
		fmt.Println("WAT")
	})
	mainPage.ButtonSys = mainPage.BasePage.AddButton(LoadImage("button_sys.png"), 240, 199, func() {
		fmt.Println("SYS")
	})

	return mainPage

}

