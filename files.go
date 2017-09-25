package fsflags

import "os"

// flag value for existing file
type FileValue struct{
    *os.File
}

func (fsf *FileValue) Set(v string) (err error) {
    fsf.File,err=os.Open(v)
    return
}

func (fsf *FileValue) String() string {
    if fsf==nil || fsf.File==nil {return "<nil>"}
    return fsf.File.Name()
}


// flag value for file, creates if needed.
type CreateFileValue struct{
    FileValue
}

func (fsf *CreateFileValue) Set(v string) (err error) {
	fsf.File,err=os.Create(v)
	return
}

