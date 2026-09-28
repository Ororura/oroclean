package model

type Category string

const (
	CategoryGeneric   Category = "generic"
	CategoryCache     Category = "cache"
	CategoryLogs      Category = "logs"
	CategoryDownloads Category = "downloads"
	CategoryTrash     Category = "trash"
	CategoryDeveloper Category = "developer"
)

func (c Category) Valid() bool {
	switch c {
	case
		CategoryGeneric,
		CategoryCache,
		CategoryLogs,
		CategoryDownloads,
		CategoryTrash,
		CategoryDeveloper:
		return true
	default:
		return false
	}
}
