package request

type CreateTagsRequest struct {
	Name string `validate:"required, min=3, max=10" json:"name"`
}
