// Copyright 2018 gf Author(https://github.com/qnsoft/common). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/qnsoft/common.

// go test *.go -bench=".*"

package qn_rand_test

import (
	"crypto/rand"
	"encoding/binary"
	"testing"

	"github.com/qnsoft/common/util/qn_rand"
)

var (
	buffer         = make([]byte, 8)
	randBuffer4    = make([]byte, 4)
	randBuffer1024 = make([]byte, 1024)
	strForStr      = "我爱GoFrame"
)

func Benchmark_Rand_Buffer4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rand.Read(randBuffer4)
	}
}

func Benchmark_Rand_Buffer1024(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rand.Read(randBuffer1024)
	}
}

func Benchmark_qn_rand_Intn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		qn_rand.N(0, 99)
	}
}

func Benchmark_Perm10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		qn_rand.Perm(10)
	}
}

func Benchmark_Perm100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		qn_rand.Perm(100)
	}
}

func Benchmark_Rand_N1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		qn_rand.N(0, 99)
	}
}

func Benchmark_Rand_N2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		qn_rand.N(0, 999999999)
	}
}

func Benchmark_B(b *testing.B) {
	for i := 0; i < b.N; i++ {
		qn_rand.B(16)
	}
}

func Benchmark_S(b *testing.B) {
	for i := 0; i < b.N; i++ {
		qn_rand.S(16)
	}
}

func Benchmark_S_Symbols(b *testing.B) {
	for i := 0; i < b.N; i++ {
		qn_rand.S(16, true)
	}
}

func Benchmark_Str(b *testing.B) {
	for i := 0; i < b.N; i++ {
		qn_rand.Str(strForStr, 16)
	}
}

func Benchmark_Symbols(b *testing.B) {
	for i := 0; i < b.N; i++ {
		qn_rand.Symbols(16)
	}
}

func Benchmark_Uint32Converting(b *testing.B) {
	for i := 0; i < b.N; i++ {
		binary.LittleEndian.Uint32([]byte{1, 1, 1, 1})
	}
}

func Benchmark_Buffer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if _, err := rand.Read(buffer); err == nil {
			binary.LittleEndian.Uint64(buffer)
		}
	}
}
