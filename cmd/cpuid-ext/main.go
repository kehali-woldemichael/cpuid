// Copyright (c) 2021 Klaus Post, released under MIT License. See LICENSE file.

// Package cpuid provides information about the CPU running the current program.
//
// CPU features are detected on startup, and kept for fast access through the life of the application.
// Currently x86 / x64 (AMD64) as well as arm64 is supported.
//
// You can access the CPU information by accessing the shared CPU variable of the cpuid library.
//
// Package home: https://github.com/klauspost/cpuid
package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/klauspost/cpuid/v2"
)


func main() {
	cpuid.Flags()
	flag.Parse()
	cpuid.Detect()
	//cpuid.DetectARM()
	// Test if we have these specific features:

	fmt.Println("Features:", strings.Join(cpuid.CPU.FeatureSet(), ","))
	fmt.Println("Model:", cpuid.CPU.Model)
}
