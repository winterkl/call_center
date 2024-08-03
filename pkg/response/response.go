package response

import "contact_center/pkg/postgres/utils/paginate"

type Response struct {
	Message  string             `json:"message"`
	Paginate *paginate.Paginate `json:"paginate,omitempty"`
	Data     interface{}        `json:"data,omitempty"`
}

func New(message string) *Response {
	return &Response{Message: message}
}

func (m *Response) SetData(data interface{}) *Response {
	m.Data = data
	return m
}

func (m *Response) SetPaginate(paginate *paginate.Paginate) *Response {
	m.Paginate = paginate
	return m
}
