# jsontime

Parsing time in JSON compatible with some common format.

```go
import (
	"encoding/json"
	"github.com/bingoohuang/jsontime"
)

type Msg struct {
	O jsontime.Time
	A jsontime.Time
	B jsontime.Time
	C jsontime.Time
	E jsontime.Time
	F jsontime.Time
	D jsontime.Time `json:"d"`
}

func main() {
	j := `{
		"O": "",
		"A": "1604022846",
		"F": 1604022846000,
		"B": "2020-03-18 10:51:54.198",
		"C": "2020-03-18 10:51:54,198",
		"E": "2020-03-18T10:51:54,198",
		"d": "2020-03-18T10:51:54.198000Z",
		"G": "XYZ"
	}`

    var msg Msg
    err := json.Unmarshal([]byte(j), &msg)
}
```

supporting time format:

1. direct seconds/milliseconds/nanoseconds between now - 1 year and now + 1 year（整数的时间，是否在当前时间前后1年内的合法值（秒数/毫秒数/纳秒数，只处理前后1年的有效时间）
1. string format of millis seconds
1. string format of yyyy-MM-dd HH:mm:ss.SSS
1. string format of yyyy-MM-dd HH:mm:ss,SSS
1. string format of yyyy-MM-ddTHH:mm:ss.SSS
1. string format of yyyy-MM-ddTHH:mm:ss,SSS
1. string format of yyyy-MM-dd HH:mm:ss.SSSSSSZ
1. string format of yyyy-MM-dd HH:mm:ss,SSSSSSZ
1. 完整年月日时分秒的字符串
