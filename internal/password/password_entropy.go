package password

import (
	"errors"
	"math"
	"unicode"
)

const (
	minPasswordLength = 10

	entropyTooWeakMax = 40.0
	entropyWeakMax    = 60.0
	entropyNormalMax  = 80.0
	entropyGoodMax    = 100.0
)

var ErrPasswordTooShort = errors.New("password must be more than 10 characters")

func passwordEntropy(password string) (string, error) {
	if len(password) < minPasswordLength {
		return "", ErrPasswordTooShort
	}

	// symbols
	var hasLower, hasUpper, hasDigit, hasSpecial bool

	for _, ch := range password {
		if unicode.IsLower(ch) {
			hasLower = true
		}
		if unicode.IsUpper(ch) {
			hasUpper = true
		}
		if unicode.IsDigit(ch) {
			hasDigit = true
		}
		if unicode.IsPunct(ch) || unicode.IsSymbol(ch) {
			hasSpecial = true
		}
	}

	// bits
	var charsetSize float64

	if hasLower {
		charsetSize += 26
	}
	if hasUpper {
		charsetSize += 26
	}
	if hasDigit {
		charsetSize += 10
	}
	if hasSpecial {
		charsetSize += 32
	}

	// password entropy
	entropy := math.Log2(charsetSize) * float64(len(password))

	if entropy < entropyTooWeakMax {
		return "Password is too weak", nil
	}
	if entropy < entropyWeakMax {
		return "Password is weak", nil
	}
	if entropy < entropyNormalMax {
		return "Normal password", nil
	}
	if entropy < entropyGoodMax {
		return "Good password", nil
	}

	return "Excellent password", nil
}
