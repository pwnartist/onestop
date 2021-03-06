package database

type OnestopDatabaseContext interface {
	Session() interface{}
	Insert(OnestopDataContext) error
	Select(OnestopDataContext) ([]interface{}, error)
}

type OnestopDataContext interface {
	JsonString() (string, error)
	Parse(interface{}) ([]interface{}, error)
	Conditions() interface{}
	TableName() string
	Data() interface{}
}
