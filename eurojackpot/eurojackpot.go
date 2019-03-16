package eurojackpot

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
)

func CheckNumbers() {
	lines := getNumbers()

	buildResults(lines)
}

func getNumbers() []string {
	eurojackpotUrl := os.Getenv("EUROJACKPOT_URL")

	resp, err := http.Get(eurojackpotUrl)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	reader := bufio.NewReader(resp.Body)

	var lines []string

	for {
		line, _, err := reader.ReadLine()

		if err == io.EOF {
			break
		}

		lineString := fmt.Sprintf("%s", line)

		lines = append(lines, lineString)
	}

	return lines
}

func buildResults(lines []string) {
	for index, element := range lines {
		fmt.Printf("%d: %s \n", index, element)
	}
}