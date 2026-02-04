package spentenergy

import (
	"fmt"
	"time"
)

const (
	mInKm                      = 1000
	minInH                     = 60
	stepLengthCoefficient      = 0.45
	walkingCaloriesCoefficient = 0.5
)

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// Чтобы не дублировать код - беру подсчет калорий при беге, если нет ошибок - возвращаю калории с корректировкой по коэффициенту при ходьбе
	calories, err := RunningSpentCalories(steps, weight, height, duration)
	if err != nil {
		return 0, err
	}

	return calories * walkingCaloriesCoefficient, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {

	if steps <= 0 {
		return 0, fmt.Errorf("the number of steps is negative or equal to zero [spentenergy.go]") // все ошибки в этом пакете возвращаются в trainings и daysteps и оттуда возвращаются и выводятся через log.Printf в actioninfo.go
	}
	if weight <= 0 {
		return 0, fmt.Errorf("user weight is negative or equal to zero [spentenergy.go]")
	}
	if height <= 0 {
		return 0, fmt.Errorf("user height is negative or equal to zero [spentenergy.go]")
	}
	if duration <= 0 {
		return 0, fmt.Errorf("duration is negative or equal to zero [spentenergy.go]")
	}

	avgSpeed := MeanSpeed(steps, height, duration)

	return weight * avgSpeed * duration.Minutes() / minInH, nil
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if steps < 0 || duration <= 0 {
		return 0
	}
	return Distance(steps, height) / duration.Hours()
}

func Distance(steps int, height float64) float64 {
	return height * stepLengthCoefficient * float64(steps) / mInKm
}
