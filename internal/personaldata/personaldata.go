package personaldata

import "fmt"

// структура Personal
type Personal struct {
	Name   string
	Weight float64
	Height float64
}

// метод Print() для структуры Personal
func (p Personal) Print() {
	fmt.Printf("Имя: %s\nВес: %.2f кг.\nРост: %.2f м.\n", p.Name, p.Weight, p.Height)
}
