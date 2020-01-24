package hop

// Network indicates the blockchain network that is intended to be the next hop
// for a forwarded HTLC. The existence of this field within the ForwardingInfo
// struct enables the ability for HTLC to cross chain-boundaries at will.
type Network uint8

const (
	// BitcoinNetwork denotes that an HTLC is to be forwarded along the
	// Bitcoin link with the specified short channel ID.
	BitcoinNetwork Network = iota

	// MonacoinNetwork denotes that an HTLC is to be forwarded along the
	// Monacoin link with the specified short channel ID.
	MonacoinNetwork
)

// String returns the string representation of the target Network.
func (c Network) String() string {
	switch c {
	case BitcoinNetwork:
		return "Bitcoin"
	case MonacoinNetwork:
		return "Monacoin"
	default:
		return "Kekcoin"
	}
}
