// Package wgtypes provides shared types for the wireguardctrl family
// of packages.
package wgtypes

import (
	"encoding/base64"
	"fmt"
	"net"
	"time"
)

// A Device is a WireGuard device.
type Device struct {
	Index        int
	Name         string
	PrivateKey   Key
	PublicKey    Key
	ListenPort   int
	FirewallMark int
	Peers        []Peer
}

const keyLen = 32 // wgh.KeyLen

// A Key is a public or private key.
type Key [keyLen]byte

// NewKey creates a Key from a byte slice.  The byte slice must be exactly
// 32 bytes in length.
func NewKey(b []byte) (Key, error) {
	if len(b) != keyLen {
		return Key{}, fmt.Errorf("wireguardnl: incorrect key size: %d", len(b))
	}

	var k Key
	copy(k[:], b)

	return k, nil
}

// MustKey calls NewKey, but panics if an error occurs.
func MustKey(b []byte) Key {
	k, err := NewKey(b)
	if err != nil {
		panic(err)
	}

	return k
}

// String returns the base64 string representation of a Key.
func (k Key) String() string {
	return base64.StdEncoding.EncodeToString(k[:])
}

// A Peer is a WireGuard peer to a Device.
type Peer struct {
	PublicKey                   Key
	PresharedKey                Key
	Endpoint                    *net.UDPAddr
	PersistentKeepaliveInterval time.Duration
	LastHandshakeTime           time.Time
	ReceiveBytes                int
	TransmitBytes               int
	AllowedIPs                  []net.IPNet
}