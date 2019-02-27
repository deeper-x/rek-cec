# rek-cec [WIP]
http service availability checker

_Problem:_
Check http services availability, status codes, result.

_Solution:_
Given a list of http services to check, return a failure/success status list.     

# Usage #
Example: check if services are available [response: 200 OK]
```bash
$ rek-cec [HTTP CODE - Default: 200]
```
Explanation: given an input status code (optional), http services listed in "to_check.txt" are called, and responses are checked.

# Configuration, build, install #

System/env stuff settings:
```bash
$ export GOPATH="/home/${USER}/go/
$ export PATH=${PATH}:${GOPATH}/bin
```
Check it out:
```bash
$ go version
go version go1.10.4 linux/amd64
$ go env GOPATH
/home/<YOURUSER>/go
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
$ go build -o bin/rek-check # Pass your preferred bin path
$ ./bin/rek-cec 
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

# Service list configuration #
List http services to scan:

```bash
$ cat /home/${USER}/go/src/github.com/deeper-x/rek-cec/assets/to_check.txt
http://google.it
http://www.bbc.com
http://youtube.com
```