package tlog

import (
	"strconv"
	"time"
)

type TypeLog struct {
	Error   string `json:"error"`
	Warning string `json:"warning"`
	Info    string `json:"info"`
}

type TransactionLog struct {
	Created string  `json:"created"`
	Message string  `json:"message"`
	Source  string  `json:"source"`
	Type    TypeLog `json:"type"`
}

func (t *TransactionLog) SetTransactionLog() {
	t.setTypeLog()
	t.createdTime()

}
func (t *TransactionLog) createdTime() {
	t.Created = strconv.Itoa(int(time.Now().Unix()))
}

func (t *TransactionLog) setTypeLog() {
	t.Type.Error = "Error"
	t.Type.Warning = "Warning"
	t.Type.Info = "Info"
}
