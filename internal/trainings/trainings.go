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

// описание структуры Training с необходимыми полями
type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

// метод структуры Training
func (t *Training) Parse(datastring string) (err error) {
	// разбитие строки по знаку запятой
	dataSlice := strings.Split(datastring, ",")
	// проверка корректности разбития
	if len(dataSlice) != 3 {
		return errors.New("некорректный datastring")
	}
	// конвертация строки в int, обработка ошибки
	t.Steps, err = strconv.Atoi(dataSlice[0])
	if err != nil {
		return fmt.Errorf("Ошибка конвертации количества шагов: %w", err)
	}
	// проверка значения количества шагов
	if t.Steps <= 0 {
		return errors.New("Шаги")
	}
	// присвоение структуре значения тренировки
	t.TrainingType = dataSlice[1]
	// присвоение структуре значения продолжительности и обработка возможной ошибки
	t.Duration, err = time.ParseDuration(dataSlice[2])
	if err != nil {
		return fmt.Errorf("Ошибка парсинга продолжительности: %w", err)
	}
	// проверка значения продолжительности
	if t.Duration <= 0 {
		return errors.New("Неверная продолжительность")
	}
	// возврат nil в случае удачного парсинга
	return nil
}

// метод структуры Training
func (t Training) ActionInfo() (string, error) {
	// переменные для вывода в строке
	var calories float64
	trainingType := t.TrainingType
	hours := t.Duration.Hours()
	way := spentenergy.Distance(t.Steps, t.Personal.Height)
	speed := spentenergy.MeanSpeed(t.Steps, t.Personal.Height, t.Duration)
	// вычисление потраченных калорий в зависимости от типа тренировки
	switch trainingType {
	case "Бег":
		cal, err := spentenergy.RunningSpentCalories(t.Steps, t.Personal.Weight, t.Personal.Height, t.Duration)
		if err != nil {
			return "", err
		}
		calories = cal
	case "Ходьба":
		cal, err := spentenergy.WalkingSpentCalories(t.Steps, t.Personal.Weight, t.Personal.Height, t.Duration)
		if err != nil {
			return "", err
		}
		calories = cal
	// возврат ошибки в случае неизвестной тренировки
	default:
		return "", errors.New("неизвестный тип тренировки")
	}
	// формирование результируещей строки и возврат с нулевой ошибкой
	res := fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", trainingType, hours, way, speed, calories)
	return res, nil
}
