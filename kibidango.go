package kibidango

import "errors"

type kibidango struct {
	id string
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
