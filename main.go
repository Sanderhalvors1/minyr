package main

import (
	"bufio"
	"fmt"
	"log"
	"minyr/conv"
	"os"
	"strconv"
	"strings"
)

func openFile(filename string) (*os.File, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func readLines(file *os.File) ([]string, error) {
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "Navn") {
			continue // ignore header line
		}
		lines = append(lines, line)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return lines, nil
}
func closeFile(file *os.File) {
	err := file.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func writeLines(lines []string, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer closeFile(file)

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	for _, line := range lines {
		fmt.Fprintln(writer, line)
	}
	return nil
}
func CelsiusToFahrenheit(celsius float64) float64 {
	return conv.CelsiusToFahrenheit(celsius)
}

func convertTemperatures() ([]string, error) {
	file, err := openFile("kjevik-temp-celsius-20220318-20230318.csv")
	if err != nil {
		return nil, err
	}

	defer closeFile(file)
	lines, err := readLines(file)
	if err != nil {
		return nil, err
	}

	convertedTemperatures := make([]string, 0, len(lines)-1) // exclude header line

	for i, line := range lines {
		if i == 0 {
			continue // ignore header line
		}

		fields := strings.Split(line, ";")
		if i == 0 {
			continue // skip header line
		}
		if len(fields) != 4 {
			return nil, fmt.Errorf("unexpected number of fields in line %d: %d", i, len(fields))
		}

		timestamp := fields[0]
		temperatureCelsius, err := strconv.ParseFloat(fields[3], 64)
		if err != nil {
			return nil, fmt.Errorf("could not parse temperature in line %d: %s", i, err)
		}
		temperatureFahrenheit := temperatureCelsius*(9.0/5.0) + 32.0

		convertedTemperature := fmt.Sprintf("%s,%.2fC,%.2fF", timestamp, temperatureCelsius, temperatureFahrenheit)
		convertedTemperatures = append(convertedTemperatures, convertedTemperature)
	}

	return convertedTemperatures, nil
}

func main() {
	// Call the convertTemperatures function to convert Celsius temperatures to Fahrenheit
	convertedTemperatures, err := convertTemperatures()
	if err != nil {
		log.Fatal(err)
	}

	// Write the converted temperatures to a new file
	err = writeLines(convertedTemperatures, "kjevik-temp-fahrenheit-20220318-20230318.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Temperatures converted successfully!")
}
