package spentcalories

import (
	"time"
	"strings"
	"strconv"
	"fmt"
	"log"
)

// Основные константы, необходимые для расчетов.
const (
	lenStep                    = 0.65 // средняя длина шага.
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе
)

func parseTraining(data string) (int, string, time.Duration, error) {
	// TODO: реализовать функцию
	parts := strings.Split(data, ",")
 if len(parts) != 3 {
  return 0, "", 0, fmt.Errorf("invalid data format")
 }

 stepsStr := strings.TrimSpace(parts[0])
 trainingType := strings.TrimSpace(parts[1])
 durationStr := strings.TrimSpace(parts[2])

 steps, err := strconv.Atoi(stepsStr)
 if err != nil {
  return 0, "", 0, fmt.Errorf("invalid steps value: %v", err)
 }
 if steps <= 0 {
  return 0, "", 0, fmt.Errorf("steps must be positive")
 }

 duration, err := time.ParseDuration(durationStr)
 if err != nil {
  return 0, "", 0, fmt.Errorf("invalid duration value: %v", err)
 }
 if duration <= 0 {
  return 0, "", 0, fmt.Errorf("duration must be positive")
 }

 return steps, trainingType, duration, nil
}

func distance(steps int, height float64) float64 {
	// TODO: реализовать функцию
	stepLength := height * stepLengthCoefficient 
	distanceMeters := float64(steps) * stepLength
	distanceKm := distanceMeters / mInKm
	return distanceKm
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию
	if duration <= 0 {
		return 0
	}
	distanceKm := distance(steps, height)
	hours := duration.Hours()
	if hours == 0 {
		return 0
	}
	speed := distanceKm / hours
	return speed
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	// TODO: реализовать функцию
	steps, trainingType, duration, err := parseTraining(data)
 if err != nil {
	log.Println(err)
  return "", err
 }

 var calories float64

 
 switch trainingType {

 case "Бег":
  calories, err = RunningSpentCalories(steps, weight, height, duration)
  if err != nil {
   return "", err
  }

 case "Ходьба":
  calories, err = WalkingSpentCalories(steps, weight, height, duration)
  if err != nil {
   return "", err
  }

 default:
  return "", fmt.Errorf("неизвестный тип тренировки: %s", trainingType)
 }

 
 hours := duration.Hours()
 distanceKm := distance(steps, height)
 speed := meanSpeed(steps, height, duration)

 
 result := fmt.Sprintf(
  "Тип тренировки: %s\n"+
   "Длительность: %.2f ч.\n"+
   "Дистанция: %.2f км.\n"+
   "Скорость: %.2f км/ч\n"+
   "Сожгли калорий: %.2f\n",
  trainingType,
  hours,
  distanceKm,
  speed,
  calories,
 )

 return result, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, fmt.Errorf("all input values must be positive")
	}
	speed := meanSpeed(steps, height, duration)
	minutes := duration.Minutes()
	
	calories := (weight * speed * minutes) / minInH
	return calories, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, fmt.Errorf("all input values must be positive")
	}
	speed := meanSpeed(steps, height, duration)
	minutes := duration.Minutes()

	calories := ((weight * speed * minutes) / minInH) * walkingCaloriesCoefficient
	return calories, nil

}