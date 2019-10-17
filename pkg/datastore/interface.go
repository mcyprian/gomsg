package datastore

// MsgDatastore defines API for message datastores
type MsgDatastore interface {
	Insert(interface{})
	GetAll() []interface{}
}
