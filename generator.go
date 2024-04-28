package main

import "strings"

type Generator interface {
	Generate(req string) error
}

type GenCreator func(*strings.Builder) Generator

var (
	G = map[string]GenCreator{}
)
