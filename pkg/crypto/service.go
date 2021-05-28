package crypto

type Service struct {
	crypt dataCrypt
}

func NewService(crypt dataCrypt) UseCase {
	return &Service{
		crypt: crypt,
	}
}

func (s *Service) Encrypt() error {
	err := s.crypt.Encrypt()
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Decrypt() error {
	err := s.crypt.Decrypt()
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) GetDecryptRaw() []byte {
	return s.crypt.GetDecryptRaw()
}
