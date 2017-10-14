// Copyright 2017 fakefs contributors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package fakefs

import (
	"io"
	"os"
)

// File implements net/https's File in memory.
type File struct {
	io.ReadSeeker
	info   FileInfo
	closed bool
}

// Close closes the File, rendering it unusable for I/O.
func (f *File) Close() (err error) {
	if f.closed {
		// os.ErrClosed was added in Go go1.8
		return io.ErrClosedPipe
	}
	f.ReadSeeker = &ErrorReadSeeker{
		error: io.ErrClosedPipe,
	}
	f.closed = true
	return nil
}

// Readdir reads the contents of the directory associated with file
// and returns a slice of up to n FileInfo values, as would be
// returned by Lstat, in directory order. Subsequent calls on the same
// file will yield further FileInfos.
//
// Readdir is not currently implemented, and will always panic.
func (f *File) Readdir(count int) (info []os.FileInfo, err error) {
	panic("File.Readdir() is not implemented")
}

// Stat returns the FileInfo structure describing file.
func (f *File) Stat() (info os.FileInfo, err error) {
	i := f.info
	return &i, nil
}
