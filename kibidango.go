package kibidango

type Kibidango struct{}

func (c *Kibidango) Create(creater Creater, args ...string) error {
	return create(creater, args...)
}

func create(creater Creater, args ...string) error {
	return creater.Create(args...)
}

type Creater interface {
	Create(args ...string) error
}

func (c *Kibidango) Init(initer Initializer) error {
	return initialize(initer)
}

func initialize(initer Initializer) error {
	return initer.Init()
}

type Initializer interface {
	Init() error
}
