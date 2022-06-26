package db

type DatabasePort interface {
	Get(resource string, collection string) (objs interface{}, err error)
	Save(resource string, collection string, objs interface{}) (err error)
}
