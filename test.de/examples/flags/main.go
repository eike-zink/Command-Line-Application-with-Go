package main

import (
	"flag"
	"fmt"
	"log"
	"strings"
	"time"
)

type arrayFlags []string

func (a *arrayFlags) String() string {
	return strings.Join(*a, ", ")
}

func (a *arrayFlags) Set(value string) error {
	*a = append(*a, value)
	return nil
}

func main() {

	betrieb_flag := flag.String("betrieb", "", "Nummer des Betriebs")

	// oder
	var betriebe_flags arrayFlags
	flag.Var(&betriebe_flags, "betriebe", "Betriebe (Nummern durch Komma getrennt)")

	// oder
	var jahr_flag int
	flag.IntVar(&jahr_flag, "jahr", time.Now().Year(), "Jahr")
	flag.IntVar(&jahr_flag, "j", time.Now().Year(), "Jahr (shorthand)")

	// Allgemeine Hinweise
	flag.Usage = func() {
		fmt.Println("Das Testprogramm verwendet die folgenden Parameter:")
		// erforderlich, da sonst die Defaults
		flag.PrintDefaults()
	}

	flag.Parse()

	// Validierung
	if jahr_flag < 2025 {
		log.Fatalf("Ungültiges Jahr: %v", jahr_flag)
	}

	fmt.Println("Betrieb: ", *betrieb_flag)
	fmt.Println("Betriebe: ", betriebe_flags)
	fmt.Println("Jahr: ", jahr_flag)

}
