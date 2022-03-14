package responses

/**
  Basic response for HTTP request
  Include "Status Code" & "Message"
*/
type BasicResponse struct {
	StatusCode int    `form:"statusCode" json:"statusCode"`
	Message    string `form:"message" json:"message"`
}
