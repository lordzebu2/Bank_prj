package banking

// Account struct
type Account struct {
	//대문자여야 다른곳에서 사용 가능. public.
	// owner   string
	// balance int
	Owner   string
	Balance int
	//소문자는 private라서 다른곳에서 못씀..
}