package test

import (
	"project/binar/module"
	"testing"
)

func TestModule(t *testing.T) {
	var data = []string{
		"Dana",
		"Nada",
		"Nadia",
		"Ami",
		"Uni",
	}
	t.Run("Test Login", func(t *testing.T) {
		for _, item := range data {
			module.PlaYGame(item)
		}
	})
}
