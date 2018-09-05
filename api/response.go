package api

type Response interface{}

type OkResponse struct {
	Data interface{} `json:"data"`
}

type BadResponse struct {
	Err ErrResponse `json:"err"`
}

type ServerErrorResponse struct {
	Err ErrResponse `json:"err"`
}

type ErrResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func (e *ErrResponse) Error() string {
	return "Error when processing your request: " + e.Msg
}
