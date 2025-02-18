package interceptors

func interceptionError(req interface{}, errCode, errMsg string) interface{} {

	switch req.(type) {

	default:
		return nil
	}
}
