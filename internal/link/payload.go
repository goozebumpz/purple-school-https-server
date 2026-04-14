package link

type CreateRequest struct {
	Url string `json:"url" validate:"required,url"`
}
