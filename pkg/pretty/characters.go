package pretty

var (
	colon = []string{
		"██╗",
		"╚═╝",
		"░░░",
		"░░░",
		"██╗",
		"╚═╝",
	}

	numbers = map[int][]string{
		0: {
			"░█████╗░",
			"██╔══██╗",
			"██║░░██║",
			"██║░░██║",
			"╚█████╔╝",
			"░╚════╝░",
		},
		1: {
			"░░███╗░░",
			"░████║░░",
			"██╔██║░░",
			"╚═╝██║░░",
			"███████╗",
			"╚══════╝",
		},
		2: {
			"██████╗░",
			"╚════██╗",
			"░░███╔═╝",
			"██╔══╝░░",
			"███████╗",
			"╚══════╝",
		},

		3: {
			"██████╗░",
			"╚════██╗",
			"░█████╔╝",
			"░╚═══██╗",
			"██████╔╝",
			"╚═════╝░",
		},

		4: {
			"░░██╗██╗",
			"░██╔╝██║",
			"██╔╝░██║",
			"███████║",
			"╚════██║",
			"░░░░░╚═╝",
		},

		5: {
			"███████╗",
			"██╔════╝",
			"██████╗░",
			"╚════██╗",
			"██████╔╝",
			"╚═════╝░",
		},

		6: {
			"░█████╗░",
			"██╔═══╝░",
			"██████╗░",
			"██╔══██╗",
			"╚█████╔╝",
			"░╚════╝░",
		},

		7: {
			"███████╗",
			"╚════██║",
			"░░░░██╔╝",
			"░░░██╔╝░",
			"░░██╔╝░░",
			"░░╚═╝░░░",
		},

		8: {
			"░█████╗░",
			"██╔══██╗",
			"╚█████╔╝",
			"██╔══██╗",
			"╚█████╔╝",
			"░╚════╝░",
		},

		9: {
			"░█████╗░",
			"██╔══██╗",
			"╚██████║",
			"░╚═══██║",
			"░█████╔╝",
			"░╚════╝░",
		},
	}
)

// FormatTime takes hours, minutes, and seconds and returns a slice of strings
// that represent the time in a pretty format.
//
// Special case: if seconds is -1, the seconds will not be displayed.
func FormatTime(hours, minutes, seconds int) []string {

	characters := [][]string{{}}
	// hours
	characters = append(characters, numbers[hours/10])
	characters = append(characters, numbers[hours%10])

	characters = append(characters, colon)

	// minutes
	characters = append(characters, numbers[minutes/10])
	characters = append(characters, numbers[minutes%10])

	if seconds != -1 {
		characters = append(characters, colon)

		// seconds
		characters = append(characters, numbers[seconds/10])
		characters = append(characters, numbers[seconds%10])
	}

	lines := []string{"", "", "", "", "", ""}

	// Transpose characters
	for i := 0; i < len(characters); i++ {
		for j := 0; j < len(characters[i]); j++ {
			lines[j] += characters[i][j]
		}
	}

	return lines
}
