package daysteps

import (
	"time"
	"strings"
	"strconv"
	"fmt"
	"log"
	"github.com/Yandex-Practicum/tracker/internal/spentcalories"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

func parsePackage(data string) (int, time.Duration, error) {
	// TODO: реализовать функцию
parts := strings.Split(data, ",")
 if len(parts) != 2 {
  return 0, 0, fmt.Errorf("invalid data format")
 }

 stepsPart := parts[0]
 durationPart := parts[1]

 if stepsPart != strings.TrimSpace(stepsPart) {
  return 0, 0, fmt.Errorf("invalid steps format")
 }
 if durationPart != strings.TrimSpace(durationPart) {
  return 0, 0, fmt.Errorf("invalid duration format")
 }

 steps, err := strconv.Atoi(stepsPart)
 if err != nil {
  return 0, 0, fmt.Errorf("invalid steps value: %v", err)
 }
 if steps <= 0 {
  return 0, 0, fmt.Errorf("steps must be positive")
 }

 duration, err := time.ParseDuration(durationPart)
 if err != nil {
  return 0, 0, fmt.Errorf("invalid duration value: %v", err)
 }
 if duration <= 0 {
  return 0, 0, fmt.Errorf("duration must be positive")
 }

 return steps, duration, nil
}
	

func DayActionInfo(data string, weight, height float64) string {
	// TODO: реализовать функцию
 steps, duration, err := parsePackage(data)
 if err != nil {
  log.Println(err)
  return ""
 }

 
 distanceKm := float64(steps) * stepLength / mInKm

 
 calories, err := spentcalories.WalkingSpentCalories(steps, weight, height, duration)
 if err != nil {
  log.Println(err)
  return ""
 }

 return fmt.Sprintf(
  "Количество шагов: %d.\n"+
   "Дистанция составила %.2f км.\n"+
   "Вы сожгли %.2f ккал.\n",
  steps,
  distanceKm,
  calories,
 )
}