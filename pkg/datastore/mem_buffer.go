package datastore

import (
	"container/ring"
)

// MsgMemoryBuffer implements MsgDatastore interface storing messages in memory cyclic list
type MsgMemoryBuffer struct {
	buff *ring.Ring
}

// NewMsgMemoryBuffer constructs a buffer of desired length
func NewMsgMemoryBuffer(lenght int) *MsgMemoryBuffer {
	return &MsgMemoryBuffer{
		buff: ring.New(lenght),
	}
}

// Insert adds a new value to the buffer
func (m *MsgMemoryBuffer) Insert(value interface{}) {
	m.buff.Value = value
	m.buff = m.buff.Next()
}

// GetAll returns slice of all not nil data stored in the buffer
func (m *MsgMemoryBuffer) GetAll() []interface{} {
	var data []interface{}
	m.buff.Do(func(x interface{}) {
		if x != nil {
			data = append(data, x)
		}
	})
	return data
}
