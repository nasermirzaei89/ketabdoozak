package templates

type Head struct {
	Title string
	Meta  []Meta
}

const (
	MetaNameAuthor             = "author"
	MetaNameDescription        = "description"
	MetaNameKeywords           = "keywords"
	MetaNameOGTitle            = "og:title"
	MetaNameOGImage            = "og:image"
	MetaNameOGDescription      = "og:description"
	MetaNameTwitterTitle       = "twitter:title"
	MetaNameTwitterImage       = "twitter:image"
	MetaNameTwitterDescription = "twitter:description"
)

type Meta struct {
	Name    string
	Content string
}

func (head Head) PageTitle() string {
	if head.Title == "" {
		return "کتابدوزک"
	}

	return head.Title + " | کتابدوزک"
}

func EmptyHead() Head {
	return Head{
		Title: "",
		Meta:  nil,
	}
}

func ErrorHead() Head {
	return Head{
		Title: "خطا",
		Meta:  nil,
	}
}
