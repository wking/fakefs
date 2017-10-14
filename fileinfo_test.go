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
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	opener := &Opener{
		FileInfo: FileInfo{
			FileMode: 0x644,
		},
	}
	file, err := opener.Open("a")
	if err != nil {
		t.Fatal(err)
	}

	info, err := file.Stat()
	if err != nil {
		t.Fatal(err)
	}

	t.Run("mode", func(t *testing.T) {
		assert.Equal(t, opener.FileMode, info.Mode())
	})

	t.Run("sys", func(t *testing.T) {
		assert.Equal(t, nil, info.Sys())
	})
}
