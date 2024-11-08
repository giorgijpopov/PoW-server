package api

type Response struct {
	Quote string `json:"quote"`
	Error string `json:"error"`
}

func newErrorResponse(err error) Response {
	return Response{Error: err.Error()}
}

func newSuccessResponse(quote string) Response {
	return Response{Quote: quote}
}
