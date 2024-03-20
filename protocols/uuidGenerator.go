package protocols

type UUIDGenerator interface {
	Generate() string
}
