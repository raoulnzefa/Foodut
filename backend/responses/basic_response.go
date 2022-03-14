package responses

type BasicResponse struct {
	StatusCode int    `form:"statusCode" json:"statusCode"`
	Message    string `form:"message" json:"message"`
}
