package errors

type ErrorObject interface {
	error
	Type() ErrorType
	Set(string, interface{}) ErrorObject
	Get(string) interface{}
	Expand(bool) ErrorObject
	JSON() map[string]interface{}
}

type errorObject struct {
	errorType ErrorType
	text      string
	error     map[string]interface{}
	expand    bool
}

func (e *errorObject) Error() string {
	return e.text
}

func (e *errorObject) JSON() map[string]interface{} {
	json := map[string]interface{}{"type": e.errorType, "text": e.text, "expand": e.expand}
	if e.expand {
		json["error"] = e.error
	}
	return json
}

func (e *errorObject) Set(k string, v interface{}) ErrorObject {
	if v != nil {
		e.error[k] = v
	}
	return e
}

func (e *errorObject) Get(k string) interface{} {
	if v, ok := e.error[k]; ok {
		return v
	}
	return nil
}

func (e *errorObject) Expand(b bool) ErrorObject {
	e.expand = b
	return e
}

func (e *errorObject) Type() ErrorType {
	return e.errorType
}

func newErrorObject(errorType ErrorType, text string) ErrorObject {
	return &errorObject{errorType: errorType, text: text, error: make(map[string]interface{}, 0), expand: false}
}
