package errors

type KeystoreError string

func (e KeystoreError) Error() string {
	return string(e)
}

type KeyError string

func (e KeyError) Error() string {
	return string(e)
}
