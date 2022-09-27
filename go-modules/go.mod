module go-modules

go 1.17

replace go-modules/math => ./math

require (
	github.com/google/uuid v1.0.0 // indirect
	github.com/pborman/uuid v1.2.1 // indirect
	go-modules/math v0.0.0-00010101000000-000000000000 // indirect
)
