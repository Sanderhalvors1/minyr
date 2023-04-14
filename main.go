package main

import (
	"bufio"
	"fmt"
	"log"
	"minyr/yr"
	"os"
	"strings"
)

func main() { // main funksjon som kjører og gir valg om: convert, average eller exit.
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Velg convert, average, eller exit: ")
		scanner.Scan()
		text := strings.TrimSpace(scanner.Text()) //bruker får valg

		if text == "convert" { //convert valg, lager ny fil kalt KONV.csv
			_, err := os.Stat("kjevik-temp-fahr-20220318-20230318.csv")
			if err == nil {
				fmt.Print("Filen eksisterer allerede. Vil du generere filen på nytt? (j/n): ") //filen eksisterer, skal det konverteres igjen?
				scanner.Scan()
				answer := strings.ToLower(scanner.Text())
				if answer != "j" && answer != "n" {
					log.Fatal("Ugyldig svar")
				} else if answer == "n" {
					return
				}
			}

			convertedTemperatures, err := yr.KonverterGrader() //kjører KonverterGrader funksjon fra yr.go
			if err != nil {
				log.Fatal(err)
			}

			err = yr.SkrivLinjer(convertedTemperatures, "kjevik-temp-fahr-20220318-20230318.csv") //skriver linjer i ny fil.
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("Konvertering fullført!")

		} else if text == "average" { //kjører GjsnittTemp funkjson
			fmt.Print("Velg 'c' for celsius og 'f' for fahrenheit (c/f): ")
			scanner.Scan()
			tempType := strings.TrimSpace(scanner.Text())
			if tempType == "c" {
				_, err := yr.CelsiusGjennomsnitt()
				if err != nil {
					log.Fatal(err)
				}
			} else if tempType == "f" {
				_, err := yr.FahrenheitGjennomsnitt()
				if err != nil {
					log.Fatal(err)
				}
			} else {
				fmt.Println("Ugyldig kommando!")
			}
		} else if text == "exit" { //avslutter og går ut av programmet.
			break
		} else {
			fmt.Println("Ugyldig kommando!")
		}
	}
}
