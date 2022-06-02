package response

import (
	"fmt"
	"time"
)

type JsonTime time.Time

func (j JsonTime) MarshalJSON() ([]byte, error) {
	var stamp = fmt.Sprintf("\"%s\"", time.Time(j).Format("2006-01-02"))
	return []byte(stamp), nil
}

type User struct {
	Id       uint64   `json:"id"`
	Mobile   string   `json:"mobile"`
	NickName string   `json:"nickName"`
	BirthDay JsonTime `json:"birthDay"`
	Gender   uint32   `json:"gender"`
	Role     uint32   `json:"role"`
}
