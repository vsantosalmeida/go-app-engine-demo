package crypto

type DataCrypt interface {
	createHash() []byte
	Encrypt() error
	Decrypt() error
	GetRaw() []byte
}

type CryptService interface {
	Encrypt() error
	Decrypt() error
	GetRaw() []byte
}

type UseCase interface {
	CryptService
}
