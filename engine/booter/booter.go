package booter

func ForOS(os string, root string) Booter {
	switch os {
	case "linux":
		return ForLinux(root)
	default:
		return nil
	}
}

type Booter interface {
	Boot() error
}
