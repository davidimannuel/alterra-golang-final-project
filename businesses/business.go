package businesses

type BaseParameter struct {
	ID      int
	Search  string
	OrderBy string
	Sort    string
	Page    int
	PerPage int
}

func (param *BaseParameter) GetOffset() int {
	return (param.Page - 1) * param.PerPage
}

func (param *BaseParameter) LikeChar(field string) string {
	return "%" + field + "%"
}
