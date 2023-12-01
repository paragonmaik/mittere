package writer

import (
	"fmt"
	"github.com/fatih/color"
	"io"
	"net/http"
	"os"
)

func colorizedPrint(s, colour string) {
	var c *color.Color
	switch colour {
	case "green":
		c = color.New(color.FgGreen)
	case "yellow":
		c = color.New(color.FgYellow)
	case "blue":
		c = color.New(color.FgBlue)
	case "white":
		c = color.New(color.FgWhite)
	case "black":
		c = color.New(color.FgBlack)
	default:
		c = color.New(color.FgRed)
	}
	c.Println(s)
}

func Write(res *http.Response,
	colorize bool, selectedColor string) (string, string) {
	content, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	color.NoColor = !colorize
	status := res.Status
	response := string(content)

	colorizedPrint(status, selectedColor)
	colorizedPrint(response, selectedColor)

	defer res.Body.Close()

	return status, response
}
