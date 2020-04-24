package transmit

// Serializable is an interface describing structures that can be encoded and decoded
type Serializable interface {
	Serialize() ([]byte, error)
	Deserialize([]byte) error
}
