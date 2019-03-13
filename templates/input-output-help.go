package main

import "github.com/splace/fsflags"
import "flag"

import "log"
import "os"
import	"io/ioutil"
import	"time"


func main() {
	var source fsflags.FileValue
	flag.Var(&source, "i", "input file/device.(default:<Stdin>)")
	flag.Var(&source, "input", flag.Lookup("i").Usage)
	var sink fsflags.NewFileValue
	flag.Var(&sink, "o", "output file/device.(default:Stdout)")
	flag.Var(&sink, "output", flag.Lookup("o").Usage)
	var over fsflags.CreateFileValue
	flag.Var(&over, "oo", "output file/device.(silent overwrite, ignored if output set).")

	var help bool
	flag.BoolVar(&help, "h", false, "display help/usage.")
	flag.BoolVar(&help, "help", false, flag.Lookup("h").Usage)
	flag.Parse()

	var logToo fsflags.CreateFileValue
	flag.Var(&logToo, "log", "log destination.(default:Stderr)")
	var logInterval time.Duration
	flag.DurationVar(&logInterval, "interval", time.Second, "time between log status reports.")
	var limit time.Duration
	flag.DurationVar(&limit, "end", 0, "time limit.")
	var quiet bool
	flag.BoolVar(&quiet, "quiet", false, "no progress logging.")
	flag.BoolVar(&quiet, "q", false, "no progress logging.")

	if help {
		flag.PrintDefaults()
		os.Exit(0)
	}

	if sink.File == nil {
		if over.File == nil {
			sink.File = os.Stdout
		}else{
			sink.File = over.File
		}
	}
	
	if source.File == nil {
		source.File = os.Stdin
	}

	
	if logToo.File == nil {
		logToo.File = os.Stderr
	}
	
	var progressLog *log.Logger
	if quiet {
		progressLog = log.New(ioutil.Discard, "", log.LstdFlags)
	} else {
		progressLog = log.New(logToo, "", log.LstdFlags)
	}
	
	progressLog.Printf("Loading:%q", &source)

	doLog := time.NewTicker(logInterval)

	startTime := time.Now()
	go func() {
		for {
			select {
			case t := <-doLog.C:
				if limit > 0 && t.Sub(startTime) > limit {
					progressLog.Printf("Timed out:%q", &source)
					os.Exit(124)
				}
				progressLog.Printf("\t@%v",t.Sub(startTime)/time.Second*time.Second)
			}
		}
	}()
	
	// code here

//	var format string
//	// if available use extension of sink file for format
//	if ext:=filepath.Ext(sink.Name());ext!=""{
//		format=ext[1:]
//	}

//		if ext:=filepath.Ext(sink.Name());ext!="" && ext!=".png"{
//			os.Rename(sink.Name(),sink.Name()+".png")
//		}


	
	progressLog.Printf("Saving:%q", &sink)

}

