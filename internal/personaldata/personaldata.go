package personaldata

import "fmt"

// Personal хранит имя, вес и рост
type Personal struct {
	Name   string
	Weight float64
	Height float64
}

// Print выводит данные о человеке
func (p Personal) Print() {
	// Выводим всё с нужной точностью (2 знака после запятой для веса и роста)
	fmt.Printf("Имя: %s\nВес: %.2f\nРост: %.2f\n", p.Name, p.Weight, p.Height)
}