package fsflags

import "os"
import "fmt"
import "time"
import "io/ioutil"

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

func (fsf *AppendFileValue) Set(v string) (err error) {
	fsf.File, err = os.OpenFile(v, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	return
}

// flag value for a file, creates in given dir (also created if needed) new each day, appended too if pre-exists.
// same file for each invocation, to maintain strict daily log will need re-making after midnight.
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
	}
	fi,err:=f.Stat()
	if err!=nil{ return}
	if !fi.IsDir(){return os.ErrNotExist}

	y,m,d:=time.Now().Date()
	fsf.File, err = os.OpenFile(fmt.Sprintf("%s/%4d-%02d-%02d",f.Name(),y,m,d), os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
	return
}

// flag value for a file, creates in given dir (also created if needed) new each day, appended too if pre-exists.
// same file for each invocation, to maintain strict daily log will need re-making after midnight.
// erases up to two files (oldest)) to maintain at least required number of files.
// used for logging dont need date, so could use flag:LTime in the log standard package
type DailyErasingFileValue struct{
    DailyFileValue
    Required int 
}

func (fsf *DailyErasingFileValue) Set(v string) (err error) {
	err=fsf.DailyFileValue.Set(v)
	if err==nil{
	     files, derr := ioutil.ReadDir(v)
     	if derr!=nil || len(files)<fsf.Required+2 { return}
    	  oldestTime := time.Now()
    	  var oldestFile,secondoldestFile os.FileInfo
		for _, file := range files {
              if file.Mode().IsRegular() && file.ModTime().Before(oldestTime) {
                      secondoldestFile = oldestFile
                      oldestFile = file
                      oldestTime = file.ModTime()
              }
         if secondoldestFile!=nil{os.Remove(v+"/"+secondoldestFile.Name())}
         if oldestFile!=nil{os.Remove(v+"/"+oldestFile.Name())}
      }

	}
	return
}


