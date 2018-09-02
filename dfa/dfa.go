package dfa

type MatchType uint8

const (
	MinMatchType MatchType = 1 << iota
	MaxMatchType
)

type SensitiveWorldLibrary map[string]interface{}

// 定义一个全家变量
var sensitiveWorldLibrary SensitiveWorldLibrary

/**
	 解析一个文件里的名词
	@Param {string} path 文件路径 支持 ini
	@Returns {[]string] 敏感词切片

*/
//func NewSensitiveLibary(path string) {
//
//}


// 初始化敏感词库
func NewSensitiveWorld(words []string) SensitiveWorldLibrary {

	if sensitiveWorldLibrary == nil {
		sensitiveWorldLibrary = make(SensitiveWorldLibrary)
	}
	// 敏感词个数
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

// 检测铭感词
// 把一句话每个字转换成数组
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
