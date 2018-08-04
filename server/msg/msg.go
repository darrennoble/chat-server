package msg

import (
	"time"
)

type Msg struct {
	User    string
	Ts      time.Time
	Message string
}

type MsgList struct {
	msgs []*Msg
	pos  int
	len  int
}

func NewMsgList(capacity int) *MsgList {
	ml := &MsgList{
		msgs: make([]*Msg, capacity, capacity),
	}
	return ml
}

func (ml *MsgList) Add(m *Msg) {
	ml.msgs[ml.pos] = m
	ml.pos++
	if ml.pos >= cap(ml.msgs) {
		ml.pos = 0
	}
	if ml.len < cap(ml.msgs) {
		ml.len++
	}
}

func (ml MsgList) Get(i int) *Msg {
	if i > cap(ml.msgs) {
		return nil
	}

	if ml.len == 0 {
		return nil
	}

	pos := ml.pos - ml.len + i
	if pos < 0 {
		pos += cap(ml.msgs)
	}

	return ml.msgs[pos]
}

func (ml MsgList) Len() int {
	return ml.len
}
