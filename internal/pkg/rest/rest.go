package rest

type Consumer interface {
}

type consumer struct {
}

func NewConsumer() (Consumer, error) {
	return &consumer{}, nil
}
