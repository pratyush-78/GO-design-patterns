package display

import (
	"fmt"

	weather "../weather"
)

type CurrentConditionsDisplay struct {
	temperature float64
	humidity    float64

	weatherData *weather.WeatherData
}

/*
	to implement the Observer Interface,
	add the update() method here,
	with same func signature as parent ( Observer interface )
*/
func (c *CurrentConditionsDisplay) Update(T, P, H float64) {
	c.temperature = T
	c.humidity = H

	// call the display func
	c.Display()
}

func (c *CurrentConditionsDisplay) Display() {
	fmt.Printf("Current Condition : %.2f degree and %.2f humidity\n", c.temperature, c.humidity)
}

func NewCurrentConditionDisplay(weatherData *weather.WeatherData) *CurrentConditionsDisplay {

	// create
	c := &CurrentConditionsDisplay{
		weatherData: weatherData,
	}

	//register
	weatherData.RegisterObserver(c)

	return c
}
