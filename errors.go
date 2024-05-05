package goutil

type HttpResponseError struct {
	Code            int
	Message         string
	ResponseBodyRaw *[]byte
	ResponseBody    interface{}
}

func (e HttpResponseError) Error() string {
	return e.Message
}
