package latex

import (
	"fmt"
	"strings"
)

func ConnectWithMinusSign(a, b string) string {
	return connectWithSign("-", a, b)
}

func ConnectWithPlusSign(a, b string) string {
	return connectWithSign("+", a, b)
}

func WrapInBrackets(text string) string {
	return fmt.Sprintf(`\left[%s\right]`, text)
}

func WrapInParentheses(text string) string {
	return fmt.Sprintf(`\left(%s\right)`, text)
}

func WriteMath(text string) string {
	return fmt.Sprintf(`$%s$`, text)
}

func WriteMathLine(text string) string {
	return fmt.Sprintf("$%s$ \\newline\n", text)
}

// #region Private Methods

func connectWithSign(sign, a, b string) string {
	if sign == "+" {
		if strings.Index(b, "-") == 0 {
			return fmt.Sprintf("%s - %s", a, b[1:])
		}
		return fmt.Sprintf("%s + %s", a, b)
	} else if sign == "-" {
		if strings.Index(b, "-") == 0 {
			return fmt.Sprintf("%s + %s", a, b[1:])
		}
		return fmt.Sprintf("%s - %s", a, b)
	}

	return fmt.Sprintf("%s %s", a, b)
}

// #endregion
