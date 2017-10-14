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

// Package fakefs is an in-memory net/http FileSystem for mock
// testing.
//
// Status: DEPRECATED
//
// I'll keep this repository around because it's pretty tiny, but you
// can accomplish its goals using mapfs and httpfs:
//
// * https://godoc.org/golang.org/x/tools/godoc/vfs/httpfs
// * https://godoc.org/golang.org/x/tools/godoc/vfs/mapfs
//
// For example:
//
//   fakeFS := httpfs.New(mapfs.New(map[string]string{
//     "etc/passwd": "root:x:0:0:root:/root:/bin/bash\n"
//   }
//   transport := &http.Transport{}
//   transport.RegisterProtocol("file", http.NewFileTransport(fakeFS))
//
// And here's the fakefs version:
//
//   fakeFS := &fakefs.FileSystem{
//     Files: map[string]*fakefs.Opener{
//       "/etc/passwd": &fakefs.Opener{
//         FileInfo: fakefs.FileInfo{
//           FileMode: 0x644,
//         },
//         Content: "root:x:0:0:root:/root:/bin/bash\n"
//       },
//     },
//   }
//   transport := &http.Transport{}
//   transport.RegisterProtocol("file", http.NewFileTransport(fakeFS))
//   client := &http.Client{Transport: transport}
//   response, err := client.Get("file:///etc/passwd")
package fakefs

import (
	"net/http"
	"os"
	"path/filepath"
)

// FileSystem implements net/http's FileSystem in memory.
type FileSystem struct {
	Files map[string]*Opener
}

// Open opens the named file for reading.
func (fs *FileSystem) Open(name string) (f http.File, err error) {
	spawner, ok := fs.Files[name]
	if !ok {
		return nil, os.ErrNotExist
	}

	base := filepath.Base(name)
	return spawner.Open(base)
}
