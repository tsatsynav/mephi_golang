package actioninfo

import (
	"fmt"
	"log"
)

// DataParser - это интерфейс, которому соответствуют и DaySteps, и Training
type DataParser interface {
	Parse(datastring string) error
	ActionInfo() (string, error)
}

// Info проходит по всем строкам датасета, парсит их и выводит результат
func Info(dataset []string, dp DataParser) {
	for _, data := range dataset {
		// Пытаемся распарсить строку
		err := dp.Parse(data)
		if err != nil {
			// Если ошибка парсинга, пишем в лог и идем к следующей строке
			log.Println(err)
			continue
		}

		// Если всё ок, получаем красивую строку
		info, err := dp.ActionInfo()
		if err != nil {
			// Если ошибка при расчетах, тоже пишем в лог
			log.Println(err)
			continue
		}

		// Выводим результат на экран
		fmt.Println(info)
	}
}