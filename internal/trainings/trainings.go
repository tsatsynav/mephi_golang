package trainings

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type Training struct {
	personaldata.Personal
	Steps        int
	TrainingType string
	Duration     time.Duration
}

// Parse разбирает строку тренировки "3456,Ходьба,3h00m"
func (t *Training) Parse(datastring string) (err error) {
	parts := strings.Split(datastring, ",")
	if len(parts) != 3 {
		return errors.New("неверный формат строки тренировки")
	}

	steps, err := strconv.Atoi(parts[0])
	if err != nil || steps <= 0 {
		return errors.New("неверное количество шагов")
	}

	duration, err := time.ParseDuration(parts[2])
	if err != nil || duration <= 0 {
		return errors.New("неверное время")
	}

	t.Steps = steps
	t.TrainingType = parts[1]
	t.Duration = duration
	return nil
}

// ActionInfo собирает инфу о тренировке и считает калории
func (t Training) ActionInfo() (string, error) {
	if t.Steps <= 0 || t.Duration <= 0 || t.Weight <= 0 || t.Height <= 0 {
		return "", errors.New("неверные данные для расчета")
	}

	var calories float64
	var err error

	// Смотрим, какая это тренировка: ходьба или бег
	switch t.TrainingType {
	case "Ходьба":
		calories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	case "Бег":
		calories, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	default:
		return "", errors.New("неизвестный тип тренировки")
	}

	if err != nil {
		return "", err
	}

	dist := spentenergy.Distance(t.Steps, t.Height)
	speed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)
	hours := t.Duration.Hours()

	return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", t.TrainingType, hours, dist, speed, calories), nil
}