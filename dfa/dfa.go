package dfa

type MatchType uint8

const (
	MinMartchType MatchType = 1 << iota
	MaxMatchType
)

type SensitiveWorldLibary map[string]interface{}

// 定义一个全家变量
var sensitiveWorldLibary SensitiveWorldLibary

// 初始化敏感词库
func NewSensitiveWorld(words []string) SensitiveWorldLibary {

	if sensitiveWorldLibary == nil {
		sensitiveWorldLibary = make(SensitiveWorldLibary)
	}
	// 敏感词个数
	length := len(words)
	currentSWL := sensitiveWorldLibary

	for i := 0; i < length; i++ {

		wds := words[i]
		for _, value := range wds {
			word := string(value)
			if _, ok := currentSWL[word]; ok {

				currentSWL = currentSWL[word].(SensitiveWorldLibary)
			} else {

				newSWL := make(SensitiveWorldLibary)
				newSWL["isEnd"] = 0
				currentSWL[word] = newSWL
				currentSWL = newSWL
			}
		}
		currentSWL["isEnd"] = 1
		currentSWL = sensitiveWorldLibary
	}
	return sensitiveWorldLibary
}

// 检测铭感词
// 把一句话每个字转换成数组
// TODO matchType
func (sw SensitiveWorldLibary) CheckSensitiveWord(words []string, startIndex int, matchType MatchType) int {

	count := 0
	flag := false
	length := len(words)
	currentSWL := sw

	for i := startIndex; i < length-1; i++ {

		word := words[i]
		if _, ok := currentSWL[word]; ok {

			count++
			currentSWL = currentSWL[word].(SensitiveWorldLibary)
			if currentSWL["isEnd"] == 1 {
				flag = true
				if matchType == MinMartchType {
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
	// fmt.Println("count, ", count)
	return count
}
