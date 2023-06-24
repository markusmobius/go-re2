//go:build re2_cgo

package cre2

/*
#cgo pkg-config: re2
#cgo LDFLAGS: -L. -lstdc++
#cgo CXXFLAGS: -std=c++17
*/
import "C"
