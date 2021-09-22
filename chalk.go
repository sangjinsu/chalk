// Package chalk https://github.com/ttacon/chalk
// Copyright (c) 2014 Trey Tacon
// The MIT License (MIT)

// Package chalk 는 터미널 및 콘솔 출력을 꾸며주는 패키지로 ttacon/chalk 패키지를 기반으로 제작되었습니다.
package chalk

import (
	"fmt"
)

// 글자색, 모양(굵게, 이탤릭, 흐릿하게 ...)을 나타내는 변수입니다.
var (
	Black      = Color{0}
	Red        = Color{1}
	Green      = Color{2}
	Yellow     = Color{3}
	Blue       = Color{4}
	Magenta    = Color{5}
	Cyan       = Color{6}
	White      = Color{7}
	ResetColor = Color{9}

	BrightBlack   = BrightColor{Black}
	BrightRed     = BrightColor{Red}
	BrightGreen   = BrightColor{Green}
	BrightYellow  = BrightColor{Yellow}
	BrightBlue    = BrightColor{Blue}
	BrightMagenta = BrightColor{Magenta}
	BrightCyan    = BrightColor{Cyan}
	BrightWhite   = BrightColor{White}

	Bold          = TextStyle{1, 22}
	Dim           = TextStyle{2, 22}
	Italic        = TextStyle{3, 23}
	Underline     = TextStyle{4, 24}
	Inverse       = TextStyle{7, 27}
	Hidden        = TextStyle{8, 28}
	Strikethrough = TextStyle{9, 29}

	emptyTextStyle = TextStyle{}

	Reset = &style{
		foreground: ResetColor,
		background: ResetColor,
	}
)

// Color 는 기본적으로 제공되는 색을 나타냅니다.
type Color struct {
	value int
}

func (c Color) Value() int {
	return c.value
}

// Set 은 문자열의 해당 색으로 변경하는데 사용됩니다.
func (c Color) Set(value string) string {
	return fmt.Sprintf("%s%s%s", c, value, ResetColor)
}

// String 함수는 값에 30을 더해 밝을 색을 표현합니다.
func (c Color) String() string {
	return fmt.Sprintf("\u001b[%dm", 30+c.value)
}

// NewStyle 은 새로운 스타일을 만들어 냅니다. ex) chalk.BrightRed.NewStyle().WithBackground(chalk.White)
func (c Color) NewStyle() Style {
	return &style{foreground: c}
}

// BrightColor 는 Color 구조체를 embedding 하며 밝은 색을 나타냅니다.
type BrightColor struct {
	Color
}

// String 함수는 값에 90을 더해 밝을 색을 표현합니다.
func (bc BrightColor) String() string {
	return fmt.Sprintf("\u001b[%dm", 90+bc.value)
}

type textStyleDemarcation int

func (tsd textStyleDemarcation) String() string {
	return fmt.Sprintf("\u001b[%dm", tsd)
}

// TextStyle 은 굵게, 흐리게, 기울임꼴, 밑줄, 반전, 숨김 또는 취소선을 설정합니다
type TextStyle struct {
	start textStyleDemarcation
	end   textStyleDemarcation
}

func (ts TextStyle) TextStyle(value string) string {
	if ts == emptyTextStyle {
		return value
	}
	return fmt.Sprintf("%s%s%s", ts.start, value, ts.end)
}

func (ts TextStyle) String() string {
	return fmt.Sprintf("%s%s", ts.start, ts.end)
}

func (ts TextStyle) NewStyle() Style {
	return &style{textStyle: ts}
}

type Style interface {
	Foreground(Color)
	Background(Color)
	Style(string) string
	WithBackground(Color) Style
	WithForeground(Color) Style
	WithTextStyle(TextStyle) Style
	String() string
}

type style struct {
	foreground Color
	background Color
	textStyle  TextStyle
}

func (s *style) Foreground(color Color) {
	s.foreground = color
}

func (s *style) Background(color Color) {
	s.background = color
}

func (s *style) Style(value string) string {
	return fmt.Sprintf("%s%s%s", s, s.textStyle.TextStyle(value), Reset)
}

func (s *style) WithBackground(color Color) Style {
	s.Background(color)
	return s
}

func (s *style) WithForeground(color Color) Style {
	s.Foreground(color)
	return s
}

func (s *style) WithTextStyle(textStyle TextStyle) Style {
	s.textStyle = textStyle
	return s
}

func (s *style) String() string {
	return fmt.Sprintf("\u001b[%dm\u001b[%dm", 40+s.background.Value(), 30+s.foreground.Value())
}
