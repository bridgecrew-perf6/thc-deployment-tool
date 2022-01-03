// Code generated by qtc from "dockerignore.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line dockerignore.qtpl:1
package templates

//line dockerignore.qtpl:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line dockerignore.qtpl:1
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line dockerignore.qtpl:1
func StreamDockerignore(qw422016 *qt422016.Writer) {
//line dockerignore.qtpl:1
	qw422016.N().S(`
Dockerfile

`)
//line dockerignore.qtpl:4
}

//line dockerignore.qtpl:4
func WriteDockerignore(qq422016 qtio422016.Writer) {
//line dockerignore.qtpl:4
	qw422016 := qt422016.AcquireWriter(qq422016)
//line dockerignore.qtpl:4
	StreamDockerignore(qw422016)
//line dockerignore.qtpl:4
	qt422016.ReleaseWriter(qw422016)
//line dockerignore.qtpl:4
}

//line dockerignore.qtpl:4
func Dockerignore() string {
//line dockerignore.qtpl:4
	qb422016 := qt422016.AcquireByteBuffer()
//line dockerignore.qtpl:4
	WriteDockerignore(qb422016)
//line dockerignore.qtpl:4
	qs422016 := string(qb422016.B)
//line dockerignore.qtpl:4
	qt422016.ReleaseByteBuffer(qb422016)
//line dockerignore.qtpl:4
	return qs422016
//line dockerignore.qtpl:4
}
