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
node_modules
Dockerfile
.dockerignore
.git
.gitignore
docker-compose*

`)
//line dockerignore.qtpl:9
}

//line dockerignore.qtpl:9
func WriteDockerignore(qq422016 qtio422016.Writer) {
//line dockerignore.qtpl:9
	qw422016 := qt422016.AcquireWriter(qq422016)
//line dockerignore.qtpl:9
	StreamDockerignore(qw422016)
//line dockerignore.qtpl:9
	qt422016.ReleaseWriter(qw422016)
//line dockerignore.qtpl:9
}

//line dockerignore.qtpl:9
func Dockerignore() string {
//line dockerignore.qtpl:9
	qb422016 := qt422016.AcquireByteBuffer()
//line dockerignore.qtpl:9
	WriteDockerignore(qb422016)
//line dockerignore.qtpl:9
	qs422016 := string(qb422016.B)
//line dockerignore.qtpl:9
	qt422016.ReleaseByteBuffer(qb422016)
//line dockerignore.qtpl:9
	return qs422016
//line dockerignore.qtpl:9
}
