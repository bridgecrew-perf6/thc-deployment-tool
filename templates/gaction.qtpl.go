// Code generated by qtc from "gaction.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line gaction.qtpl:1
package templates

//line gaction.qtpl:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line gaction.qtpl:1
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line gaction.qtpl:1
func StreamGaction(qw422016 *qt422016.Writer, answer1 string, answer2 string, answer3 string, answer4 string, answer5 string, answerEnv string) {
//line gaction.qtpl:1
	qw422016.N().S(`
name: publish

on:
  push:
    branches: [ `)
//line gaction.qtpl:6
	qw422016.E().S(answer1)
//line gaction.qtpl:6
	qw422016.N().S(` ]

jobs:
  build:
    name: Cloud Run Deployment
    runs-on: ubuntu-latest
    steps:

      - name: Checkout
        uses: actions/checkout@master

      - name: Setup GCP Service Account
        uses: google-github-actions/setup-gcloud@master
        with:
          version: 'latest'
          service_account_email: ${{ secrets.GCP_SA_EMAIL }}
          service_account_key: ${{ secrets.GCP_SA_KEY }}
          export_default_credentials: true

      - name: Configure Docker
        run: |
          gcloud auth configure-docker
      
      - name: Build
        run: |
          docker build -t gcr.io/${{ secrets.GCP_PROJECT_ID }}/`)
//line gaction.qtpl:31
	qw422016.E().S(answer2)
//line gaction.qtpl:31
	qw422016.N().S(`:latest .
      - name: Push
        run: |
          docker push gcr.io/${{ secrets.GCP_PROJECT_ID }}/`)
//line gaction.qtpl:34
	qw422016.E().S(answer2)
//line gaction.qtpl:34
	qw422016.N().S(`:latest
      - name: Deploy
        run: |
          gcloud run deploy `)
//line gaction.qtpl:37
	qw422016.E().S(answer2)
//line gaction.qtpl:37
	qw422016.N().S(` \
          --port `)
//line gaction.qtpl:38
	qw422016.E().S(answer3)
//line gaction.qtpl:38
	qw422016.N().S(` \
          --region `)
//line gaction.qtpl:39
	qw422016.E().S(answer4)
//line gaction.qtpl:39
	qw422016.N().S(` \
          --image gcr.io/${{ secrets.GCP_PROJECT_ID }}/`)
//line gaction.qtpl:40
	qw422016.E().S(answer2)
//line gaction.qtpl:40
	qw422016.N().S(` \
          --platform managed \
          --`)
//line gaction.qtpl:42
	qw422016.E().S(answer5)
//line gaction.qtpl:42
	qw422016.N().S(` \
          --quiet \
          --project ${{ secrets.GCP_PROJECT_ID }} \
          `)
//line gaction.qtpl:45
	qw422016.E().S(answerEnv)
//line gaction.qtpl:45
	qw422016.N().S(`
          
`)
//line gaction.qtpl:47
}

//line gaction.qtpl:47
func WriteGaction(qq422016 qtio422016.Writer, answer1 string, answer2 string, answer3 string, answer4 string, answer5 string, answerEnv string) {
//line gaction.qtpl:47
	qw422016 := qt422016.AcquireWriter(qq422016)
//line gaction.qtpl:47
	StreamGaction(qw422016, answer1, answer2, answer3, answer4, answer5, answerEnv)
//line gaction.qtpl:47
	qt422016.ReleaseWriter(qw422016)
//line gaction.qtpl:47
}

//line gaction.qtpl:47
func Gaction(answer1 string, answer2 string, answer3 string, answer4 string, answer5 string, answerEnv string) string {
//line gaction.qtpl:47
	qb422016 := qt422016.AcquireByteBuffer()
//line gaction.qtpl:47
	WriteGaction(qb422016, answer1, answer2, answer3, answer4, answer5, answerEnv)
//line gaction.qtpl:47
	qs422016 := string(qb422016.B)
//line gaction.qtpl:47
	qt422016.ReleaseByteBuffer(qb422016)
//line gaction.qtpl:47
	return qs422016
//line gaction.qtpl:47
}
