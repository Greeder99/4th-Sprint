package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

var (
	StepLength = 0.65 // длина шага в метрах
)

func parsePackage(data string) (int, time.Duration, error) {
	// ваш код ниже
	list := strings.Split(data, ",")
	if len(list) != 2 {
		return 0, 0, errors.New("Разделение прошло неверно")
	}
	steps, err := strconv.Atoi(list[0])
	if err != nil {
		return 0, 0, err
	}
	duration, err := time.ParseDuration(list[1])
	if err != nil {
		return 0, 0, err
	}
	return steps, duration, nil
}

// DayActionInfo обрабатывает входящий пакет, который передаётся в
// виде строки в параметре data. Параметр storage содержит пакеты за текущий день.
// Если время пакета относится к новым суткам, storage предварительно
// очищается.
// Если пакет валидный, он добавляется в слайс storage, который возвращает
// функция. Если пакет невалидный, storage возвращается без изменений.
func DayActionInfo(data string, weight, height float64) string {
	// ваш код ниже
	steps, duration, err := parsePackage(data)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	if steps <= 0 {
		return ""
	}
	distance := float64(steps) * StepLength / float64(1000)
	calories := WalkingSpentCalories(duration)
	text := fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал. ", steps, distance, calories)
	return text
}
