package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Measurement struct {
	Min   float64
	Max   float64
	Sum   float64
	Count int64
}

func min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

func max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

func main() {
	measurements, err := os.Open("measurements.txt")
	if err != nil {
		panic(err)
	}
	defer measurements.Close()

	dados := make(map[string]Measurement)

	scanner := bufio.NewScanner(measurements)

	for scanner.Scan() {
		rawData := scanner.Text()

		semicolun := strings.Index(rawData, ";")
		if semicolun == -1 {
			fmt.Println("Linha inválida:", rawData)
			continue
		}

		location := strings.TrimSpace(rawData[:semicolun])
		Rawtemp := strings.TrimSpace(rawData[semicolun+1:])

		temp, err := strconv.ParseFloat(Rawtemp, 64)
		if err != nil {
			fmt.Printf("Erro ao converter temperatura '%s': %v\n", Rawtemp, err)
			continue
		}

		measurement := dados[location]
		if measurement.Count == 0 {
			measurement = Measurement{
				Min:   temp,
				Max:   temp,
				Sum:   temp,
				Count: 1,
			}
		} else {
			measurement.Min = min(measurement.Min, temp)
			measurement.Max = max(measurement.Max, temp)
			measurement.Sum += temp
			measurement.Count++
		}

		dados[location] = measurement
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Erro ao ler o arquivo:", err)
	}

	locations := make([]string, 0, len(dados))
	for name := range dados {
		locations = append(locations, name)
	}
	sort.Strings(locations)

	fmt.Print("{")
	for i, name := range locations {
		measurement := dados[name]
		fmt.Printf(
			"%s=%.1f/%.1f/%.1f",
			name,
			measurement.Min,
			measurement.Sum/float64(measurement.Count),
			measurement.Max,
		)
		if i < len(locations)-1 {
			fmt.Print(", ")
		}
	}
	fmt.Println("}")
}
