// Package tern implements ternary-like operator
package tern

// String returns non nil value if available
func String(a, b string) string {
	if len(a) > 0 {
		return a
	}
	return b
}

// Uint returns non nil value if available
func Uint(a, b uint) uint {
	if a > 0 {
		return a
	}
	return b
}

// Int returns non nil value if available
func Int(a, b int) int {
	if a > 0 {
		return a
	}
	return b
}
