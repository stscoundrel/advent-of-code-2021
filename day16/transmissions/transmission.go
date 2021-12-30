package transmissions

import (
	"fmt"
	"strconv"
	"strings"
)

func toBinaryTransmission(transmission string) []string {
	binary := ""

	for _, character := range transmission {
		num, _ := strconv.ParseInt(string(character), 16, 64)
		binary = binary + fmt.Sprintf("%04b", num)
	}

	return strings.Split(binary, "")
}
