package fsflags

import "os"
import "fmt"

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


// flag value for file, creates if needed, overwrite without error.
type CreateFileValue struct{
    FileValue
}

func (fsf *CreateFileValue) Set(v string) (err error) {
	fsf.File,err=os.Create(v)
	return
}

// flag value for file, appends, creates if needed.
type AppendFileValue struct{
    FileValue
}

func (fsf *CreateFileValue) Set(v string) (err error) {
	fsf.File, err := os.OpenFile(v, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	return
}

// flag value for a file, creates in given dir (also created if needed) new each day, appended too if pre-exists.
// to maintain strict daily log will need re-making after midnight.
// used for logging dont need date, so could use flag:LTime in the log standard package
type DailyFileValue struct{
    FileValue
}

func (fsf *DailyFileValue) Set(v string) (err error) {
	f,err:=os.Open(v)
	if os.IsNotExist(err){
		err=os.Mkdir(v,0777)
 		if err!=nil{ return}
		f,err=os.Open(v)
 		if err!=nil{ return}
	}else{
		fi,err:=f.Stat()
		if err!=nil{ return}
		if !fi.IsDir(){return os.ErrNotExist}
	}
	y,m,d:=time.Now().Date()
	fsf.File, err := os.OpenFile(fmt.Sprintf("%s/%4d-%2d-%2d",f.Name(),y,m,d), os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	return
}


