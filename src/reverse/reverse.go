/*
Package reverse

Reverse a string
*/
package reverse

func Reverse(input []rune) []rune {
	for i, j := 0, len(input)-1; i < j; i, j = i+1, j-1 {
		input[i], input[j] = input[j], input[i]
	}
	return input
}

/*
	ReverseString:

	input: string

	output: string

	return a reverse of given string
*/
func ReverseString(input string) string {
	output := make([]byte, len(input))
	for i, j := 0, len(input)-1; j >= 0; i, j = i+1, j-1 {
		output[i] = input[j]
	}
	return string(output)
}
