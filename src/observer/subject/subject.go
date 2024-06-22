package subject

type Subject interface {
	RegisterObserver(o Observer)
	RemoveObserver(o Observer)
	NotifyObserver()
}

type Observer interface {
	Update(temp, pressure, humidity float64)
}

type DisplayElement interface {
	Display()
}

// 1. Define the 'subject' package
