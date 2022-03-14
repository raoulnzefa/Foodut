package responses

// /**
//   The action was successfully received, understood, and accepted.
//   200 - OK
//   201 - Created
//   204 - No Content
// */
// type SuccessfullResponse struct {
// 	Http BasicResponse `form:"http" json:"http"`
// 	Data interface{}   `form:"data" json:"data"`
// }

// /**
//   Get Successfull Response object
// */
// func (s SuccessfullResponse) GetSuccessfullResponse() SuccessfullResponse {
// 	return s
// }

// /**
//   The request succeeded. The result meaning of "success" depends on the HTTP method:
//   GET: The resource has been fetched and transmitted in the message body.
//   HEAD: The representation headers are included in the response without any message body.
//   PUT or POST: The resource describing the result of the action is transmitted in the message body.
//   TRACE: The message body contains the request message as received by the server.

//   Status Code: "200"
//   Message: "OK"
// */
// func (s *SuccessfullResponse) Response_200(data interface{}) {
// 	s.Http.StatusCode = 200
// 	s.Http.Message = "OK"
// 	s.Data = data
// }

// /**
//   The request succeeded, and a new resource was created as a result.
//   This is typically the response sent after POST requests, or some PUT requests.

//   Status Code: "201"
//   Message: "Created"
// */
// func (s *SuccessfullResponse) Response_201() {
// 	s.Http.StatusCode = 201
// 	s.Http.Message = "Created"
// }

// /**
//   There is no content to send for this request, but the headers may be useful. This status code confirms that the server has fulfilled the request but does not need to return information.
//   The user agent may update its cached headers for this resource with the new ones.
//   Examples of this status code include delete requests or if a request was sent via a form and the response should not cause the form to be refreshed or for a new page to load.

//   Status Code: "204"
//   Message: "No Content"
// */
// func (s *SuccessfullResponse) Response_204() {
// 	s.Http.StatusCode = 201
// 	s.Http.Message = "No Content"
// }
