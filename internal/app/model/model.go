package model

import "time"

type (
	ProxyType          uint8
	ProxySEcurityLevel uint8

	IProxyDesc interface {
		Country() string
		Type() ProxyType
		Addres() string
		Port() int
		User() string
		Passwd() string
	}
	IProxyItem interface {
		LastChecked() time.Time
		Speed() int
		Security() ProxySEcurityLevel
	}
)

// Proxy types definitions
const (
	ProxyTypeHTTP ProxyType = iota + 1
	ProxyTypeHTTPS ProxyType
	PoxyTypeSocks4 ProxyType
	ProxyTypeSocks4A ProxyType
	ProxyTypeSocks5 ProxyType
)
//
const (
	SecurityLevelLow ProxySEcurityLevel = iota +1
	SecurityLevelAnonim ProxySEcurityLevel
	SecurtyLeveHigh ProxySEcurityLevel
)

//
