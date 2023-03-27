package main

import (
	"bufio"
	"fmt"
	"minyr/yr"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var fileName string // variabel for å lagre filnavnet

	for {
		fmt.Print("Skriv 'convert' for å konvertere grader, 'average' for å beregne gjennomsnittstemperatur eller 'exit' for å avslutte programmet: ")
		scanner.Scan()
		input := strings.ToLower(scanner.Text())

		switch input {
		case "convert":
			if fileName == "" {
				fmt.Print("Skriv inn filnavn for å lagre konverterte temperaturer: ")
				scanner.Scan()
				fileName = scanner.Text() + ".csv"
			} else {
				fmt.Print("Vil du lage en ny fil? Skriv 'ja' eller 'nei': ")
				scanner.Scan()
				answer := strings.ToLower(scanner.Text())
				if answer == "ja" {
					fmt.Print("Skriv inn filnavn for å lagre konverterte temperaturer: ")
					scanner.Scan()
					fileName = scanner.Text() + ".csv"
				}
			}
			convertedTemperatures, err := yr.KonverterGrader()
			if err != nil {
				fmt.Println("Kunne ikke konvertere grader:", err)
				continue
			}
			err = yr.SkrivLinjer(convertedTemperatures, fileName)
			if err != nil {
				fmt.Println("Kunne ikke skrive til fil:", err)
			} else {
				fmt.Println("Konvertering fullført!")
			}

		case "average":
			fmt.Print("Vil du ha gjennomsnittstemperaturen i Celsius eller Fahrenheit? Skriv 'c' for Celsius eller 'f' for Fahrenheit: ")
			scanner.Scan()
			unit := strings.ToLower(scanner.Text())

			var avgTemp float64
			var err error
			switch unit {
			case "c":
				avgTemp, err = yr.GjsnittTemp()
			case "f":
				avgTemp, err = yr.GjsnittTempFahrenheit()
			default:
				fmt.Println("Ugyldig enhet. Prøv igjen.")
				continue
			}

			if err != nil {
				fmt.Println("Kunne ikke beregne gjennomsnittstemperatur:", err)
			} else {
				if unit == "c" {
					fmt.Printf("Gjennomsnittstemperatur for perioden: %.2f°C\n", avgTemp)
				} else {
					fmt.Printf("Gjennomsnittstemperatur for perioden: %.2f°F\n", avgTemp)
				}
			}

		case "exit":
			fmt.Println("Avslutter programmet...")
			return

		default:
			fmt.Println("Ugyldig input. Prøv igjen.")
		}
	}
}
