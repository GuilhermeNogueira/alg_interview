package main

var closedBrackets = []string{")", "]", "}"}

func isBalanced(input string) bool {

	stack := make([]string, 0)

	for _, v := range input {

		size := len(stack)

		stringVal := string(v)
		if isClosedBracket(stringVal) {
			if size == 0 {
				return false
			}
			if reverse(stack[size-1]) == stringVal {
				stack = stack[:size-1]
			} else {
				return false
			}
		} else {
			stack = append(stack, stringVal)
		}
	}

	return len(stack) == 0
}

func isClosedBracket(i string) bool {
	for _, v := range closedBrackets {
		if v == i {
			return true
		}
	}
	return false
}

func reverse(i string) string {
	switch i {
	case "(":
		return ")"
	case "{":
		return "}"
	case "[":
		return "]"
	default:
		return ""
	}
}
