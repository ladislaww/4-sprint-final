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
		return 0, 0, fmt.Errorf("Длинна %d не != 2", len(splitedData))
	}

	// Вычленяем из слайса кол-во шагов
	numberOfSteps, err := strconv.Atoi(splitedData[0])
	if err != nil {
		return 0, 0, fmt.Errorf("Неверно введено количество шагов")
	}
	if numberOfSteps == 0 {
		return 0, 0, fmt.Errorf("Количество шагов не может быть нулевым")
	}

	// Вычленяем из слайса время прогулки
	walkingDuration, err := time.ParseDuration(splitedData[1])
	if err != nil {
		return 0, 0, fmt.Errorf("Ошибка парсинга продолжительности прогулки")
	}
	if walkingDuration == 0 {
		return 0, 0, fmt.Errorf("Длительность прогулки не может быть нулевой")
	}

	return numberOfSteps, walkingDuration, nil
}

func DayActionInfo(data string, weight, height float64) string {
	
	// Парсим данные 
	numberOfSteps, walkingDuration, err := parsePackage(data)
	if err != nil {
		return fmt.Sprintf("Ошибка %v", err) 
	} 
	if numberOfSteps == 0 || walkingDuration == 0 { 
		return ""
	} 
	
	// Получаем дистанцию в метрах 
	distance :=  float64(numberOfSteps) * stepLength / mInKm   
	
	// Реализовать подсчет ккал с помощью WalkingSpentCalories()
	walkingCalories, _ := spentcalories.WalkingSpentCalories(numberOfSteps, weight, height, walkingDuration)
	
	
	// Сохраняем все данные в одну строку 
	resultStr := fmt.Sprintf("Количество шагов: %d.\nДистанция составила: %.2f.\n Вы сожгли: %.2f ккал.", numberOfSteps, distance, walkingCalories)

	return resultStr

}
