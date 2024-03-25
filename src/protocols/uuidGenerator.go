package protocols

type UUIDGenerator interface {
	Generate() string
	IsValidUUID(id string) error
}
