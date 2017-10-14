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
	"io/ioutil"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFileSystem(t *testing.T) {
	fakeFS := &FileSystem{
		Files: map[string]*Opener{
			"/etc/passwd": &Opener{
				FileInfo: FileInfo{
					FileMode: 0x644,
				},
				Content: "root:x:0:0:root:/root:/bin/bash\nbin:x:1:1:bin:/bin:/bin/false\n",
			},
		},
	}

	t.Run("transport protocol", func(t *testing.T) {
		transport := &http.Transport{}
		transport.RegisterProtocol("file", http.NewFileTransport(fakeFS))
		client := &http.Client{Transport: transport}

		response, err := client.Get("file:///etc/passwd")
		if err != nil {
			t.Fatal(err)
		}

		content, err := ioutil.ReadAll(response.Body)
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, fakeFS.Files["/etc/passwd"].Content, string(content))
	})

	t.Run("does not exist", func(t *testing.T) {
		file, err := fakeFS.Open("/does/not/exist")
		assert.Equal(t, err, os.ErrNotExist)
		assert.Equal(t, nil, file)
	})
}
