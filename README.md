## rek-cec 
http service availability checker

## Usage ##
Example: check if services are available (= response is 200 OK)
```bash
$ rek-cec [HTTP CODE - Default: 200]
```
Explanation: given an input status code (optional), http services listed in "to_check.txt" are called, and responses are checked.

## Configuration, build, install ##

System/env stuff settings:
```bash
$ export GOPATH=${HOME}/go
$ export GOBIN=${GOPATH}/bin
$ export PATH=${PATH}:${GOBIN}
```

Check it out:
```bash
$ go version
go version go1.13.3 linux/amd64

$ go env 
/home/<YOURUSER>/go
...
$ echo $PATH
/a/long/sequence/:/of/paths/:/ending/with:/home/<YOURUSER>/go/bin
```

Download package:
```bash
$ go get github.com/deeper-x/rek-cec
```

Look at the downloaded package:
```bash
$ cd ${GOPATH}/src/github.com/deeper-x/rek-cec
$ ls
[...root content...]
```

Build & run:
```bash
$ go build 
$ ./rek-cec 
http://google.com [200]: WORKS!
http://www.bbc.com [200]: WORKS!
http://youtube.com [200]: WORKS!
```

Or install & run:
```bash
$ go install 
$ rek-cec # Assuming your env path stuff is well configured
http://google.it [200]: WORKS!
http://www.bbc.com [200]: WORKS!
http://youtube.com [200]: WORKS!
```

## Unit test ##
```bash
$ go test -v -cover ./...
?   	github.com/deeper-x/rek-cec	[no test files]
=== RUN   TestCallService
--- PASS: TestCallService (2.25s)
PASS
coverage: 85.7% of statements
ok  	github.com/deeper-x/rek-cec/checker	2.255s	coverage: 85.7% of statements
=== RUN   TestWriteResponseOutput
--- PASS: TestWriteResponseOutput (0.00s)
=== RUN   TestLoadConfigurationFile
--- PASS: TestLoadConfigurationFile (0.00s)
=== RUN   TestGetExpectedStatus
--- PASS: TestGetExpectedStatus (0.00s)
PASS
coverage: 72.0% of statements
ok  	github.com/deeper-x/rek-cec/utils	0.002s	coverage: 72.0% of statements
```


## Service list configuration ##
List http services to scan:

```bash
$ cat /home/${USER}/go/src/github.com/deeper-x/rek-cec/assets/to_check.txt
http://google.it
http://www.bbc.com
http://youtube.com
```
