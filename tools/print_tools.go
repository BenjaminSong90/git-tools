package tools

import "fmt"

type printColor int

const (
	Red     printColor = iota // 0
	Green                     // 1（自动递增）
	Yellow                    // 2
	Blue                      // 3
	Magenta                   // 4
	Cyan                      // 5
	White                     // 6
)

func (color printColor) formatTextColor(info string) (formatString string) {
	switch color {
	case Red:
		formatString = fmt.Sprintf("\x1b[31m %s \x1b[0m", info)
	case Green:
		formatString = fmt.Sprintf("\x1b[32m %s \x1b[0m", info)
	case Yellow:
		formatString = fmt.Sprintf("\x1b[33m %s \x1b[0m", info)
	case Blue:
		formatString = fmt.Sprintf("\x1b[34m %s \x1b[0m", info)
	case Magenta:
		formatString = fmt.Sprintf("\x1b[35m %s \x1b[0m", info)
	case Cyan:
		formatString = fmt.Sprintf("\x1b[36m %s \x1b[0m", info)
	case White:
		formatString = fmt.Sprintf("\x1b[37m %s \x1b[0m", info)
	default:
		formatString = fmt.Sprintf("\x1b[31m %s \x1b[0m", info)
	}

	return
}

func Println(info string, color printColor) {
	fmt.Println(color.formatTextColor(info))
}
