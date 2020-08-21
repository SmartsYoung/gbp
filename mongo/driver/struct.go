package driver

type mgo struct {
	database   string
	collection string
}

func NewMgo(database, collection string) *mgo {
	return &mgo{
		database,
		collection,
	}
}
