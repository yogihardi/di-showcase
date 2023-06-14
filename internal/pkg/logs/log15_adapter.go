package logs

type Log15Adapter interface {
}

type log15Adapter struct {
}

func NewLog15Adapter() (Log15Adapter, error) {
	return &log15Adapter{}, nil
}
