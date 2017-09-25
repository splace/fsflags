package fsflags

import "os"

// flag value for existing file
type fileValue struct{
    *os.File
}

func (fsf *fileValue) Set(v string) (err error) {
    fsf.File,err=os.Open(v)
    return
}

func (fsf *fileValue) String() string {
    if fsf==nil || fsf.File==nil {return "<nil>"}
    return fsf.File.Name()
}


// flag value for file, creates if needed.
type createFileValue struct{
    fileValue
}

func (fsf *createFileValue) Set(v string) (err error) {
	fsf.File,err=os.Create(v)
	return
}

