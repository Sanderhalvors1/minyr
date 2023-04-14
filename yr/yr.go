package yr

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"minyr/conv"
	"os"
	"strconv"
	"strings"
)

func OpenFil(filename string) (*os.File, error) { // funksjon for å åpne fil
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func LesLinjer(file *os.File) ([]string, error) { // funksjon for å lese fil
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return lines, nil
}

func LukkFil(file *os.File) { // funksjon for å lukke fil
	err := file.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func SkrivLinjer(lines []string, filename string) error { //funksjon for å skrive linjene i fila
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer LukkFil(file)

	writer := bufio.NewWriterSize(file, 4096) // bruker en buffer med størrelse 4096 bytes.
	defer writer.Flush()

	fmt.Fprintln(writer, lines[0])

	for _, line := range lines[1 : len(lines)-1] {
		fields := strings.Split(line, ";")
		if len(fields) != 4 {
			return fmt.Errorf("unexpected number of fields in line: %s", line)
		}

		location := fields[0]
		station := fields[1]
		timestamp := fields[2]
		temperatureCelsius, err := strconv.ParseFloat(fields[3], 64)
		if err != nil {
			return fmt.Errorf("could not parse temperature in line: %s", line)
		}

		temperatureFahrenheit := CelsiusToFahrenheit(temperatureCelsius)

		convertedTemperature := fmt.Sprintf("%s;%s;%s;%.1f°F", location, station, timestamp, temperatureFahrenheit)
		fmt.Fprintln(writer, convertedTemperature)
	}

	fmt.Fprintln(writer, "Data er gyldig per 18.03.2023 (CC BY 4.0), Meteorologisk institutt (MET); endringen er gjort av Amadeus Hovden")
	return nil
}

func CelsiusToFahrenheit(celsius float64) float64 { //funksjon for konvertere gradene. Hentet fra conv
	return conv.CelsiusToFahrenheit(celsius)
}

func KonverterGrader() ([]string, error) { // konevrterer gardene i kjevik fila til fahrenheit
	file, err := OpenFil("kjevik-temp-celsius-20220318-20230318.csv")
	if err != nil {
		return nil, err
	}
	defer LukkFil(file)

	lines, err := LesLinjer(file)
	if err != nil {
		return nil, err
	}

	err = SkrivLinjer(lines, "kjevik-temp-fahr-20220318-20230318.csv")
	if err != nil {
		return nil, err
	}

	return lines, nil
}

func CelsiusGjennomsnitt() (float64, error) { //funksjon for gj.snitt i celsius
	// funksjon for å regne gj.snitts temp.

	// åpner kjevik fila
	file, err := OpenFil("kjevik-temp-celsius-20220318-20230318.csv")
	if err != nil {
		return 0, err
	}
	defer LukkFil(file)

	// leser linjene
	lines, err := LesLinjer(file)
	if err != nil {
		return 0, err
	}

	// kalkulerer var
	sumCelsius := 0.0
	count := 0

	for i, line := range lines {
		if i == 0 {
			continue // ignorerer første linje
		}

		fields := strings.Split(line, ";")
		if len(fields) != 4 {
			return 0, fmt.Errorf("unexpected number of fields in line %d: %d", i, len(fields))
		}

		if fields[3] == "" {
			continue //ignorer linje uten temp field
		}

		temperatureCelsius, err := strconv.ParseFloat(fields[3], 64)
		if err != nil {
			return 0, fmt.Errorf("could not parse temperature in line %d: %s", i, err)
		}

		sumCelsius += temperatureCelsius

		count++
	}

	averageCelsius := sumCelsius / float64(count)

	averageCelsius = math.Round(averageCelsius*100) / 100 // runder opp til 2 desimaler

	fmt.Println("Gjennomsnittstemperaturen er:", averageCelsius, "°Celsius")

	return averageCelsius, nil
}

func FahrenheitGjennomsnitt() (float64, error) { // funksjon for gj.snitt i fahrenheit
	// funksjon for å regne gj.snitts temp.

	// åpner kjevik fila
	file, err := OpenFil("kjevik-temp-celsius-20220318-20230318.csv")
	if err != nil {
		return 0, err
	}
	defer LukkFil(file)

	// leser linjene
	lines, err := LesLinjer(file)
	if err != nil {
		return 0, err
	}

	// kalkulerer var
	sumFahrenheit := 0.0
	count := 0

	for i, line := range lines {
		if i == 0 {
			continue // ignorerer første linje
		}

		fields := strings.Split(line, ";")
		if len(fields) != 4 {
			return 0, fmt.Errorf("unexpected number of fields in line %d: %d", i, len(fields))
		}

		if fields[3] == "" {
			continue //ignorer linje uten temp field
		}

		temperatureCelsius, err := strconv.ParseFloat(fields[3], 64)
		if err != nil {
			return 0, fmt.Errorf("could not parse temperature in line %d: %s", i, err)
		}

		temperatureFahrenheit := CelsiusToFahrenheit(temperatureCelsius) //bruker funkjson fra funtemps

		sumFahrenheit += temperatureFahrenheit

		count++
	}

	averageFahrenheit := sumFahrenheit / float64(count)

	averageFahrenheit = math.Round(averageFahrenheit*100) / 100 // runder opp til 2 desimaler

	fmt.Println("Gjennomsnittstemperaturen er:", averageFahrenheit, "°Fahrenheit")

	return averageFahrenheit, nil
}
