# fsflags


Overview/docs: [![GoDoc](https://godoc.org/github.com/splace/fsplags?status.svg)](https://godoc.org/github.com/splace/fsflags) 

Installation:

     go get github.com/splace/fsflags

Example: get input, output and daily (self-maintaining) log folder from parameters. use stdin and stdout by default.
```go
package main

import "os"
import "flag"
import "github.com/splace/fsflags"
import "log"

func main() {
	var read fsflags.FileValue
	flag.Var(&read, "i", "input file.")
	var write fsflags.NewFileValue
	flag.Var(&write, "o", "output file (will not overwrite).")
	var over fsflags.CreateFileValue
	flag.Var(&over, "oo", "output file (will overwrite).")
	var daily fsflags.DailyErasingFileValue
	flag.Var(&daily, "d", "folder for log writes, appending to new file (named "YYYY-MM-DD") each day (self erasing).")
	var verbose bool
	flag.BoolVar(&verbose, "v", false, "display log writes.")
	flag.Parse()

	if read.File == nil {
		read.File = os.Stdin
	}

	if write.File == nil {
		if over.File == nil {
			write.File = os.Stdout
		}else{
			write.File = over.File
		}
	}

	// make logger that writes as flags indicate.
	var logger *log.Logger
	if verbose && daily.File!=nil {
		logger=log.New(io.MultiWriter(os.Stdout,daily.File),"",log.Ltime)
	} else {
		if daily.File!=nil{
			logger=log.New(daily.File,"",log.Ltime)
		} else {
			if verbose{
				logger=log.New(os.Stdout,"",log.Ltime)
			}
		}
	}

}
```
