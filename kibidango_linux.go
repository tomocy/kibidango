package kibidango

func ForLinux(id string) *Linux {
	return &Linux{
		kibidango: &kibidango{
			id: id,
		},
	}
}

type Linux struct {
	*kibidango
}
