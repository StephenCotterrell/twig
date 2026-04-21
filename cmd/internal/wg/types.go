// Package wg provides utilities to access the wg and wg-quick interfaces
package wg

type ProfileState struct {
	Profile   Profile
	IsActive  bool
	Interface *InterfaceStatus
}

type Profile struct {
	Name string
	Path string
}

type InterfaceStatus struct {
	Name         string
	ListenPort   int
	FirewallMark string
	PublicKey    string
	Peers        []PeerStatus
}

type PeerStatus struct {
	PublicKey         string
	Endpoint          string
	AllowedIPs        []string
	LatestHandshake   string
	ReceivedBytes     int64
	TransmitBytes     int64
	KeepAliveInterval string
}
