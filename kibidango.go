package kibidango

import "errors"

type kibidango struct {
	id   string
	root string
}

func (k *kibidango) ID() string {
	return k.id
}

func (k *kibidango) UpdateID(id string) error {
	if id == "" {
		return errors.New("id should not be empty")
	}

	k.id = id

	return nil
}

func (k *kibidango) UpdateRoot(root string) error {
	k.root = root
	return nil
}
