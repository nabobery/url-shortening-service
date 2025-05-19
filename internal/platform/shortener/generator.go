package shortener

import (
	"crypto/rand"
	"encoding/base64"
)

const (
	// DefaultLength is the default length for generated short codes
	DefaultLength = 6
)

// Generator handles the generation of short codes
type Generator struct {
	length int
}

// NewGenerator creates a new Generator instance
func NewGenerator(length int) *Generator {
	if length <= 0 {
		length = DefaultLength
	}
	return &Generator{length: length}
}

// Generate creates a new short code
func (g *Generator) Generate() string {
	// Calculate how many random bytes we need to generate the desired length
	// base64 encoding: 3 bytes -> 4 chars
	numBytes := (g.length * 3) / 4
	if (g.length*3)%4 != 0 {
		numBytes++
	}

	// Generate random bytes
	b := make([]byte, numBytes)
	_, err := rand.Read(b)
	if err != nil {
		// In case of error, fallback to a simple string
		return "fallback"
	}

	// Encode to base64 and trim to desired length
	encoded := base64.URLEncoding.EncodeToString(b)
	if len(encoded) > g.length {
		encoded = encoded[:g.length]
	}

	return encoded
}
