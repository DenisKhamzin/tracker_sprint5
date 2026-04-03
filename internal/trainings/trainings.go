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
	Steps        int
	TrainingType string
	Duration     time.Duration
	Personal     personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {
	dataSlice := strings.Split(datastring, ",")
	if len(dataSlice) != 3 {
		return errors.New("некорректный datastring")
	}
	t.Steps, err = strconv.Atoi(dataSlice[0])
	if err != nil {
		return err
	}
	if t.Steps <= 0 {
		return errors.New("Шаги")
	}
	t.TrainingType = dataSlice[1]
	t.Duration, err = time.ParseDuration(dataSlice[2])
	if err != nil {
		return err
	}
	if t.Duration <= 0 {
		return errors.New("Неверная продолжительность")
	}
	return nil
}

func (t Training) ActionInfo() (string, error) {
	var calories float64
	trainingType := t.TrainingType
	hours := t.Duration.Hours()
	way := spentenergy.Distance(t.Steps, t.Personal.Height)
	speed := spentenergy.MeanSpeed(t.Steps, t.Personal.Height, t.Duration)
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
	default:
		//default:
		return "", errors.New("неизвестный тип тренировки")
	}
	res := fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", trainingType, hours, way, speed, calories)
	// TODO: реализовать функцию
	return res, nil
}
