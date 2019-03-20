package eurojackpot

import (
	"bufio"
	"fmt"
	"github.com/thoas/go-funk"
	"io"
	"lottomonitor/notification"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func CheckNumbers() {
	lines := getNumbers()

	r := buildResults(lines)
	myCoupon := getChoosenNumbers()

	if myCoupon.date == r.date {
		hitNumbers := funk.Intersect(r.numbers, myCoupon.numbers).([]int)
		hitNumbersCount := len(hitNumbers)

		hitExtraNumbers := funk.Intersect(r.extraNumbers, myCoupon.extraNumbers).([]int)
		hitExtraNumbersCount := len(hitExtraNumbers)

		hitResult := checkWin(hitNumbersCount, hitExtraNumbersCount)

		sendNotification(hitResult)
	}
}

func sendNotification(hitResult bool) {
	if hitResult {
		notification.Send("Eurojackpot")
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

func getChoosenNumbers() result {
	couponString := os.Getenv("EUROJACKPOT_COUPON")
	couponSlices := strings.Split(couponString, ",")

	r := result{
		date: time.Now().Format("2006-01-02"),
	}

	for index, element := range couponSlices {
		if (index > 0) && (index < 5) {
			number, _ := strconv.Atoi(element)
			r.numbers = append(r.numbers, number)
		} else if index >= 5 {
			number, _ := strconv.Atoi(element)
			r.extraNumbers = append(r.extraNumbers, number)
		}
	}

	return r
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