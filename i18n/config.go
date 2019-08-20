package i18n

import (
	i18nstt "github.com/rhernandez-itemsoft/helpers/i18n/structs"
)

//DefaultConfig regresa la configuración por default
func DefaultConfig() i18nstt.Config {
	return i18nstt.Config{
		Default:      "es-MX",
		URLParameter: "language",
		Languages: map[string]string{
			"en-US": "./resources/languages/en-US.ini",
			"es-MX": "./resources/languages/es-MX.ini",
		},
	}
}
