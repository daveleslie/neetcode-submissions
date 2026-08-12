func evalRPN(tokens []string) int {
	var operands []int
	operators := map[string]func(a,b int) int {
		"-": func(a, b int) int { return a - b },
		"+": func(a, b int) int { return a + b },
		"/": func(a, b int) int { return a / b },
		"*": func(a, b int) int { return a * b },
	}

	for _ , token := range tokens {
		opFunc, isOperator := operators[token]
		if isOperator {
			operands = append(operands[:len(operands)-2], opFunc(operands[len(operands) - 2], operands[len(operands) - 1]))
		} else {
			intval, _ := strconv.Atoi(token)
			operands = append(operands, intval)
		}
	}

	return operands[len(operands) - 1]

}
