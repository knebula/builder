package tmpl

var Kbt = `// Code generated by kbt:{{.Version}}. DO NOT EDIT.
//go:build mage
// +build mage

package main

// Install tools
func Install() error {
	return nil
}

// Generate source code
func Generate() error {
	return nil
}

// Test code
func Test() error {
	return nil
}

// Build cmds
func Build() error {
	return nil
}

// Package cmds in Docker images
func Package() error {
	return nil
}

// Deploy cmds Docker images to Kubernetes
func Deploy() error {
	return nil
}

`
