# fsflags


Overview/docs: [![GoDoc](https://godoc.org/github.com/splace/fsplags?status.svg)](https://godoc.org/github.com/splace/fsflags) 

Installation:

     go get github.com/splace/fsflags

Example:
```go
package main

import "os"
import "flag"
import "github.com/splace/fsflags"
import "log"
import "path/filepath"

func main(){
	var logFolder fsflags.newDirValue   // create or reuse folder, emptying it on startup.
	flag.Var(&logFolder, "f", "folder for log files.")
	flag.Parse()

	if logFolder.File == nil {
		progressLog=log.New(os.Stdout, "", log.LstdFlags)
		errorLog=log.New(os.Stderr, "", log.LstdFlags)
	}else{
		progressLogFile,err:=os.Create(filepath.Join(logFolder,"progress.log"))
		if err!=nil{log.Printf(err)}
		progressLog=log.New(progressLogFile, "", log.LstdFlags)
		errorLogFile,err:=os.OpenFile(filepath.Join(logFolder,"errors.log"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err!=nil{log.Printf(err)}
 		errorLog=log.New(errorLogFile, "", log.LstdFlags)
	}
	
}
```