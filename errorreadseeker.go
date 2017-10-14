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

// ErrorReadSeeker returns an error for all message calls.  This is useful for representing a closed ReadSeeker.
type ErrorReadSeeker struct {
	error
}

// Read always returns zero and an error.
func (ers *ErrorReadSeeker) Read(p []byte) (n int, err error) {
	return 0, ers.error
}

// Seek always returns zero and an error.
func (ers *ErrorReadSeeker) Seek(offset int64, whence int) (newOffset int64, err error) {
	return 0, ers.error
}
