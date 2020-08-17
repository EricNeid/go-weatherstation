package assets

// ResourceKey represent an enum for the used language.
type ResourceKey int

// Locale represents an enum for the currently used locale.
type Locale int

const (
	// EN is english
	EN Locale = iota
	// DE is german
	DE
)

// CurrentLocale configures the locale used to return localized strings
var CurrentLocale Locale = EN

const (
	// Today is today
	Today ResourceKey = iota
	// Tomorrow is tomorrow
	Tomorrow
	// AfterTomorrow is the day after tomorrow
	AfterTomorrow
	// Close is close
	Close
	// CurrentTemperature is current temperature
	CurrentTemperature
	// LastUpdate is last update
	LastUpdate
	// DayTimeTemperature is day time temperature
	DayTimeTemperature
	// MinTemperature is the lowest temperature of the day
	MinTemperature
	// MaxTemperature is the hightest temperature of the day
	MaxTemperature
)

// Labels contains all localized label strings
var Labels = map[ResourceKey]map[Locale]string{
	Today: {
		DE: "Heute",
		EN: "Today",
	},
	Tomorrow: {
		DE: "Morgen",
		EN: "Tomorrow",
	},
	AfterTomorrow: {
		DE: "Übermorgen",
		EN: "Day after tomorrow",
	},
	Close: {
		DE: "Schließen",
		EN: "Close",
	},
	CurrentTemperature: {
		DE: "aktuelle Temperatur: %.2f°",
		EN: "current temperature: %.2f°",
	},
	LastUpdate: {
		DE: "Stand: %s",
		EN: "last update: %s",
	},
	DayTimeTemperature: {
		DE: "Tagestemperatur: %.2f°",
		EN: "daytime temperature: %.2f°",
	},
	MinTemperature: {
		DE: "Tiefsttemperatur: %.2f°",
		EN: "lowest temperature: %.2f°",
	},
	MaxTemperature: {
		DE: "Höchsttemperatur: %.2f°",
		EN: "maximum  temperature: %.2f°",
	},
}

// GetLabel returns localiced label for given key. See Labels for valid keys
func GetLabel(key ResourceKey) string {
	return Labels[key][CurrentLocale]
}
