package webmania

type ErrorMessage struct {
	Err string `json:"error"`
}

func (e ErrorMessage) Error() string {
	return e.Err
}

type SuccessMessage struct {
	Success string `json:"success"`
}
