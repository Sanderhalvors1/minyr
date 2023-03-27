package yr

import (
	"os"
	"testing"
)

func TestCheckCSVLineCount(t *testing.T) {
	// Åpner CSV-filen for testing
	file, err := os.Open("kjevik-temp-celsius-20220318-20230318.csv")
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	// Leser innholdet i CSV-filen
	lines, err := lesLinjer(file)
	if err != nil {
		t.Fatal(err)
	}

	// Sjekker om antall linjer er lik 16756
	if len(lines) != 16755 {
		t.Errorf("Forventet 16756 linjer, fikk %d", len(lines))
	}
}
