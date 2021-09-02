package main

import (
	"fmt"
	"github.com/chalk"
)

func main() {
	fmt.Println(chalk.Green, "Hello! my name is jinsu.", chalk.BrightGreen, "Thank you!!", chalk.Reset)
	fmt.Println(chalk.Yellow, "These examples are written for use.", chalk.BrightYellow, "Thank you for reviewing it.", chalk.Reset)
	fmt.Println(chalk.Magenta.Set("I prefer red to purple."), chalk.BrightMagenta.Set("But she likes purple."))

	fmt.Println(chalk.Bold.TextStyle("I added a bright color to ttacon's module. Thank you, Ttacon."))
	fmt.Println(chalk.Italic.TextStyle("I hope you have a good day today."))
	fmt.Println(chalk.Italic, "Please don't use it like this.")
	fmt.Println(chalk.Underline.TextStyle("Korea is a really good country. I want to get a job."))

	first := chalk.BrightYellow.NewStyle().WithForeground(chalk.BrightCyan.Color)
	second := chalk.BrightRed.NewStyle().WithBackground(chalk.White)
	third := second.WithTextStyle(chalk.Italic)
	fmt.Println(first, "There's no end to the development study.", chalk.Reset)
	fmt.Println(second, "There's no end to the development study.", chalk.Reset)
	fmt.Println(third.Style("There's no end to the development study."))
	fmt.Printf("%s%s%s\n", chalk.BrightMagenta, "I prefer coffee in the morning.", chalk.Reset)

	lime := chalk.Green.NewStyle().WithBackground(chalk.Black).WithTextStyle(chalk.Bold).Style
	fmt.Println(lime("Is my lime orange tree this tree or that tree?"))
}
