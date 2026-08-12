
func evalRPN(tokens []string) int {
	var stack []int

	for _, token := range tokens {
		switch token {
		case "+", "-", "*", "/":
			var result int
			left := stack[len(stack)-2]
			right := stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			switch token {
			case "+":
				result = left + right

			case "-":
				result = left - right

			case "*":
				result = left * right

			case "/":
				result = left / right

			}
			stack = append(stack, result)

		default:
			intVal, _ := strconv.Atoi(token)
			stack = append(stack, intVal)
		}
	}

	return stack[len(stack)-1]
}

