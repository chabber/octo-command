package datastore

type DatabasePort interface {
	Get(resource string, collection string, obj interface{}) error
	GetAll(collection string, obj []interface{}) error
	Save(resource string, collection string, objs interface{}) (err error)
}
