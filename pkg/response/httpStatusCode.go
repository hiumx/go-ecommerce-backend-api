package response

const (
	ErrCodeSuccess      = 20001
	ErrCodeParamInvalid = 20002
	ErrInvalidToken     = 300001
)

var msg = map[int]string{
	ErrCodeSuccess:      "Success",
	ErrCodeParamInvalid: "Param invalid!",
	ErrInvalidToken:     "Invalid token!",
}
