package literature

type Literature interface {
	Read() bool
	GetAuthor() string
	GetTitle() string
}
