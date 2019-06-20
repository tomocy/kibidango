package initializer

func ForOS(os string, root string) Initializer {
	switch os {
	case "linux":
		return ForLinux(root)
	default:
		return nil
	}
}

type Initializer interface {
	Init() error
}
