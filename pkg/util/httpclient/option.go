package httpclient

type option func(*httpOption)

type httpOption struct {
	withoutReqLog bool
	withoutResLog bool
}

func WithoutReqLog() option {
	return func(ho *httpOption) {
		ho.withoutReqLog = true
	}
}

func WithoutResLog() option {
	return func(ho *httpOption) {
		ho.withoutResLog = true
	}
}
