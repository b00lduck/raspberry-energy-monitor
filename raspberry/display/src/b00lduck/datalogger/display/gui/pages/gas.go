package pages
import "fmt"

type GasPage struct {
	BasePage
}

func CreateGasPage() Page {

	arrowUp := LoadImage("arrow_up.gif")
	arrowDown := LoadImage("arrow_down.gif")

	gasPage := *new(GasPage)
	gasPage.BasePage = NewBasePage()

	for i := 0; i < 8; i ++ {
		gasPage.BasePage.AddButton(arrowUp, 20 + i * 35, 60 , func() {
			fmt.Printf("Digit %d >UP< pressed", i)
		})
		gasPage.BasePage.AddButton(arrowDown, 20 + i * 35, 120 , func() {
			fmt.Printf("Digit %d >DOWN< pressed", i)
		})	}

	return gasPage

}
