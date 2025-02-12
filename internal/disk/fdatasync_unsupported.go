//go:build !linux && !netbsd && !freebsd && !darwin && !openbsd
// +build !linux,!netbsd,!freebsd,!darwin,!openbsd

// Copyright (c) 2015-2021 MinIO, Inc.
//
// This file is part of MinIO Object Storage stack
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package disk

import (
	"os"
)

// Fdatasync is a no-op
func Fdatasync(f *os.File) error {
	return nil
}

// fdavise advice constants
const (
	FadvSequential = 0
	FadvNoReuse    = 0
)

// Fadvise implements possibility of choosing
// offset: 0, length: 0
func Fadvise(f *os.File, advice int) error {
	return nil
}
