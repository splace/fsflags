package fsflags

import "os"

// flag value for existing file, error if it doesn't
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


// flag value for file, creates if needed, overwrite without error..
type CreateFileValue struct{
    FileValue
}

func (fsf *CreateFileValue) Set(v string) (err error) {
	fsf.File,err=os.Create(v)
	return
}

// flag value for file, creates, error if exists.
type NewFileValue struct{
    FileValue
}

func (fsf *NewFileValue) Set(v string) (err error) {
    fsf.File,err=os.Open(v)
    if !os.IsNotExist(err){
		return os.ErrExist
	}    
	fsf.File,err=os.Create(v)    
	return
}


