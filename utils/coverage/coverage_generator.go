package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Abre el archivo de cobertura
	file, err := os.Open("coverage-report.out")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var totalLines, coveredLines int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		if len(parts) == 3 {
			count, err := strconv.Atoi(parts[1])
			if err != nil {
				panic(err)
			}
			totalLines += count

			if parts[2] == "1" {
				coveredLines += count
			}
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	// Calcular el porcentaje de cobertura
	coverage := 0.0
	if totalLines > 0 {
		coverage = (float64(coveredLines) / float64(totalLines)) * 100
	}

	// Leer el archivo README.md
	readmeFile, err := os.Open("README.md")
	if err != nil {
		panic(err)
	}
	defer readmeFile.Close()

	var updatedLines []string
	coverageLineFound := false

	scanner = bufio.NewScanner(readmeFile)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "| Cobertura") && !coverageLineFound {
			// Actualizar la línea de cobertura
			updatedLines = append(updatedLines, fmt.Sprintf("| Cobertura     | %.2f%%        |", coverage))
			coverageLineFound = true
		} else {
			updatedLines = append(updatedLines, line)
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	// Escribir el contenido actualizado al README.md
	err = os.WriteFile("README.md", []byte(strings.Join(updatedLines, "\n")), 0644)
	if err != nil {
		panic(err)
	}
}
