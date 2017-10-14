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
	"net/http"
	"strings"
)

// Opener is a FileSystem.Files value used for creating new File
// instances.
type Opener struct {
	FileInfo
	Content string
}

// Open creates a new, open File instance.
func (s *Opener) Open(name string) (f http.File, err error) {
	info := s.FileInfo
	info.name = name
	info.size = int64(len(s.Content))
	return &File{
		ReadSeeker: strings.NewReader(s.Content),
		info:       info,
	}, nil
}
