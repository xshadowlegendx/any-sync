package list

type ACLRecord struct {
	Id                 string
	PrevId             string
	CurrentReadKeyHash uint64
	Timestamp          int64
	Data               []byte
	Identity           []byte
	Model              interface{}
	Signature          []byte
}