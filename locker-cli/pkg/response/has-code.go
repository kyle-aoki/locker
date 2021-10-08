package response

type HasCode interface {
	getCode() string
}

func (r BasicResponse) getCode() string {
	return r.Code
}
func (r MessageResponse) getCode() string {
	return r.Code
}
func (r StrResponse) getCode() string {
	return r.Code
}
func (r ListResponse) getCode() string {
	return r.Code
}
func (r KeyValueResponse) getCode() string {
	return r.Code
}
