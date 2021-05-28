package crypto

type dataCrypt interface {
	createHash() []byte
	Encrypt() error
	Decrypt() error
	GetEncryptRaw() string
	GetDecryptRaw() []byte
}

type cryptService interface {
	Encrypt() error
	Decrypt() error
	GetDecryptRaw() []byte
}

type UseCase interface {
	cryptService
}
