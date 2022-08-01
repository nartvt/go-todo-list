package response

type Pagination struct {
	Data  interface{}
	Limit int
	Page  int
	Total int
}
