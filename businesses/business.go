package businesses

type BaseParameter struct {
	ID      int
	Search  string
	OrderBy string
	Sort    string
	Page    int
	PerPage int
}

type Pagination struct {
	Page      int
	PerPage   int
	TotalData int
}

func (param *BaseParameter) GetOffset() int {
	return (param.GetPage() - 1) * param.GetPerPage()
}

func (param *BaseParameter) GetPage() int {
	if param.Page < 0 {
		param.Page = 1
	}
	return param.Page
}

func (param *BaseParameter) GetPerPage() int {
	if param.PerPage < 0 {
		param.PerPage = 10
	}
	return param.PerPage
}

func (param *BaseParameter) LikeChar(field string) string {
	return "%" + field + "%"
}

func (param *BaseParameter) GetPageInfo(totalData int) Pagination {
	return Pagination{
		Page:      param.GetPage(),
		PerPage:   param.GetPerPage(),
		TotalData: totalData,
	}
}
