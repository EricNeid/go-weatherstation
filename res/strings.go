package res

// DE stand for german language
const DE Locale = "de"

// EN stands for english language
const EN Locale = "en"

// Locale represent an enum for the used language
type Locale string

// CurrentLocale configures the locale used to return localized strings
var CurrentLocale Locale = EN

// Labels contains all localized label strings
var Labels = map[string]map[Locale]string{
	"today": {
		DE: "Heute",
		EN: "Today",
	},
	"tomorrow": {
		DE: "Morgen",
		EN: "Tomorrow",
	},
	"aftertomorrow": {
		DE: "Übermorgen",
		EN: "Day after tomorrow",
	},
}

// GetLabel returns localiced label for given key. See Labels for valid keys
func GetLabel(key string) string {
	return Labels[key][CurrentLocale]
}
