package eurojackpot

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
)

func CheckNumbers() {
	lines := getNumbers()

	r := buildResults(lines)
	currentDate := time.Now().Format("2006-01-02")

	if currentDate == r.date {
		fmt.Println("Today is lottery day")
		return
	}

	fmt.Println("Today is not a lottery day")
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

func buildResults(lines []string) result {
	r := result{}
	for index, element := range lines {
		if index == 0 {
			r.date = element
		} else if (index > 0) && (index < 6) {
			number, _ := strconv.Atoi(element)
			r.numbers = append(r.numbers, number)
		} else if index >= 6 {
			number, _ := strconv.Atoi(element)
			r.extraNumber = append(r.extraNumber, number)
		}
	}

	return r
}

type result struct {
	date        string
	numbers     []int
	extraNumber []int
}