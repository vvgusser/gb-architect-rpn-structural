package main

import (
	"bufio"
	"fmt"
	"github.com/fatih/color"
	rpn2 "gusser/rpn/rpn"
	"io"
	"os"
	"strings"
)

func main() {
	PrintWelcome()
	RunRepl()
}

func PrintWelcome() {
	fmt.Println("🍺 Hi!")
	fmt.Println("Just write infix here and get RPN to answer with evaluated value.")
	fmt.Println("For example you can write (1 + 2) * 4 + 3 and get back RPN like this")
	fmt.Println("1 2 + 4 * 3 + and evaluated answer is 15")
	fmt.Println("🤖 Just fun")
}

func RunRepl() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print(">>> ")
		text, err := reader.ReadString('\n')

		if err != nil {
			if err != io.EOF {
				color.Red("error: %s", err.Error())
			}
			break
		}

		text = strings.ReplaceAll(text, "\n", "")

		rpn, ans, err := rpn2.Rpn(text)

		if err != nil {
			color.Red(err.Error())
			continue
		}

		color.Green("> %v\n", rpn)
		color.Green("> %v\n", ans)
	}
}
