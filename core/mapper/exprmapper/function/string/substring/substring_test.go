package substring

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var s = &Substring{}

func TestSubstring(t *testing.T) {
	// Substring(string, int, int)
	// string : input string
	// int : starting position for substring
	// int : length of the substring

	// Should produce "Flogo"
	sub := s.Eval("Flogo is the most awesome project ever", 0, 5)
	fmt.Printf("Result [%s] should be equal to: Flogo\n", sub)
	assert.Equal(t, "Flogo", sub)

	// Should produce "awesome"
	sub = s.Eval("Flogo is the most awesome project ever", 18, 7)
	fmt.Printf("Result [%s] should be equal to: awesome\n", sub)
	assert.Equal(t, "awesome", sub)

	// Should produce "ever"
	// When setting the length to negative, it will always start at the end of
	// the string and work backwards (ignoring the starting position)
	sub = s.Eval("Flogo is the most awesome project ever", 0, -4)
	fmt.Printf("Result [%s] should be equal to: ever\n", sub)
	assert.Equal(t, "ever", sub)

	// Should produce "ever"
	// When setting the length to negative, it will always start at the end of
	// the string and work backwards (ignoring the starting position)
	sub = s.Eval("Flogo is the most awesome project ever", 2, -4)
	fmt.Printf("Result [%s] should be equal to: ever\n", sub)
	assert.Equal(t, "ever", sub)

	// Should produce ""
	// When setting the length to a negative number higher than the length of
	// the string it will return an empty string
	sub = s.Eval("Flogo is the most awesome project ever", 0, -40)
	fmt.Printf("Result [%s] should be equal to: <empty>\n", sub)
	assert.Equal(t, "", sub)
}
