package defs


type Err struct {
	Error string `json:"error"`
	ErrorCode string `json:"error_code"'`
}

type ErrorResponse struct {
	HttpSc int
	Error Err
}
var (
	ErrorRequestBodyParseFailed =ErrorResponse{HttpSc: 400,Error: Err{"Request body is not correct", "001"}}
    ErrorNotAuthUser = ErrorResponse{401,Err{Error: "User authentication failed", ErrorCode: "002"}}
	)
