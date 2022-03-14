package responses

/*
  An error may have been caused by the client. The request contains bad syntax or cannot be fulfilled.
  400 - Bad Request
  401 - Unauthorized
  403 - Forbidden
  404 - Not Found
  405 - Method Not Allowed
*/
type ClientErrorResponse struct {
	Http BasicResponse `form:"response" json:"response"`
}

/*
  Get Client Error Response object
*/
func (c ClientErrorResponse) GetClientErrorResponse() ClientErrorResponse {
	return c
}

/*
  The server cannot or will not process the request due to something that is perceived to be a client error
  (e.g., malformed request syntax, invalid request message framing, or deceptive request routing).

  Status Code: "400"
  Message: "Bad Request"
*/
func (c *ClientErrorResponse) Response_400() {
	c.Http.StatusCode = 400
	c.Http.Message = "Bad Request"
}

/*
  Although the HTTP standard specifies "unauthorized", semantically this response means "unauthenticated".
  That is, the client must authenticate itself to get the requested response.

  Status Code: "401"
  Message: "Unauthorized"
*/
func (c *ClientErrorResponse) Response_401() {
	c.Http.StatusCode = 401
	c.Http.Message = "Unauthorized"
}

/*
  The client does not have access rights to the content; that is, it is unauthorized,
  so the server is refusing to give the requested resource. Unlike 401 Unauthorized,
  the client identity is known to the server.

  Status Code: "403"
  Message: "Forbidden"
*/
func (c *ClientErrorResponse) Response_403() {
	c.Http.StatusCode = 403
	c.Http.Message = "Forbidden"
}

/*
  The server can not find the requested resource. In the browser, this means the URL is not recognized.
  In an API, this can also mean that the endpoint is valid but the resource itself does not exist.
  Servers may also send this response instead of '403 Forbidden' to hide the existence of a resource from an unauthorized client.
  This response code is probably the most well known due to its frequent occurrence on the web.

  Status Code: "404"
  Message: "Not Found"
*/
func (c *ClientErrorResponse) Response_404() {
	c.Http.StatusCode = 404
	c.Http.Message = "Not Found"
}

/*
  The request method is known by the server but is not supported by the target resource.
  For example, an API may not allow calling DELETE to remove a resource.

  Status Code: "405"
  Message: "Method Not Allowed"
*/
func (c *ClientErrorResponse) Response_405() {
	c.Http.StatusCode = 405
	c.Http.Message = "Method Not Allowed"
}
