package link

type CreateRequest struct {
	Url string `json:"url" validate:"required,url"`
}

type UpdateRequest struct {
	Url  string `json:"url" validate:"required,url"`
	Hash string `json:"hash"`
}

type GetLinkResponse struct {
	Links []Link `json:"links"`
	Total int64  `json:"total"`
}
