package answer

const (
	responseHeaderJSON = "application/json; charset=utf-8"
)

type key int

const (
	keyUserEmail       key = iota
	keyQuestionID      key = iota
	keyNewQuestion     key = iota
	keyUpdatedQuestion key = iota
	keyNewAnswer       key = iota
)
