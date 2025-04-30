package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"
	"github.com/Yandex-Practicum/tracker/internal/spentcalories"

)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

func parsePackage(data string) (int, time.Duration, error) {
	
	// Разделяем данные по запятой
	splitedData := strings.Split(data, ",")
	if len(splitedData) != 2 {
		return 0, 0, fmt.Errorf("invalid data format, expected 2 values")
	}

	// Вычленяем из слайса кол-во шагов
	numberOfSteps, err := strconv.Atoi(splitedData[0])
	if err != nil {
		return 0, 0, fmt.Errorf("incorrect number of steps entered")
	}
	if numberOfSteps == 0 {
		return 0, 0, fmt.Errorf("the number of steps cannot be zero")
	}

	// Вычленяем из слайса время прогулки
	walkingDuration, err := time.ParseDuration(splitedData[1])
	if err != nil {
		return 0, 0, fmt.Errorf("activity duration parsing error")
	}
	if walkingDuration == 0 {
		return 0, 0, fmt.Errorf("the duration time of an walking cannot be zero")
	}

	return numberOfSteps, walkingDuration, nil
}

func DayActionInfo(data string, weight, height float64) string {
	
	// Парсим данные 
	numberOfSteps, walkingDuration, err := parsePackage(data)
	if err != nil {
		return ""
	} 
	if numberOfSteps == 0 || walkingDuration == 0 { 
		return ""
	} 
	
	// Получаем дистанцию в метрах 
	distance :=  float64(numberOfSteps) * stepLength / mInKm   
	
	// Реализовать подсчет ккал с помощью WalkingSpentCalories()
	walkingCalories, _ := spentcalories.WalkingSpentCalories(numberOfSteps, weight, height, walkingDuration)
	
	
	// Сохраняем все данные в одну строку 
	resultStr := fmt.Sprintf("Количество шагов: %d.\nДистанция составила: %.2f км.\nВы сожгли: %.2f ккал.", numberOfSteps, distance, walkingCalories)

	return resultStr

}
