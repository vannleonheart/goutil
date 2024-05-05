package goutil

type HttpResponseError struct {
	Code         int
	Message      string
	ResponseBody *[]byte
}

func (e HttpResponseError) Error() string {
	return e.Message
}
