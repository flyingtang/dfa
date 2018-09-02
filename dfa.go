package DFA

type MatchType uint8

const (
	MinMatchType MatchType = 1 << iota
	MaxMatchType
)

type SensitiveWorldLibrary map[string]interface{}

// define a global variable
var sensitiveWorldLibrary SensitiveWorldLibrary

// inital a sensitive word library by []sting
func NewSensitiveWorld(words []string) SensitiveWorldLibrary {

	if sensitiveWorldLibrary == nil {
		sensitiveWorldLibrary = make(SensitiveWorldLibrary)
	}

	length := len(words)
	currentSWL := sensitiveWorldLibrary

	for i := 0; i < length; i++ {

		wds := words[i]
		for _, value := range wds {
			word := string(value)
			if _, ok := currentSWL[word]; ok {

				currentSWL = currentSWL[word].(SensitiveWorldLibrary)
			} else {

				newSWL := make(SensitiveWorldLibrary)
				newSWL["isEnd"] = 0
				currentSWL[word] = newSWL
				currentSWL = newSWL
			}
		}
		currentSWL["isEnd"] = 1
		currentSWL = sensitiveWorldLibrary
	}
	return sensitiveWorldLibrary
}

// check sensitive word in a sentence, you must transform type of string to []sting, and them pass it here 
func (sw SensitiveWorldLibrary) CheckSensitiveWord(words []string, startIndex int, matchType MatchType) int {

	count := 0
	flag := false
	length := len(words)
	currentSWL := sw

	for i := startIndex; i < length-1; i++ {

		word := words[i]
		if _, ok := currentSWL[word]; ok {

			count++
			currentSWL = currentSWL[word].(SensitiveWorldLibrary)
			if currentSWL["isEnd"] == 1 {
				flag = true
				if matchType == MinMatchType {
					break
				}
			}
		} else {
			break
		}
	}

	if !flag && count < 2 {
		count = 0
	}
	return count
}
