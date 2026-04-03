package actioninfo

import (
	"fmt"
	"log"
)

// описание интерфейса DataParser
type DataParser interface {
	Parse(string) error
	ActionInfo() (string, error)
}

// функция прогоняет стрез строк через объекты, удовлетворяюще интерфейсу DataParser
func Info(dataset []string, dp DataParser) {
	for _, value := range dataset {
		// объект DataParser обрабатывает строку, возможные ошибки логируются
		err := dp.Parse(value)
		if err != nil {
			log.Printf("Ошибка %v", err)
		}
		// объект DataParser возвращает результирующую строку или ошибку, ошибка логируется
		str, err := dp.ActionInfo()
		if err != nil {
			log.Printf("Ошибка %v", err)
		}
		fmt.Println(str)
	}
}
