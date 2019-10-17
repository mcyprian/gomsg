package datastore

import (
	"container/ring"
)

// MsgMemoryBuffer implements MsgDatastore interface storing messages in memory cyclic list
type MsgMemoryBuffer struct {
	buff *ring.Ring
}

func NewMsgMemoryBuffer(lenght int) *MsgMemoryBuffer {
	return &MsgMemoryBuffer{
		buff: ring.New(lenght),
	}
}

func (m *MsgMemoryBuffer) Insert(value interface{}) {
	m.buff.Value = value
	m.buff = m.buff.Next()
}

func (m *MsgMemoryBuffer) GetAll() []interface{} {
	var data []interface{}
	m.buff.Do(func(x interface{}) {
		if x != nil {
			data = append(data, x)
		}
	})
	return data
}
