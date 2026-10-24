package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type DaySteps struct {
	personaldata.Personal // встраиваем данные о пользователе
	Steps                 int
	Duration              time.Duration
}

// Parse разбирает строку вида "678,0h50m"
func (ds *DaySteps) Parse(datastring string) (err error) {
	// Разделяем строку по запятой
	parts := strings.Split(datastring, ",")
	if len(parts) != 2 {
		return errors.New("неверный формат строки")
	}

	// Парсим шаги
	steps, err := strconv.Atoi(parts[0])
	if err != nil || steps <= 0 {
		return errors.New("неверное количество шагов")
	}

	// Парсим время с помощью встроенной функции (она сама понимает "1h30m")
	duration, err := time.ParseDuration(parts[1])
	if err != nil || duration <= 0 {
		return errors.New("неверное время")
	}

	ds.Steps = steps
	ds.Duration = duration
	return nil
}

// ActionInfo формирует красивую строку с результатами
func (ds DaySteps) ActionInfo() (string, error) {
	// Проверяем, что данные вообще есть и они правильные
	if ds.Steps <= 0 || ds.Duration <= 0 || ds.Weight <= 0 || ds.Height <= 0 {
		return "", errors.New("неверные данные для расчета")
	}

	dist := spentenergy.Distance(ds.Steps, ds.Height)
	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", err
	}

	// Возвращаем отформатированную строку
	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", ds.Steps, dist, calories), nil
}