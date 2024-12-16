package main

import (
	"fmt"
	"os"
)

// 10 варіант (парний)

type Worker struct {
	name      string
	year      int
	month     int
	workPlace Company
}

func NewWorker(name string, year int, month int, workPlace Company) Worker {
	return Worker{name, year, month, workPlace}
}

func (w *Worker) setName(name string) {
	w.name = name
}

func (w Worker) getName() string {
	return w.name
}

func (w *Worker) setYear(year int) {
	w.year = year
}

func (w Worker) getYear() int {
	return w.year
}

func (w *Worker) setMonth(month int) {
	w.month = month
}

func (w Worker) getMonth() int {
	return w.month
}

func (w *Worker) setWorkPlace(workPlace Company) {
	w.workPlace = workPlace
}

func (w Worker) getWorkPlace() Company {
	return w.workPlace
}

func (w Worker) GetWorkerPosition() string {
	return w.workPlace.position
}

func (w Worker) GetWorkExperience() int {
	return ((2022 - w.year) * 12) + (12 - w.month)
}

func (w Worker) GetTotalMoney() float64 {
	return float64(w.GetWorkExperience()) * w.workPlace.salary
}

type Company struct {
	Name     string
	position string
	salary   float64
}

func NewCompany(Name string, position string, salary float64) *Company {
	comp := new(Company)
	comp.Name = Name
	comp.position = position
	comp.salary = salary
	return comp
}

func (c *Company) setCompanyName(Name string) {
	c.Name = Name
}

func (c Company) getCompanyName() string {
	return c.Name
}

func (c *Company) setPosition(position string) {
	c.position = position
}

func (c Company) getPosition() string {
	return c.position
}

func (c *Company) setSalary(salary float64) {
	c.salary = salary
}

func (c Company) getSalary() float64 {
	return c.salary
}

func (c Company) getSalaryInDollars() float64 {
	return c.salary / 36.90
}

func (c Company) getSalaryInEuro() float64 {
	return c.salary / 38.09
}

func (c Company) getSalaryInPound() float64 {
	return c.salary / 43.90
}

func readWorkersArray() ([]Worker, error) {
	var workers []Worker
	var n int
	fmt.Println("Введіть кількість працівників: ")
	_, err := fmt.Fscan(os.Stdin, &n)
	if err != nil {
		return nil, fmt.Errorf("помилка при введенні кількості працівників: %w", err)
	}

	for i := 0; i < n; i++ {
		var name, workPlaceName, workPlacePosition string
		var year, month int
		var workPlaceSalary float64

		fmt.Println("\nВведіть прізвище працівника: ")
		_, err := fmt.Fscan(os.Stdin, &name)
		if err != nil {
			return nil, fmt.Errorf("помилка при введенні прізвища: %w", err)
		}

		fmt.Println("Введіть назву компанії: ")
		_, err = fmt.Fscan(os.Stdin, &workPlaceName)
		if err != nil {
			return nil, fmt.Errorf("помилка при введенні назви компанії: %w", err)
		}

		fmt.Println("Введіть рік початку роботи: ")
		_, err = fmt.Fscan(os.Stdin, &year)
		if err != nil {
			return nil, fmt.Errorf("помилка при введенні року початку роботи: %w", err)
		}

		fmt.Println("Введіть місяць початку роботи: ")
		_, err = fmt.Fscan(os.Stdin, &month)
		if err != nil {
			return nil, fmt.Errorf("помилка при введенні місяця початку роботи: %w", err)
		}

		fmt.Println("Введіть посаду працівника: ")
		_, err = fmt.Fscan(os.Stdin, &workPlacePosition)
		if err != nil {
			return nil, fmt.Errorf("помилка при введенні посади: %w", err)
		}

		fmt.Println("Введіть зарплату працівника: ")
		_, err = fmt.Fscan(os.Stdin, &workPlaceSalary)
		if err != nil {
			return nil, fmt.Errorf("помилка при введенні зарплати: %w", err)
		}

		w := NewWorker(name, year, month, Company{Name: workPlaceName, position: workPlacePosition, salary: workPlaceSalary})
		workers = append(workers, w)
	}

	return workers, nil
}

func PrintWorker(w Worker) {
	fmt.Println("Прізвище: " + w.name)
	fmt.Printf("Рік початку роботи: %d\n", w.year)
	fmt.Printf("Місяць початку роботи: %d\n", w.month)
	fmt.Println("Назва компанії: " + w.workPlace.Name)
	fmt.Println("Посада працівника: " + w.workPlace.position)
	fmt.Printf("Зарплата працівника: %.2f\n", w.workPlace.salary)
}

func PrintWorkers(ws []Worker) {
	for i := 0; i < len(ws); i++ {
		fmt.Printf("\nПрацівник %d-й:\n", i+1)
		PrintWorker(ws[i])
	}
}

func GetWorkersInfo(ws []Worker) (min Worker, max Worker) {
	min = ws[0]
	max = ws[0]
	for i := 0; i < len(ws); i++ {
		if ws[i].workPlace.salary < min.workPlace.salary {
			min = ws[i]
		}
		if ws[i].workPlace.salary > max.workPlace.salary {
			max = ws[i]
		}
	}
	return min, max
}

func main() {
	ws, err := readWorkersArray()
	if err != nil {
		fmt.Println("Помилка:", err)
		return
	}

	PrintWorkers(ws)
	min, max := GetWorkersInfo(ws)

	fmt.Println("\nПрацівник з мінімальною зарплатою:")
	PrintWorker(min)

	fmt.Println("\nПрацівник з максимальною зарплатою:")
	PrintWorker(max)

	var n int
	fmt.Println("\nВведіть номер працівника :")
	_, err = fmt.Fscan(os.Stdin, &n)
	if err != nil || n < 1 || n > len(ws) {
		fmt.Println("Некоректний номер працівника.")
		return
	}

	fmt.Printf("\nПрацівник %d-й:\n", n)
	fmt.Println("\nПрізвище: ", ws[n-1].name)
	fmt.Println("\nПосада: ", ws[n-1].GetWorkerPosition())
	fmt.Println("\nСтаж роботи в місяцях: ", ws[n-1].GetWorkExperience())
	fmt.Println("\nЗарплата за всі місяці: ", ws[n-1].GetTotalMoney())
	fmt.Println("\nЗарплата у доларах: ", ws[n-1].workPlace.getSalaryInDollars())
	fmt.Println("\nЗарплата у євро: ", ws[n-1].workPlace.getSalaryInEuro())
	fmt.Println("\nЗарплата у фунтах: ", ws[n-1].workPlace.getSalaryInPound())
}
