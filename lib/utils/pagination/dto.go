package pagination

type CursorPage[T any] struct {
	Data   []T    `json:"data"`
	Cursor string `json:"cursor"`
}
