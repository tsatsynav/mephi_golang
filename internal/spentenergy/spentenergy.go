package spentenergy

import (
	"errors"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000
	minInH                     = 60
	stepLengthCoefficient      = 0.45
	walkingCaloriesCoefficient = 0.5
)

// Distance считает дистанцию в километрах
func Distance(steps int, height float64) float64 {
	// Шаги * коэфф длины шага * рост / 1000 (чтобы перевести метры в километры)
	return float64(steps) * stepLengthCoefficient * height / mInKm
}

// MeanSpeed считает среднюю скорость
func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	// Если время или шаги нулевые/отрицательные, то скорость 0
	if duration <= 0 || steps <= 0 {
		return 0
	}
	dist := Distance(steps, height)
	return dist / duration.Hours()
}

// RunningSpentCalories считает калории при беге
func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// Проверяем, что все значения адекватные
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, errors.New("все значения должны быть больше нуля")
	}
	speed := MeanSpeed(steps, height, duration)
	// Вес * скорость * время в часах
	return weight * speed * duration.Hours(), nil
}

// WalkingSpentCalories считает калории при ходьбе
func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// При ходьбе тратится в 2 раза меньше (используем коэфф 0.5)
	// Поэтому сначала считаем как для бега, а потом умножаем на коэффициент
	calories, err := RunningSpentCalories(steps, weight, height, duration)
	if err != nil {
		return 0, err
	}
	return calories * walkingCaloriesCoefficient, nil
}