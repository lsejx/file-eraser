module eraser

replace example.com/me/hashed_random => ../hashed_random

replace example.com/me/filepath => ../filepath

go 1.19

require (
	example.com/me/filepath v0.0.0-00010101000000-000000000000
	example.com/me/hashed_random v0.0.0-00010101000000-000000000000
)

require (
	golang.org/x/crypto v0.5.0 // indirect
	golang.org/x/sys v0.4.0 // indirect
)
