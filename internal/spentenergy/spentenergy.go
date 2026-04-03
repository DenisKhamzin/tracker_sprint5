package spentenergy

import (
	"errors"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

// функция для подсчета калорий при ходьбе
func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// обработка некорректных параметров
	if steps <= 0 {
		return 0, errors.New("количество шагов должно быть положительным")
	}
	if weight <= 0 {
		return 0, errors.New("вес должен быть положительным")
	}
	if height <= 0 {
		return 0, errors.New("рост должен быть положительным")
	}
	if duration <= 0 {
		return 0, errors.New("продолжительность должна быть продолжительной")
	}
	// вычисление количества потраченных калорий
	averageSpeed := MeanSpeed(steps, height, duration)
	minutes := duration.Minutes()
	res := (weight * averageSpeed * minutes * walkingCaloriesCoefficient) / float64(minInH)
	return res, nil
}

// функция для подсчета калорий при беге
func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// обработка некорректных параметров
	if steps <= 0 {
		return 0, errors.New("количество шагов должно быть положительным")
	}
	if weight <= 0 {
		return 0, errors.New("вес должен быть положительным")
	}
	if height <= 0 {
		return 0, errors.New("рост должен быть положительным")
	}
	if duration <= 0 {
		return 0, errors.New("продолжительность должна быть продолжительной")
	}
	// вычисление количества потраченных калорий
	averageSpeed := MeanSpeed(steps, height, duration)
	minutes := duration.Minutes()
	res := (weight * averageSpeed * minutes) / float64(minInH)
	return res, nil
}

// функция для рассчета средней скорости
func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	// проверка значения продолжительности тренировки
	if duration <= 0 {
		return 0
	}
	return Distance(steps, height) / duration.Hours()
}

// функция для рассчета дистанции
func Distance(steps int, height float64) float64 {
	return height * stepLengthCoefficient * float64(steps) / float64(mInKm)
}
