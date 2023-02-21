package storage

type Blog struct {
	Author string `bson:"author"`
	Title  string `bson:"title"`
	Body   string `bson:"body"`
}
