package responses

/**
  The action was successfully received, understood, and accepted.
  200 - OK
  201 - Created
  204 - No Content
*/
type Response struct {
	StatusCode int         `form:"statusCode" json:"statusCode"`
	Message    string      `form:"message" json:"message"`
	Data       interface{} `form:"data" json:"data"`
}

/**
  Get Successfull Response object
*/
func (r Response) GetResponse() Response {
	return r
}

/**
  The request succeeded. The result meaning of "success" depends on the method:
  GET: The resource has been fetched and transmitted in the message body.
  HEAD: The representation headers are included in the response without any message body.
  PUT or POST: The resource describing the result of the action is transmitted in the message body.
  TRACE: The message body contains the request message as received by the server.

  Status Code: "200"
  Message: "OK"
*/
func (r *Response) Response_200(data interface{}) {
	r.StatusCode = 200
	r.Message = "OK"
	r.Data = data
}

/**
  The request succeeded, and a new resource was created as a result.
  This is typically the response sent after POST requests, or some PUT requests.

  Status Code: "201"
  Message: "Created | Updated"
*/
func (r *Response) Response_201(data interface{}) {
	r.StatusCode = 201
	r.Message = "Created | Updated | Deleted"
	r.Data = data
}

/**
  There is no content to send for this request, but the headers may be useful. This status code confirms that the server has fulfilled the request but does not need to return information.
  The user agent may update its cached headers for this resource with the new ones.
  Examples of this status code include delete requests or if a request was sent via a form and the response should not cause the form to be refreshed or for a new page to load.

  Status Code: "204"
  Message: "No Content"
*/
func (r *Response) Response_204(data interface{}) {
	r.StatusCode = 204
	r.Message = "No Content"
	r.Data = data
}

/**
  The server cannot or will not process the request due to something that is perceived to be a client error
  (e.g., malformed request syntax, invalid request message framing, or deceptive request routing).

  Status Code: "400"
  Message: "Bad Request"
*/
func (r *Response) Response_400(data interface{}) {
	r.StatusCode = 400
	r.Message = "Bad Request"
	r.Data = data
}

/**
  Although the standard specifies "unauthorized", semantically this response means "unauthenticated".
  That is, the client must authenticate itself to get the requested response.

  Status Code: "401"
  Message: "Unauthorized"
*/
func (r *Response) Response_401(data interface{}) {
	r.StatusCode = 401
	r.Message = "Unauthorized"
	r.Data = data
}

/**
  The client does not have access rights to the content; that is, it is unauthorized,
  so the server is refusing to give the requested resource. Unlike 401 Unauthorized,
  the client identity is known to the server.

  Status Code: "403"
  Message: "Forbidden"
*/
func (r *Response) Response_403(data interface{}) {
	r.StatusCode = 403
	r.Message = "Forbidden"
	r.Data = data
}

/**
  The server can not find the requested resource. In the browser, this means the URL is not recognized.
  In an API, this can also mean that the endpoint is valid but the resource itself does not exist.
  Servers may also send this response instead of '403 Forbidden' to hide the existence of a resource from an unauthorized client.
  This response code is probably the most well known due to its frequent occurrence on the web.

  Status Code: "404"
  Message: "Not Found"
*/
func (r *Response) Response_404(data interface{}) {
	r.StatusCode = 404
	r.Message = "Not Found"
	r.Data = data
}

/**
  The request method is known by the server but is not supported by the target resource.
  For example, an API may not allow calling DELETE to remove a resource.

  Status Code: "405"
  Message: "Method Not Allowed"
*/
func (r *Response) Response_405(data interface{}) {
	r.StatusCode = 405
	r.Message = "Method Not Allowed"
	r.Data = data
}

/**
  This response is sent when the web server, after performing server-driven content negotiation,
  doesn't find any content that conforms to the criteria given by the user agent.
*/
func (r *Response) Response_406(data interface{}) {
	r.StatusCode = 406
	r.Message = "Not Acceptable"
	r.Data = data
}

/**
  The server has encountered a situation it does not know how to handle.

  Status Code: "500"
  Message: "Internal Server Error"
*/
func (r *Response) Response_500(data interface{}) {
	r.StatusCode = 500
	r.Message = "Internal Server Error"
	r.Data = data
}

/**
  The request method is not supported by the server and cannot be handled.
  The only methods that servers are required to support (and therefore that
  must not return this code) are 'GET' and 'HEAD'.

  Status Code: "501"
  Message: "Not Implemented"
*/
func (r *Response) Response_501(data interface{}) {
	r.StatusCode = 501
	r.Message = "Not Implemented"
	r.Data = data
}
