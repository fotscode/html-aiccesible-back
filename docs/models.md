# Modelos

Los modelos pensados segun el prototipo inicial fueron los siguientes:

```go
type User struct{
	Username   string
	Password   string
	Posts      Post[]
	Comments   Comment[]
	Likes      Post[]
    Config     Configuration
	UpdatedAt  time.Time
	CreatedAt  time.Time
}

type Post struct{
	Title       string
	Description string
	Before      string
	After       string
	Likes       User[]
	Comments    Comment[]  
	UpdatedAt   time.Time
	CreatedAt   time.Time
}

type Comment struct{
	Title      string
	Content    string
	UpdatedAt  time.Time
	CreatedAt  time.Time
}

type Configuration struct{
	ShowLikes    bool
	ShowComments bool
	Theme        string
	Language     string
	SizeTitle    int
	SizeText     int
	User         User
	UpdatedAt    time.Time
	CreatedAt    time.Time
}
```

Como tecnologias se pensaba utilizar mysql como base de datos y gorm como ORM.
