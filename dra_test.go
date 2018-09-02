package DFA

import (
	"encoding/json"
	"reflect"
	"strings"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewSensitiveWorld(t *testing.T) {

	// Only pass t into top-level Convey calls
	Convey("初始化初始化敏感词库", t, func() {
		var words = []string{
			"中国人",
			"黑人",
			"江泽民",
			"胡锦涛",
			"胡锦涛主席",
		}
		sensitiveWorldLibary := NewSensitiveWorld(words)
		a, _ := json.Marshal(sensitiveWorldLibary)
		var res = `{"中":{"isEnd":0,"国":{"isEnd":0,"人":{"isEnd":1}}},"江":{"isEnd":0,"泽":{"isEnd":0,"民":{"isEnd":1}}},"胡":{"isEnd":0,"锦":{"isEnd":0,"涛":{"isEnd":1,"主":{"isEnd":0,"席":{"isEnd":1}}}}},"黑":{"isEnd":0,"人":{"isEnd":1}}}`
		So(string(a), ShouldEqual, res)

	})
}

func TestCheckSensitiveWord(t *testing.T) {
	Convey("测试一句话", t, func() {
		str := "我们以前的主席是江泽民, 现在的主席是胡锦涛主席， 黑人不应该直接叫人，这样子是不礼貌的中国人的"
		// strLength := utf8.RuneCountInString(str)
		var strArr []string

		for _, value := range str {
			strArr = append(strArr, string(value))
		}

		var sensitivew []string
		for i := 0; i < len(strArr); i++ {
			if count := sensitiveWorldLibary.CheckSensitiveWord(strArr, i, MinMartchType); count > 0 {
				sensitivew = append(sensitivew, strings.Join(strArr[i:i+count], ""))
			}
		}
		res := []string{
			"江泽民",
			"胡锦涛",
			"黑人",
			"中国人",
		}

		is := reflect.DeepEqual(res, sensitivew)
		So(is, ShouldEqual, true)
	})
}
