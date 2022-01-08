package httpclient

type Option func(*httpOption)

type httpOption struct {
	withoutReqLog bool
	withoutResLog bool
}

func WithoutReqLog() Option {
	return func(ho *httpOption) {
		ho.withoutReqLog = true
	}
}

func WithoutResLog() Option {
	return func(ho *httpOption) {
		ho.withoutResLog = true
	}
}
