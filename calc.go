package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

var (
	yearInMinutes      = float64(60 * 24 * 365)
	uptimes            = []float64{99, 99.5, 99.9, 99.95, 99.99, 99.995, 99.999, 99.9995, 99.9999}
	plannedMaintenance = []float64{0, 25, 50, 100, 200, 500}
)

func main() {
	template, err := os.ReadFile("template.markdown")
	if err != nil {
		panic(err)
	}

	file, err := os.Create("readme.md")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	writer := bufio.NewWriter(file)

	writer.Write(template)

	// Header
	writer.WriteString("Uptimes | ")
	for i, maintenance := range plannedMaintenance {
		writer.WriteString(fmt.Sprintf("%.0f hours", maintenance))
		if i < len(plannedMaintenance)-1 {
			writer.WriteString(" | ")
		}
	}
	writer.WriteString("\n")

	// Header separator
	for i := 0; i < len(plannedMaintenance)+1; i++ {
		writer.WriteString("---")
		if i < len(plannedMaintenance) {
			writer.WriteString(" | ")
		}
	}
	writer.WriteString("\n")

	for _, uptime := range uptimes {
		writer.WriteString(fmt.Sprintf("%.4f | ", uptime))
		for i, maintenance := range plannedMaintenance {
			yearWithoutMaintenance := yearInMinutes - (maintenance * 60)
			downTimePercent := 100 - uptime
			downtime := yearWithoutMaintenance / 100 * downTimePercent

			writer.WriteString((time.Duration(downtime) * time.Minute).String())
			if i < len(plannedMaintenance)-1 {
				writer.WriteString(" | ")
			}
		}
		writer.WriteString("\n")
	}

	writer.Flush()
}
