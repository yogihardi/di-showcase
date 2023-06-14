package logs

type ZapAdapter interface {
}

type zapAdapter struct {
}

func NewZapAdapter() (ZapAdapter, error) {
	return &zapAdapter{}, nil
}
