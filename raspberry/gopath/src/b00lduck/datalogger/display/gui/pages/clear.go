package pages

type ClearPage struct {
	BasePage
}

func NewClearPage() Page {
	clearPage := new(ClearPage)
	clearPage.BasePage = NewBasePage()
	clearPage.BasePage.SetBackground(LoadImage("bg.png"))
	return clearPage
}

func (p *ClearPage) Process() bool {
	return true
}

