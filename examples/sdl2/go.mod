module github.com/hakobera/go-sora/examples/sdl2

go 1.14

replace github.com/hakobera/go-sora v0.3.0 => ../../../go-sora

require (
	github.com/hakobera/go-sora v0.3.0
	github.com/hakobera/go-webrtc-decoder v0.3.0
	github.com/pion/rtp v1.6.0
	github.com/pion/webrtc/v3 v3.0.0-beta.1
	github.com/veandco/go-sdl2 v0.4.4
)
