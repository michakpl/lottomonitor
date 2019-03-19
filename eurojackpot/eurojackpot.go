package eurojackpot

import (
	"bufio"
	"fmt"
	"github.com/thoas/go-funk"
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
		myNumbers, myExtraNumbers := getChoosenNumbers()
		hitNumbers := funk.Intersect(r.numbers, myNumbers).([]int)
		hitNumbersCount := len(hitNumbers)

		hitExtraNumbers := funk.Intersect(r.extraNumbers, myExtraNumbers).([]int)
		hitExtraNumbersCount := len(hitExtraNumbers)

		hitResult := checkWin(hitNumbersCount, hitExtraNumbersCount)

		fmt.Println(hitResult)
		return
	}
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
			r.extraNumbers = append(r.extraNumbers, number)
		}
	}

	return r
}

func getChoosenNumbers() ([]int, []int) {
	numbers := []int{5,19,20,35,49}
	extraNumbers := []int{2,7}

	return numbers, extraNumbers
}

func checkWin(hitNumbers int,  hitExtraNumbers int) bool {
	if hitNumbers >= 3 {
		return true
	} else if (hitExtraNumbers >= 2) && (hitNumbers >= 1) {
		return true
	} else if (hitExtraNumbers >= 1) && (hitNumbers >= 2) {
		return true
	}

	return false
}

type result struct {
	date         string
	numbers      []int
	extraNumbers []int
}