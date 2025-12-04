package algorithms

func isValidParentheses(s string) bool {
	n := len(s)
	if n%2 != 0 {
		return false
	}

	isOpening := func(char uint8) bool {
		return char == '(' || char == '[' || char == '{'
	}
	isClosure := func(char uint8) bool {
		return char == ')' || char == ']' || char == '}'
	}

	getClosureFor := func(char uint8) uint8 {
		switch char {
		case '(':
			return ')'
		case '[':
			return ']'
		case '{':
			return '}'
		default:
			return 0
		}
	}

	char := s[0]

	if !isOpening(char) {
		return false
	}

	if !isClosure(s[n-1]) {
		return false
	}

	openParenthesis := 0
	openSquareBrackets := 0
	openCurlyBraces := 0
	if char == '(' {
		openParenthesis++
	}

	for i := 1; i < n; i++ {
		switch char {
		case '(':
			openParenthesis++
		case '[':
			openSquareBrackets++
		case '{':
			openCurlyBraces++
		case ')':
			openParenthesis--
		case ']':
			openSquareBrackets--
		case '}':
			openCurlyBraces--
		}
		item := s[i]
		if isOpening(char) {
			valid := item == getClosureFor(char) || isOpening(item)
			if !valid {
				return false
			}
		}
		char = item
	}
	if openParenthesis != 0 && openSquareBrackets != 0 && openCurlyBraces != 0 {
		return false
	}
	return true
}
