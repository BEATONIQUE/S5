package trainings

import (
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
	personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {

	data := strings.Split(datastring, ",")
	if len(data) != 3 {
		return fmt.Errorf("incorrect number of arguments, expected 3 [trainings.go]") // все ошибки в этом пакете возвращаются и выводятся через log.Printf в actioninfo.go
	}

	steps, err := strconv.Atoi(data[0])
	if err != nil {
		return fmt.Errorf("failed to parse number of steps [trainings.go]")
	}
	if steps <= 0 {
		return fmt.Errorf("the number of steps is negative or equal to zero [trainings.go]")
	}
	t.Steps = steps

	t.TrainingType = data[1]

	duration, err := time.ParseDuration(data[2])
	if err != nil {
		return fmt.Errorf("failed to parse duration [trainings.go]")
	}
	if duration <= 0 {
		return fmt.Errorf("duration is negative or equal to zero [trainings.go]")
	}
	t.Duration = duration

	return nil
}

func (t Training) ActionInfo() (string, error) {

	distance := spentenergy.Distance(t.Steps, t.Personal.Height)
	avgSpeed := spentenergy.MeanSpeed(t.Steps, t.Personal.Height, t.Duration)

	var calories float64
	var err error

	switch t.TrainingType {

	case "Бег":
		calories, err = spentenergy.RunningSpentCalories(t.Steps, t.Personal.Weight, t.Personal.Height, t.Duration)
		if err != nil {
			return "", err
		}

	case "Ходьба":
		calories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Personal.Weight, t.Personal.Height, t.Duration)
		if err != nil {
			return "", err
		}

	default:
		return "", fmt.Errorf("unknown training type [trainings.go]")
	}

	summary := fmt.Sprintf("Тип тренировки: %s\n", t.TrainingType)
	summary += fmt.Sprintf("Длительность: %.2f ч.\n", t.Duration.Hours())
	summary += fmt.Sprintf("Дистанция: %.2f км.\n", distance)
	summary += fmt.Sprintf("Скорость: %.2f км/ч\n", avgSpeed)
	summary += fmt.Sprintf("Сожгли калорий: %.2f\n", calories)

	return summary, nil
}
