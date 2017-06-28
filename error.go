package sphinx

import "fmt"

var (
	// ErrReplayedPacket is an error returned when a packet is rejected
	// during processing due to being an attempted replay or probing
	// attempt.
	ErrReplayedPacket = fmt.Errorf("sphinx packet replay attempted")

	// ErrInvalidOnionVersion is returned during decoding of the onion packet, when
	// received packet has unknown version byte.
	ErrInvalidOnionVersion = fmt.Errorf("invalid onion packet version")

	// ErrInvalidOnionHMAC is returned during onion parsing process, when received
	// mac do not corresponds to the generated one.
	ErrInvalidOnionHMAC = fmt.Errorf("invalid mismathed mac")

	// ErrInvalidOnionKey is returned during onion parsing process, when
	// onion key is invalid - not belongs to the secp256k1 curve.
	ErrInvalidOnionKey = fmt.Errorf("invalid onion key: pubkey isn't on " +
		"secp256k1 curve")
)
