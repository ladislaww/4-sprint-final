package spentcalories

import (
	"fmt"
	"strconv"
	"strings"
	"time"

)

// Основные константы, необходимые для расчетов.
const (
	lenStep                    = 0.65 // средняя длина шага.
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе
)

func parseTraining(data string) (int, string, time.Duration, error) {	// "3456(шагов),Ходьба,3h00m"
	
	// Сплитим данные 
	splitedData := strings.Split(data, ",")	
	if len(splitedData) != 3 {
		return 0, "", 0, fmt.Errorf("invalid data format, expected 3 values")
	}

	// Вычленяем из слайса кол-во шагов
	numberOfSteps, err := strconv.Atoi(splitedData[0])
	if err != nil {
		return 0, "", 0, fmt.Errorf("incorrect number of steps entered")
	} 
	if numberOfSteps <= 0 {
		return 0, "", 0, fmt.Errorf("the number of steps cannot be zero or negative")
	}

	// Вычленяем из слайса вид активности
	activity := splitedData[1] 

	// Вычленяем из слайса продолжительность активности 
	activityDuration, err := time.ParseDuration(splitedData[2])
	if err != nil {
		return 0, "", 0, fmt.Errorf("activity duration parsing error")
	}
	if activityDuration <= 0 {
		return 0, "", 0, fmt.Errorf("the duration time of an activity cannot be zero or negative")
	}

	return numberOfSteps, activity, activityDuration, nil
}

func distance(steps int, height float64) float64 {
	
	// Расчет длиины шага
	stepLength := height * stepLengthCoefficient 

	// вычисляем количество пройденых километров 
	kmPassed := float64(steps) * stepLength / mInKm

	return kmPassed
 }

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	
	// Проверяем duration
	if duration <=0 {
		return 0 
	}

	if steps <=0 {
		return 0 
	}

	// Вычисляем дистанцию
	distance := distance(steps, height)

	// Скорость в км/ч
	speed := distance / duration.Hours()

	return speed
}

func TrainingInfo(data string, weight, height float64) (string, error) { // "3456,Ходьба,3h00m"
	numberOfSteps, activity, activityDuration, err := parseTraining(data)
	if err != nil {
		return "", err
	}

	dist := distance(numberOfSteps, height)
	speed :=  meanSpeed(numberOfSteps, height, activityDuration)

	switch activity {
	case "Бег":
		runningCalories, err := RunningSpentCalories(numberOfSteps, weight, height, activityDuration)
		if err != nil {
		return "", err
		}
		return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
		 activity, 
		 activityDuration.Hours(), 
		 dist, 
		 speed, 
		 runningCalories), nil
	case "Ходьба":
		walkingCalories, err := WalkingSpentCalories(numberOfSteps, weight, height, activityDuration) 
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
		 activity, 
		 activityDuration.Hours(), 
		 dist, 
		 speed, 
		 walkingCalories), nil
	default:
		return "", fmt.Errorf("неизвестный тип тренировки: %s", activity)
		}
	}	
	
	

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	
	// Проверяем корректность данных 
	if steps <= 0 || height <= 0 || duration <= 0 || weight <= 0 {
		return 0, fmt.Errorf("invalid input: steps=%d, height=%.2f, duration=%s, weigt = %.2f", steps, height, duration, weight)
	} 
	
	speed := meanSpeed(steps, height, duration)

	// Количество калорий, потраченных при беге
	spentCalories := (weight*speed*duration.Minutes()) / minInH

	return spentCalories, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	
	// Проверяем корректность данных 
	if steps <= 0 || height <= 0 || duration <= 0 || weight <= 0 {
		return 0, fmt.Errorf("invalid input: steps=%d, height=%.2f, duration=%s, weigt = %.2f", steps, height, duration, weight)
	} 
	
	speed := meanSpeed(steps, height, duration)

	spentCalories := (weight*speed*duration.Minutes()) / minInH

	// Количество калорий, потраченных при ходьбе
	spentCalories *= walkingCaloriesCoefficient

	return spentCalories, nil
}
