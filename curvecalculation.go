package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

// struct for functions
type Key struct {
	Time  float64 `json:"Time"`
	Value float64 `json:"Value"`
}

// struct for functions
type Curve struct {
	Name string `json:"Name"`
	Keys []Key  `json:"Keys"`
}

// CARGA LA CURVA JSON
func LoadCurveData(filename string) (Curve, error) {
	file, err := os.Open(filename)
	if err != nil {
		return Curve{}, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return Curve{}, fmt.Errorf("failed to read file: %w", err)
	}

	var curve Curve
	err = json.Unmarshal(data, &curve)
	if err != nil {
		return Curve{}, fmt.Errorf("failed to unmarshal JSON: %w", err)
	}

	return curve, nil
}

// EXTRAE VALOR DE CURVA
func GetCurveValue(curve Curve, rating float64) float64 {
	if len(curve.Keys) == 0 {
		return 0 // Return 0 if there's no data
	}

	// Handle edge cases where rating is out of bounds
	if rating <= curve.Keys[0].Time {
		return curve.Keys[0].Value
	}
	if rating >= curve.Keys[len(curve.Keys)-1].Time {
		return curve.Keys[len(curve.Keys)-1].Value
	}

	// Find the correct interval in the curve keys
	for i := 0; i < len(curve.Keys)-1; i++ {
		if rating >= curve.Keys[i].Time && rating <= curve.Keys[i+1].Time {
			// Linear interpolation between Keys[i] and Keys[i+1]
			t1, v1 := curve.Keys[i].Time, curve.Keys[i].Value
			t2, v2 := curve.Keys[i+1].Time, curve.Keys[i+1].Value
			// Interpolate
			return v1 + (rating-t1)*(v2-v1)/(t2-t1)
		}
	}

	return 0 // Default case, should not reach here
}

// EJECUTA CALCULO CURVA HYBRIDA 2
func runcurvehybridx(varstatone int, varstattwo int, jsonfile string) float64 {
	curve, err := LoadCurveData(jsonfile)
	if err != nil {
		fmt.Println("Error loading curve data:", err)
		return 0
	}
	rating := float64(varstatone)*0.4 + float64(varstattwo)*0.6
	resultcurve := GetCurveValue(curve, rating)
	return resultcurve
}

// EJECUTA CALCULO CURVA HYBRIDA
func runcurvehybrid(varstatone int, varstattwo int, jsonfile string) float64 {
	curve, err := LoadCurveData(jsonfile)
	if err != nil {
		fmt.Println("Error loading curve data:", err)
		return 0
	}
	rating := float64(varstatone)*0.25 + float64(varstattwo)*0.75
	resultcurve := GetCurveValue(curve, rating)
	return resultcurve
}

// EJECUTA CALCULO CURVA
func runcurve(varstatone int, jsonfile string) float64 {
	curve, err := LoadCurveData(jsonfile)
	if err != nil {
		fmt.Println("Error loading curve data:", err)
		return 0
	}
	resultcurve := GetCurveValue(curve, float64(varstatone))
	return resultcurve
}
