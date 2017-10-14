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
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDuplicateClose(t *testing.T) {
	opener := &Opener{}
	file, err := opener.Open("a")
	if err != nil {
		t.Fatal(err)
	}

	err = file.Close()
	if err != nil {
		t.Fatal(err)
	}

	err = file.Close()
	assert.Equal(t, io.ErrClosedPipe, err)
}

func TestReadAfterClose(t *testing.T) {
	opener := &Opener{}
	file, err := opener.Open("a")
	if err != nil {
		t.Fatal(err)
	}

	err = file.Close()
	if err != nil {
		t.Fatal(err)
	}

	buf := []byte("buffer for read bytes")
	_, err = file.Read(buf)
	assert.Equal(t, io.ErrClosedPipe, err)
}

func TestSeekAfterClose(t *testing.T) {
	opener := &Opener{}
	file, err := opener.Open("a")
	if err != nil {
		t.Fatal(err)
	}

	err = file.Close()
	if err != nil {
		t.Fatal(err)
	}

	_, err = file.Seek(0, io.SeekStart)
	assert.Equal(t, io.ErrClosedPipe, err)
}
