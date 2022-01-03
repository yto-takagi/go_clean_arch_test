package response

type AuthorGetByAllResponse struct {
	Author  []Author `json:"authors"`
	ExpPool ExpPool  `json:"exp_pool"`
}

func NewAuthorGetByAllResponse(authors []Author, exp_pool ExpPool) *AuthorGetByAllResponse {
	AuthorGetByAllResponse := new(AuthorGetByAllResponse)
	AuthorGetByAllResponse.Author = authors
	AuthorGetByAllResponse.ExpPool = exp_pool
	return AuthorGetByAllResponse
}
