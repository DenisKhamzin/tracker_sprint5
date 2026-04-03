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
	Steps    int
	Duration time.Duration
	personaldata.Personal
	// TODO: добавить поля
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	slice := strings.Split(datastring, ",")
	if len(slice) != 2 {
		err := errors.New("неверный формат данных")
		return err
	}
	steps, err := strconv.Atoi(slice[0])
	if err != nil {
		return err
	}
	if steps <= 0 {
		return errors.New("Неверное кол-во шагов")
	}
	ds.Steps = steps
	duration, err := time.ParseDuration(slice[1])
	if err != nil {
		return err
	}
	if duration <= 0 {
		return errors.New("Неверная продолжительность")
	}
	ds.Duration = duration
	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	if ds.Weight <= 0 {
		return "", errors.New("Вес")
	}
	if ds.Height <= 0 {
		return "", errors.New("Рост")
	}
	steps := ds.Steps
	if steps <= 0 {
		return "", errors.New("Неверное количество шагов")
	}
	if ds.Duration <= 0 {
		return "", errors.New("Продолжительность")
	}
	way := spentenergy.Distance(steps, ds.Height)
	calories, err := spentenergy.WalkingSpentCalories(steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", nil
	}
	str := fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", steps, way, calories)
	return str, nil
}
