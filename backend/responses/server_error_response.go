package responses

/*
  The server has encountered an error and failed to fulfill the request.
  500 - Internal Server Error
  501 - Created
*/
type ServerErrorResponse struct {
	Http BasicResponse `form:"response" json:"response"`
}

/*
  Get Server Error Response object
*/
func (s ServerErrorResponse) GetServerErrorResponse() ServerErrorResponse {
	return s
}

/*
  The server has encountered a situation it does not know how to handle.

  Status Code: "500"
  Message: "Internal Server Error"
*/
func (s *ServerErrorResponse) Response_500(data interface{}) {
	s.Http.StatusCode = 500
	s.Http.Message = "Internal Server Error"
}

/*
  The request method is not supported by the server and cannot be handled.
  The only methods that servers are required to support (and therefore that
  must not return this code) are 'GET' and 'HEAD'.

  Status Code: "501"
  Message: "Not Implemented"
*/
func (s *ServerErrorResponse) Response_501(data interface{}) {
	s.Http.StatusCode = 501
	s.Http.Message = "Not Implemented"
}
