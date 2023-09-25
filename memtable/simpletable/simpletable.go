package simpletable

import "errors"

type SimpleTable struct {
	data map[string]string
}

func (st *SimpleTable) Insert(key, value string) error {
	st.data[key] = value
	return nil
}

func (st *SimpleTable) Remove(key string) error {
	if _, ok := st.data[key]; !ok {
		return errors.New("key doesnt exist")
	}

	delete(st.data, key)
	return nil
}

func (st *SimpleTable) Get(key string) (string, error) {
	if value, ok := st.data[key]; ok {
		return value, nil
	}

	return "", errors.New("key doesnt exist")
}
