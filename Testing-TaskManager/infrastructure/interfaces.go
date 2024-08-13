package infrastructure

type Services interface {
	Encode(id string , email string , is_admin bool) (string , error)
	HashPassWord(spass string) string
}