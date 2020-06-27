package sora

import "github.com/pion/webrtc/v2"

// ConnectionOptions は Sora 接続設定です。
type ConnectionOptions struct {
	// Sora の URL
	SoraURL string

	// Role (sendonly | sendrecv | recvonly)
	Role string

	// Channel ID
	ChannelID string

	// Video の設定
	Video *webrtc.RTPCodec

	// Audio の設定
	Audio bool

	// Metadata
	Metadata Metadata

	// TrickleICE を利用するかどうかのフラグ
	UseTrickeICE bool

	// Debug 出力をするかどうかのフラグ
	Debug bool
}
