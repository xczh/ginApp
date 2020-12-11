package router

// 客户端错误：输入参数格式校验错误
const CLIENT_ERROR_VALIDATION = "invalid"

// 客户端错误：输入参数格式正确，但请求无法处理
const CLIENT_ERROR_UNPROCESSABLE = "unprocessable"

type ResponseMessage struct {
	Message string `json:"message"`
}
