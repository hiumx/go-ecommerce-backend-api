package response

const (
	ErrCodeSuccess      = 20001
	ErrCodeParamInvalid = 20002
)

var msg = map[int]string{
	ErrCodeSuccess:      "Success",
	ErrCodeParamInvalid: "Param invalid!",
}
