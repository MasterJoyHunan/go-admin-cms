package response

type ArticleCatePage struct {
	Total       int           `json:"total"`        // 总共多少页
	PerPage     int           `json:"per_page"`     // 当前页码
	CurrentPage int           `json:"current_page"` // 每页显示多少条
	Data        []ArticleCate `json:"data"`
}

type ArticleCate struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Sort    int    `json:"sort"`
	Desc    string `json:"desc"`
	KeyWord string `json:"key_word"`
}
