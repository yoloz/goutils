package handler

type Handler interface {
	Encrypt(key string) (cypher string, err error)
	Decrypt(key string) (text string, err error)
}
