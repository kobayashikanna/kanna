package main

type ServerConfig struct {
	ListenAddr string // 监听地址
	MTU int
	PeerTimeout int // 心跳超时
}
