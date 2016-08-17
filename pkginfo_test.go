// Copyright 2016 Lennart Espe. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.md file.

// Package pkginfo provides utilities to store package metadata.
package pkginfo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPackageVersion(t *testing.T) {
	testCases := []struct {
		Subject  PackageVersion
		Expected string
	}{
		{
			PackageVersion{Major: 0, Minor: 0, Patch: 4},
			"0.0.4",
		},
		{
			PackageVersion{Major: 0, Minor: 1, Patch: 0},
			"0.1.0",
		},
		{
			PackageVersion{Major: 0, Minor: 13, Patch: 2},
			"0.13.2",
		},
		{
			PackageVersion{Major: 1, Minor: 0, Patch: 0},
			"1.0.0",
		},
		{
			PackageVersion{Major: 1, Minor: 0, Patch: 0, Identifier: "dev"},
			"1.0.0-dev",
		},
		{
			PackageVersion{Major: 1, Minor: 0, Patch: 0, Identifier: "dev", Build: "13654"},
			"1.0.0-dev+13654",
		},
		{
			PackageVersion{Major: 1, Minor: 0, Patch: 0, Build: "13654"},
			"1.0.0+13654",
		},
	}

	for _, c := range testCases {
		assert.Equal(t, c.Subject.String(), c.Expected)
	}
}

func TestPackageInfo(t *testing.T) {
	subject := PackageInfo{
		"pkginfo",
		PackageVersion{
			Major: 1,
			Minor: 0,
			Patch: 0,
		},
	}
	assert.Equal(t, subject.String(), "pkginfo 1.0.0")
}
