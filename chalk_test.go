package chalk

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type test struct {
	expected interface{}
	actual   interface{}
}

func TestColor(t *testing.T) {
	assertions := assert.New(t)

	color := Color{1}
	tests := []test{
		{color.Value(), 1},
		{"\u001b[31m", color.String()},
		{fmt.Sprintf("%shello%s", "\u001b[31m", ResetColor), color.Set("hello")},
	}

	for _, tc := range tests {
		assertions.Equalf(tc.expected, tc.actual, "%#v should be %#v", tc.actual, tc.expected)
	}
}

func TestBrightColor(t *testing.T) {
	assertions := assert.New(t)

	color := BrightColor{Color{1}}
	tests := []test{
		{color.Value(), 1},
		{"\u001b[91m", color.String()},
		{fmt.Sprintf("%shello%s", "\u001b[91m", ResetColor), color.Set("hello")},
	}

	for _, tc := range tests {
		assertions.Equalf(tc.expected, tc.actual, "%#v should be %#v", tc.actual, tc.expected)
	}
}

func TestTextStyle(t *testing.T) {
	assertions := assert.New(t)

	textStyle := TextStyle{4, 24}
	tests := []test{
		{"\u001b[4m\u001b[24m", textStyle.String()},
		{"\u001b[4mhello\u001b[24m", textStyle.TextStyle("hello")},
		{"hello", TextStyle{}.TextStyle("hello")},
	}

	for _, tc := range tests {
		assertions.Equalf(tc.expected, tc.actual, "%#v should be %#v", tc.actual, tc.expected)
	}
}
