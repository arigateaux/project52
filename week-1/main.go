package main

import (
	"fmt"

	"github.com/arigateaux/project-52/week-1/funnel"
)

func main() {
	fmt.Println(
		funnel.Funnel("leave", "eave"),     // => true
		funnel.Funnel("reset", "rest"),     // => true
		funnel.Funnel("dragoon", "dragon"), // => true
		funnel.Funnel("eave", "leave"),     // => false
		funnel.Funnel("sleet", "lets"),     // => false
		funnel.Funnel("skiff", "ski"),      // => false
	)
	fmt.Println(
		funnel.Bonus("dragoon"),   // => [dragon]
		funnel.Bonus("boats"),     // => [oats bats bots boas boat]
		funnel.Bonus("affidavit"), // => []
	)
}
