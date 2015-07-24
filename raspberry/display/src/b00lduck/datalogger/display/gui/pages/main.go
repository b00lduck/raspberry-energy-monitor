package pages

import (
	"b00lduck/datalogger/display/gui/elems"
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
	mainPage.ButtonGas = mainPage.BasePage.AddMenuButton(LoadImage("button_gas.png"), 0, 199, "GAS_1")
	mainPage.ButtonElc = mainPage.BasePage.AddMenuButton(LoadImage("button_elc.png"), 80, 199, "ELC_1")
	mainPage.ButtonWat = mainPage.BasePage.AddMenuButton(LoadImage("button_wat.png"), 160, 199, "WAT_1")
	mainPage.ButtonSys = mainPage.BasePage.AddMenuButton(LoadImage("button_sys.png"), 240, 199, "SYS_1")

	return mainPage

}

