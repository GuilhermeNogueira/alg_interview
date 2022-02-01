package main

//func main() {
//	firstRepeatedCharacter("ABCA") // returns A
//}

func firstRepeatedCharacter(input string) *string {

	occurrences := make(map[string]int)

	for _, v := range input {
		occurrences[string(v)]++
	}

	for k, v := range occurrences {
		if v > 1 {
			return &k
		}
	}

	return nil
}
