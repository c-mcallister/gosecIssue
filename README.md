# gosecIssue

this is to show the problem with gosec when using a v2 folder
see: https://github.com/securego/gosec/issues/682

```
gosec ./...
[gosec] 2021/08/12 12:36:26 Including rules: default
[gosec] 2021/08/12 12:36:26 Excluding rules: default
[gosec] 2021/08/12 12:36:26 Import directory: /Users/cam/go/src/github.com/c-mcallister/gosecIssue/v2
[gosec] 2021/08/12 12:36:26 Checking package: gosecissue
[gosec] 2021/08/12 12:36:26 Checking file: /Users/cam/go/src/github.com/c-mcallister/gosecIssue/v2/version2.go
[gosec] 2021/08/12 12:36:26 Import directory: /Users/cam/go/src/github.com/c-mcallister/gosecIssue
[gosec] 2021/08/12 12:36:27 Checking package: gosecissue
[gosec] 2021/08/12 12:36:27 Checking file: /Users/cam/go/src/github.com/c-mcallister/gosecIssue/version1.go
Results:

Golang errors in file: [/Users/cam/go/src/github.com/c-mcallister/gosecIssue/v2/version2.go]:

  > [line 6 : column 2] - could not import github.com/gin-gonic/gin (invalid package name: "")

```