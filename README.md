# jsontime

json time parser

Parsing time in JSON like 

```json
{
    "O": "",
    "A": "123",
    "F": 123,
    "B": "2020-03-18 10:51:54.198",
    "C": "2020-03-18 10:51:54,198",
    "E": "2020-03-18T10:51:54,198",
    "d": "2020-03-18T10:51:54.198000Z"
}
```

to relative go struct like

```go
type Msg struct {
	O jsontime.Time
	A jsontime.Time
	B jsontime.Time
	C jsontime.Time
	E jsontime.Time
	F jsontime.Time
	D jsontime.Time `json:"d"`
}
```

supporing time format:

1. direct milli seconds
1. string format of millis seconds
1. string format of yyyy-MM-dd HH:mm:ss.SSS
1. string format of yyyy-MM-dd HH:mm:ss,SSS
1. string format of yyyy-MM-ddTHH:mm:ss.SSS
1. string format of yyyy-MM-ddTHH:mm:ss,SSS
1. string format of yyyy-MM-dd HH:mm:ss.SSSSSSZ
1. string format of yyyy-MM-dd HH:mm:ss,SSSSSSZ
1. 完整年月日时分秒的字符串
