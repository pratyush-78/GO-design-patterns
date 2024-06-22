package weather

import (
	subject "../subject"
)

// 2. Implement the WeatherData struct in the weather package, provided by Client

type WeatherData struct {
	temperature float64
	pressure    float64
	humidity    float64

	observers []subject.Observer
}

func NewWeatherData() *WeatherData {
	return &WeatherData{}
}

// implement the parent class methods first:

func (wd *WeatherData) RegisterObserver(o subject.Observer) {
	wd.observers = append(wd.observers, o)
}

func (wd *WeatherData) RemoveObserver(o subject.Observer) {

	for i, observer := range wd.observers {
		if observer == o {
			wd.observers = append(wd.observers[0:i], wd.observers[i+1:]...)
			break
		}
	}
}

func (wd *WeatherData) NotifyObserver() {
	for _, observer := range wd.observers {
		observer.Update(wd.temperature, wd.pressure, wd.humidity)
	}
}

// implement the child class ( weatherData methods ):

// sets the new T,P,H and call the MChanged func:

func (wd *WeatherData) SetMeasurements(T, P, H float64) {
	wd.temperature = T
	wd.pressure = P
	wd.humidity = H
	wd.MeasurementsChaneged()
}

func (wd *WeatherData) MeasurementsChaneged() {
	wd.NotifyObserver()
}
