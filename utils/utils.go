package utils

import "errors"

func Rates(currency string, target string) (float64, error) {
	// GHS := map[string]float64{
	// 	"NGN": 69.31,
	// 	"KSH": 18.24,
	// }
	// KSH := map[string]float64{
	// 	"GHS": 0.055,
	// 	"NGN": 3.80,
	// }

	// NGN := map[string]float64{
	// 	"GHS": 0.014,
	// 	"KSH": 0.26,
	// }

	ConversionValues := map[string]map[string]float64{
		"NGN": {
			"GHS": 0.014,
			"KSH": 0.26,
		},

		"GHS": {
			"NGN": 69.31,
			"KSH": 18.24,
		},
		"KSH": {
			"GHS": 0.055,
			"NGN": 3.80,
		},
	}

	rate, ok := ConversionValues[currency]
	if ok {
		rates, ok := rate[target]
		if ok {
			return rates, nil
		}

	}

	// if currency == "GHS" {
	// 	switch target {
	// 	case "NGN":
	// 		return GHS["NGN"], nil
	// 	case "KSH":
	// 		return GHS["KSH"], nil

	// 	}
	// }
	// if currency == "NGN" {
	// 	switch target {
	// 	case "GHS":
	// 		return NGN["GHS"], nil
	// 	case "KSH":
	// 		return NGN["KSH"], nil

	// 	}
	// }
	// if currency == "KSH" {
	// 	switch target {
	// 	case "NGN":
	// 		return KSH["NGN"], nil
	// 	case "GHS":
	// 		return KSH["GHS"], nil

	// 	}
	// }

	return 0.00, errors.New("Invalid Currency")
}
