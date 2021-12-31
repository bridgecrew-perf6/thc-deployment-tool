// Code generated by qtc from "commitlint.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line commitlint.qtpl:1
package templates

//line commitlint.qtpl:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line commitlint.qtpl:1
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line commitlint.qtpl:1
func StreamCommitlint(qw422016 *qt422016.Writer) {
//line commitlint.qtpl:1
	qw422016.N().S(`
# Copyright 2021 The Hacker Collective, LLC.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#    http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

name: Lint Commit Messages
on: [pull_request]

jobs:
  commitlint:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v2
        with:
          fetch-depth: 0
      - uses: wagoid/commitlint-github-action@v4

`)
//line commitlint.qtpl:28
}

//line commitlint.qtpl:28
func WriteCommitlint(qq422016 qtio422016.Writer) {
//line commitlint.qtpl:28
	qw422016 := qt422016.AcquireWriter(qq422016)
//line commitlint.qtpl:28
	StreamCommitlint(qw422016)
//line commitlint.qtpl:28
	qt422016.ReleaseWriter(qw422016)
//line commitlint.qtpl:28
}

//line commitlint.qtpl:28
func Commitlint() string {
//line commitlint.qtpl:28
	qb422016 := qt422016.AcquireByteBuffer()
//line commitlint.qtpl:28
	WriteCommitlint(qb422016)
//line commitlint.qtpl:28
	qs422016 := string(qb422016.B)
//line commitlint.qtpl:28
	qt422016.ReleaseByteBuffer(qb422016)
//line commitlint.qtpl:28
	return qs422016
//line commitlint.qtpl:28
}
