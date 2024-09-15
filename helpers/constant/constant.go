package constant

type Dialect string

func (d Dialect) String() string {
	return string(d)
}

const (
	PostgresDialect Dialect = "postgres"
)

type TableName string

func (t TableName) String() string {
	return string(t)
}

const (
	UserTableName    TableName = "users"
	PostTableName    TableName = "posts"
	FollowTableName  TableName = "follows"
	CommentTableName TableName = "comments"
)

type DateTimeFormat string

func (d DateTimeFormat) String() string {
	return string(d)
}

const (
	DateFormat DateTimeFormat = "2006-02-01"
)

type RegexFormat string

func (d RegexFormat) String() string {
	return string(d)
}
