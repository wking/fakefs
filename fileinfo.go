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
	"os"
	"time"
)

// FileInfo implements os's FileInfo in memory.
type FileInfo struct {
	os.FileMode
	name string
	size int64
}

// IsDir reports whether this FileInfo describes a directory.
func (fi *FileInfo) IsDir() bool {
	return fi.FileMode&os.ModeDir != 0
}

// Mode returns the file's mode bits.
func (fi *FileInfo) Mode() os.FileMode {
	return fi.FileMode
}

// ModTime returns the file's modification time.
func (fi *FileInfo) ModTime() time.Time {
	return time.Unix(0, 0)
}

// Name returns the base name of the file as presented to Open.
func (fi *FileInfo) Name() string {
	return fi.name
}

// Size returns the length in bytes for regular files.  The value is
// system-dependent for other files.
func (fi *FileInfo) Size() int64 {
	return fi.size
}

// Sys always returns nil.
func (fi *FileInfo) Sys() interface{} {
	return nil
}
