package main

import (
	display "../observer/display"

	weather "../observer/weather"
)

func main() {
	weatherData := weather.NewWeatherData()

	currentDisp := display.NewCurrentConditionDisplay(weatherData)

	// fmt.Println(currentDisp)
	_ = currentDisp

	weatherData.SetMeasurements(80, 65, 30.4)
	weatherData.SetMeasurements(82, 70, 34.4)
}

/*
	change the GOPATH to absolute path upto GO-design-patterns
	eg : export GOROOT=/home/user/path/to/GO-design-patterns

	change the curr directory to src/observer/
	then run the main program using: go run main.go

*/
