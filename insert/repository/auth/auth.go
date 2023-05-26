package auth

type Auth interface {
	Ident() error
}
