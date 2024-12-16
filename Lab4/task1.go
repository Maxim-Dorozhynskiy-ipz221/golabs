package main

import (
	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
	"strconv"
)

func Count() {
	// Створення вікна
	window := ui.NewWindow("Калькулятор", 300, 250, false)
	box := ui.NewVerticalBox()
	box.SetPadded(true)

	// Панель введення параметрів
	parameters := ui.NewHorizontalBox()
	parameters.SetPadded(true)

	// Група для введення розміру вікна
	input := ui.NewGroup("Розмір вікна")
	input.SetMargined(true)

	// Група для введення даних
	inputBox := ui.NewHorizontalBox()
	inputBox.SetPadded(true)

	// Форма для введення розмірів та матеріалу
	inputValue := ui.NewForm()
	inputValue.SetPadded(true)

	// Поля для введення ширини та висоти
	h := ui.NewEntry()
	w := ui.NewEntry()

	// Комбо-бокс для вибору матеріалу
	mater := ui.NewCombobox()
	mater.Append("Дерево")
	mater.Append("Метал")
	mater.Append("Металопластик")

	// Додавання полів у форму
	inputValue.Append("Ширина, см", h, false)
	inputValue.Append("Висота, см", w, false)
	inputValue.Append("Матеріали", mater, false)

	// Вкладення форми в групу
	inputBox.Append(inputValue, false)
	input.SetChild(inputBox)

	// Група для вибору типу склопакету
	glass := ui.NewGroup("Cклопакет")
	glass.SetMargined(true)

	// Вибір типу склопакету
	boxGlass := ui.NewVerticalBox()
	boxGlass.SetPadded(true)

	glassValue := ui.NewCombobox()
	glassValue.Append("Однокамерний")
	glassValue.Append("Двокамерний")

	// Перевірка наявності підвіконня
	check := ui.NewCheckbox("Підвіконня")

	// Додавання елементів до групи
	boxGlass.Append(glassValue, false)
	boxGlass.Append(check, false)
	glass.SetChild(boxGlass)

	// Додавання груп до панелі параметрів
	parameters.Append(input, false)
	parameters.Append(glass, false)

	// Додавання панелі параметрів в головний бокс
	box.Append(parameters, false)

	// Панель для результатів
	result := ui.NewHorizontalBox()
	result.SetPadded(true)

	// Панель для текстового результату
	boxText := ui.NewHorizontalBox()
	text := ui.NewLabel("")
	boxText.Append(text, false)
	result.Append(boxText, false)

	// Панель для кнопок
	boxButt := ui.NewVerticalBox()
	boxButt.SetPadded(true)

	// Кнопка для розрахунку
	button := ui.NewButton("Розрахувати")
	boxButt.Append(button, false)

	// Вікно для помилок
	errors := ui.NewWindow("ERROR", 20, 20, false)

	// Обробник події для кнопки
	button.OnClicked(func(*ui.Button) {
		// Отримання вибору користувача
		choosen_mat := mater.Selected()
		choosen_gl := glassValue.Selected()

		// Отримання значень висоти і ширини
		height := h.Text()
		width := w.Text()

		// Перетворення значень на числа
		costH, errorH := strconv.ParseInt(height, 0, 64)
		costW, errorW := strconv.ParseInt(width, 0, 64)

		// Перевірка на правильність введених значень
		if errorH == nil && errorW == nil {
			var sum float64
			// Розрахунок вартості в залежності від вибраних параметрів
			switch {
			case choosen_mat == 0 && choosen_gl == 0:
				if check.Checked() {
					sum = (float64(costH) * float64(costW) * 2.5) + 350
				} else {
					sum = (float64(costH) * float64(costW) * 2.5)
				}
			case choosen_mat == 1 && choosen_gl == 0:
				if check.Checked() {
					sum = (float64(costH) * float64(costW) * 0.5) + 350
				} else {
					sum = (float64(costH) * float64(costW) * 0.5)
				}
			case choosen_mat == 2 && choosen_gl == 0:
				if check.Checked() {
					sum = (float64(costH) * float64(costW) * 1.5) + 350
				} else {
					sum = (float64(costH) * float64(costW) * 1.5)
				}
			case choosen_mat == 0 && choosen_gl == 1:
				if check.Checked() {
					sum = (float64(costH) * float64(costW) * 3.0) + 350
				} else {
					sum = (float64(costH) * float64(costW) * 3.0)
				}
			case choosen_mat == 1 && choosen_gl == 1:
				if check.Checked() {
					sum = (float64(costH) * float64(costW) * 1.0) + 350
				} else {
					sum = (float64(costH) * float64(costW) * 1.0)
				}
			case choosen_mat == 2 && choosen_gl == 1:
				if check.Checked() {
					sum = (float64(costH) * float64(costW) * 2) + 350
				} else {
					sum = (float64(costH) * float64(costW) * 2)
				}
			}

			// Виведення результату
			text.SetText(strconv.FormatFloat(sum, 'f', -1, 64) + " грн")
		} else {
			// Виведення помилки при невірних даних
			ui.MsgBoxError(errors, "Помилка!", "Перевірте введенні дані")
		}
	})

	// Додавання кнопки до результатів
	result.Append(boxButt, true)

	// Додавання результатів в головний бокс
	box.Append(result, true)

	// Встановлення головного вікна
	window.SetChild(box)

	// Показ вікна
	window.Show()
	window.OnClosing(func(*ui.Window) bool {
		ui.Quit()
		return true
	})
}

func main() {
	err := ui.Main(Count)
	if err != nil {
		panic(err)
	}
}
