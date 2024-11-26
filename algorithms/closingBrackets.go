package algorithms

import (
	"maps"
	"slices"
)

func declareBracketsMap() map[rune]rune {
	bracketsMap := make(map[rune]rune)
	bracketsMap['['] = ']'
	bracketsMap['{'] = '}'
	bracketsMap['('] = ')'
	return bracketsMap
}

func getOpenBracketsArr(bracketsMap map[rune]rune) []rune {
	openBracketsMap := maps.Keys(bracketsMap)
	openBracketsArr := make([]rune, 0)
	for bracket := range openBracketsMap {
		openBracketsArr = append(openBracketsArr, bracket)
	}
	return openBracketsArr
}

func peek(stack []rune) rune {
	return stack[len(stack)-1]
}

func ValidClosingBrackets(brackets []rune) bool {
	bracketsMap := declareBracketsMap()

	openBracketsArr := getOpenBracketsArr(bracketsMap)

	stack := make([]rune, 0)

	for i := range brackets {
		if len(stack) == 0 {
			stack = append(stack, brackets[i])
			continue
		}
		if slices.Contains(openBracketsArr, brackets[i]) {
			stack = append(stack, brackets[i])
			continue
		}
		lastBracketInStack := peek(stack)
		targetBracket := bracketsMap[lastBracketInStack]
		if brackets[i] != targetBracket {
			return false
		}
		stack = stack[:len(stack)-1]
	}

	return len(stack) == 0
}
