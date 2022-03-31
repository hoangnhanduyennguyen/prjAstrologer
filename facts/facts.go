package facts

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

type Facts struct {
	DateOfBirth            string   `json:"date_of_birth"`
	Generation             string   `json:"generation"`
	DayOld                 string   `json:"day_old"`
	MonthOld               string   `json:"month_old"`
	YearOld                int      `json:"year_old"`
	NumOfDays              int      `json:"number_of_days"`
	NumOfCandles           int      `json:"number_of_candles"`
	NextBirthdate          int      `json:"next_birthdate"`
	NumberNextBirthdate    string   `json:"number_next_birthdate"`
	DayOfWeekBirthdate     string   `json:"day_of_week_birthdate"`
	DayOfWeekNextBirthDate string   `json:"day_of_week_next_birthdate"`
	DateFact               DateFact `json:"date_fact"`
}

//interface{} empty interface - any type
type DateFact struct {
	Text   interface{} `json:"text"`
	Year   interface{} `json:"year"`
	Number interface{} `json:"number"`
	Found  interface{} `json:"found"`
	Type   interface{} `json:"type"`
}

func GetFunFacts(dobInput string) Facts {
	var month, day string
	var birthDay, birthMonth, birthYear, numOfCandles int

	//Convert input string to array
	dob := []rune(dobInput)
	//Get day
	day = string(dob[8:10])
	birthDay, _ = strconv.Atoi(day)
	//Get month
	month = string(dob[5:7])
	birthMonth, _ = strconv.Atoi(month)
	month = GetMonth(birthMonth)
	//Get year
	birthYear, _ = strconv.Atoi(string(dob[0:4]))

	// Calculate year old, month old and day old
	//Get current time
	currentTime := time.Now()
	currentDay := currentTime.Day()
	currentMonth := int(currentTime.Month())
	currentYear := currentTime.Year()
	currentTime = time.Date(currentYear, time.Month(currentMonth), currentDay, 0, 0, 0, 0, time.UTC)

	months := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

	if birthDay > currentDay {
		currentMonth = currentMonth - 1
		currentDay = currentDay + months[birthMonth-1]
	}

	if birthMonth > currentMonth {
		currentYear = currentYear - 1
		currentMonth = currentMonth + 12
	}

	dayOld := currentDay - birthDay
	monthOld := currentMonth - birthMonth
	yearOld := currentYear - birthYear

	// Calculate number of candles
	for i := 1; i <= yearOld; i++ {
		numOfCandles += i
	}

	//Calculate number of days, day of week, when is the next birthday
	dOBirth := time.Date(birthYear, time.Month(birthMonth), birthDay, 0, 0, 0, 0, time.UTC)
	var num time.Duration = currentTime.Sub(dOBirth)
	var nextBirthdate time.Duration
	var dayOfWeekBirthdate string
	var dayOfWeekNextBirthdate string
	dOBirthThisYear := time.Date(currentYear, time.Month(birthMonth), birthDay, 0, 0, 0, 0, time.UTC)

	if currentTime.Before(dOBirthThisYear) {
		nextBirthdate = currentTime.Sub(dOBirthThisYear)
		dayOfWeekBirthdate = GetDayNumberToLetter(int(dOBirthThisYear.Weekday()))
		dOfBirthNextYear := dOBirthThisYear.AddDate(1, 0, 0)
		dayOfWeekNextBirthdate = GetDayNumberToLetter(int(dOfBirthNextYear.Weekday()))
	} else {
		dOBirthThisYear = dOBirthThisYear.AddDate(1, 0, 0)
		nextBirthdate = dOBirthThisYear.Sub(currentTime)
		dOfBirthNextYear := dOBirthThisYear.AddDate(1, 0, 0)
		dayOfWeekBirthdate = GetDayNumberToLetter(int(dOBirthThisYear.Weekday()))
		dayOfWeekNextBirthdate = GetDayNumberToLetter(int(dOfBirthNextYear.Weekday()))
	}

	// Store the results
	var result Facts
	result.DateOfBirth = GetDayNumberToLetter(int(dOBirth.Weekday())) + " " + month + " " + AddSuffix(strconv.Itoa(birthDay)) + ", " + strconv.Itoa(birthYear)
	result.Generation = GetGeneration(birthYear)
	result.DayOld = strconv.Itoa(dayOld)
	result.MonthOld = strconv.Itoa(monthOld)
	result.YearOld = yearOld
	result.NumOfDays = int(num.Seconds() / 86400)
	result.NumOfCandles = numOfCandles
	result.NextBirthdate = int(nextBirthdate.Seconds() / 86400)
	result.NumberNextBirthdate = AddSuffix(strconv.Itoa(yearOld + 1))
	result.DayOfWeekBirthdate = dayOfWeekBirthdate
	result.DayOfWeekNextBirthDate = dayOfWeekNextBirthdate
	result.DateFact = GetDateFactAPI(strconv.Itoa(birthMonth) + "/" + strconv.Itoa((birthDay)))
	return result
}

func AddSuffix(num string) string {
	dayRune := []rune(num)

	if dayRune[len(dayRune)-1] == 49 && num != "11" && num != "111" {
		return num + "st"
	} else if dayRune[len(dayRune)-1] == 50 && num != "12" && num != "112" {
		return num + "nd"
	} else if dayRune[len(dayRune)-1] == 51 && num != "13" && num != "113" {
		return num + "rd"
	} else {
		return num + "th"
	}
}

func GetMonth(month int) string {

	monthMap := map[int]string{
		1:  "January",
		2:  "February",
		3:  "March",
		4:  "March",
		5:  "May",
		6:  "June",
		7:  "July",
		8:  "August",
		9:  "September",
		10: "October",
		11: "November",
		12: "December",
	}
	return monthMap[month]
}

func GetGeneration(year int) string {
	if year > 2012 {
		return "Not found"
	}
	if year >= 1997 && year <= 2012 {
		return "Gen Z"
	}
	if year >= 1981 {
		return "Milennials"
	}
	if year >= 1965 {
		return "Gen X"
	}
	if year >= 1955 {
		return "Boomers II"
	}
	if year >= 1946 {
		return "Boomers I"
	}
	if year >= 1928 {
		return "Post War"
	}
	if year >= 1922 {
		return "WW II"
	}
	return "Not found"

}

func GetDayNumberToLetter(day int) string {
	mapDay := map[int]string{
		0: "Sunday",
		1: "Monday",
		2: "Tuesday",
		3: "Wednesday",
		4: "Thursday",
		5: "Friday",
		6: "Saturday",
	}
	return mapDay[day]
}

//API
func GetDateFactAPI(date string) DateFact {

	url := "https://numbersapi.p.rapidapi.com/" + date + "/date?fragment=true&json=true"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("X-RapidAPI-Host", "numbersapi.p.rapidapi.com")
	req.Header.Add("X-RapidAPI-Key", "50be42a795mshb68f9a3788bc691p190dffjsnd755284183c7")

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		fmt.Println(err.Error())
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	var dateFact DateFact
	err1 := json.Unmarshal(body, &dateFact) // Unmarshal parses the JSON-encoded data and stores the result in the value pointed to by v.
	if err1 != nil {
		fmt.Println(err1.Error())
	}
	return dateFact
}
