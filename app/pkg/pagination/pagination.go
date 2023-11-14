package pagination

type PageInfo struct {
	PageSize int         `json:"pageSize"`
	PageNum  int         `json:"pageNum"`
	Total    int         `json:"total"`
	LastPage int         `json:"lastPage"`
	Order    string      `json:"order"`
	Offset   int         `json:"offset"`
	Data     interface{} `json:"data"`
}

func PageHelper(pageNum, pageSize int, order string, total int) PageInfo {
	var page PageInfo
	var lastPage int
	if pageNum == 0 {
		pageNum = 1
	}
	if pageSize == 0 {
		pageSize = 10
	}
	//total := int(totalNum)
	page.Total = total
	if total == 0 {
		lastPage = 1
	} else {
		if total%pageSize > 0 {
			lastPage = total/pageSize + 1
		} else {
			lastPage = total / pageSize
		}
	}
	page.LastPage = lastPage
	if pageNum < 1 {
		page.PageNum = 1
	} else if pageNum > lastPage {
		//page.PageNum = lastPage
		page.PageNum = pageNum
	} else {
		page.PageNum = pageNum
	}
	//单页最大20
	if pageSize > 20 {
		page.PageSize = 20
	} else if pageSize < 3 {
		//单页最小3
		page.PageSize = 3
	} else {
		page.PageSize = pageSize
	}
	//offset
	page.Offset = (page.PageNum - 1) * page.PageSize
	//order
	//if order == "DESC" || order == "desc" {
	//		page.Order = "DESC"
	//}
	if order == "ASC" || order == "asc" {
		page.Order = "ASC"
	} else {
		page.Order = "DESC"
	}

	return page

}
