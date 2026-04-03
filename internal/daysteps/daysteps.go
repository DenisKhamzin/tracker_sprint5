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

// описание структуры DaySteps
type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

// метод Parse() для структуры DaySteps
func (ds *DaySteps) Parse(datastring string) (err error) {
	// разбиение входящей строки
	slice := strings.Split(datastring, ",")
	// проверка количества элементов в получившемся после разбиения слайсе
	if len(slice) != 2 {
		return errors.New("неверный формат данных")
	}
	// конвертация количества шагов и возврат возможной ошибки конвертации
	steps, err := strconv.Atoi(slice[0])
	if err != nil {
		return fmt.Errorf("Ошибка конвертации количества шагов: %w", err)
	}
	// проверка количества шагов
	if steps <= 0 {
		return errors.New("Неверное кол-во шагов")
	}
	ds.Steps = steps
	// конвертация продолжительности и возврат возможной ошибки конвертации
	duration, err := time.ParseDuration(slice[1])
	if err != nil {
		return fmt.Errorf("Ошибка парсинга продолжительности: %w", err)
	}
	// проверка значения продолжительности
	if duration <= 0 {
		return errors.New("Неверная продолжительность")
	}
	ds.Duration = duration
	// возврат nil в случае удачного парсинга
	return nil
}

// метод структуры DaySteps
func (ds DaySteps) ActionInfo() (string, error) {
	// проверка выводимых значений на корректность
	if ds.Weight <= 0 {
		return "", errors.New("Вес должен быть положительным")
	}
	if ds.Height <= 0 {
		return "", errors.New("Рост должен быть положительным")
	}
	steps := ds.Steps
	if steps <= 0 {
		return "", errors.New("Неверное количество шагов")
	}
	if ds.Duration <= 0 {
		return "", errors.New("Некорректная родолжительность тренировки")
	}
	// вычисление значений для вывода
	way := spentenergy.Distance(steps, ds.Height)
	calories, err := spentenergy.WalkingSpentCalories(steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", nil
	}
	// формирование строки для вывода
	str := fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", steps, way, calories)
	return str, nil
}
