package outbox

type Message interface {
	GetData() []byte
}