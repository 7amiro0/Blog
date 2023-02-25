package storage

type Blog struct {
	Author string `bson:"author"`
	Title  string `bson:"title"`
	Body   string `bson:"body"`
}

type Blogs struct {
	Blogs []Blog `json:"blogs"`
}

func (b Blogs) GetPosts() []Blog {
	return b.Blogs
} 