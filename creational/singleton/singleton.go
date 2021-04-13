package singleton

//type Book interface {
//	SetName() string
//	GetName() string
//	SetId() int
//	GetId() int
//}

type book struct {
	name string
	id   int
}

var instance *book

func GetBookInstance() *book {
	if instance == nil {
		instance = new(book)
	}

	return instance
}

func (b *book) SetName(name string) string {
	b.name = name
	return b.name
}

func (b *book) SetId(id int) int {
	b.id = id
	return b.id
}

func (b *book) GetName() string {
	return b.name
}

func (b *book) GetId() int {
	return b.id
}