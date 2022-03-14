package responses

type HttpResponse struct {
	StatusCode int         `form:"statusCode" json:"statusCode"`
	Message    string      `form:"message" json:"message"`
	Data       interface{} `form:"data" json:"data"`
}
