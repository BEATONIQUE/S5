package daysteps

import (
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
}

func (ds *DaySteps) Parse(datastring string) (err error) {

	data := strings.Split(datastring, ",")
	if len(data) != 2 {
		return fmt.Errorf("incorrect number of arguments, expected 2 [daysteps.go]") // все ошибки в этом пакете возвращаются и выводятся через log.Printf в actioninfo.go
	}

	steps, err := strconv.Atoi(data[0])
	if err != nil {
		return fmt.Errorf("failed to parse number of steps [daysteps.go]")
	}
	if steps <= 0 {
		return fmt.Errorf("the number of steps is negative or equal to zero [daysteps.go]")
	}
	ds.Steps = steps

	duration, err := time.ParseDuration(data[1])
	if err != nil {
		return fmt.Errorf("failed to parse duration [daysteps.go]")
	}
	if duration <= 0 {
		return fmt.Errorf("duration is negative or equal to zero [daysteps.go]")
	}
	ds.Duration = duration

	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {

	distance := spentenergy.Distance(ds.Steps, ds.Personal.Height)

	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Personal.Weight, ds.Personal.Height, ds.Duration)
	if err != nil {
		return "", err
	}

	summary := fmt.Sprintf("Количество шагов: %d.\n", ds.Steps)
	summary += fmt.Sprintf("Дистанция составила %.2f км.\n", distance)
	summary += fmt.Sprintf("Вы сожгли %.2f ккал.\n", calories)

	return summary, nil
}
