package fsflags

import "os"

// flag value for an existing directory.
type DirValue struct{
    fileValue
}

func (fsf *DirValue) Set(v string) error {
	f,err:=os.Open(v)
    if err!=nil{ return err}
    fi,err:=f.Stat()
    if err!=nil{ return err}
	if !fi.IsDir(){return os.ErrNotExist}
    fsf.File=f
    return err
}

// flag value for an existing directory, creates if needed.
type NewDirValue struct{
    fileValue
}

func (fsf *NewDirValue) Set(v string) (err error) {
	f,err:=os.Open(v)
	if os.IsNotExist(err){
		err=os.Mkdir(v,0777)
 		if err!=nil{ return}
		f,err=os.Open(v)
 		if err!=nil{ return}
 		fsf.File=f
		return
	}
    fi,err:=f.Stat()
    if err!=nil{ return}
	if !fi.IsDir(){return os.ErrNotExist}
    fsf.File=f
    return
}

// flag value for a directory, creates if needed, any pre-existing hierarchy inside it is erased.
type NewOverwriteDirValue struct{
    fileValue
}

func (fsf *NewOverwriteDirValue) Set(v string) (err error) {
	f,err:=os.Open(v)
	if os.IsNotExist(err){
		err=os.Mkdir(v,0777)
 		if err!=nil{ return}
		f,err=os.Open(v)
 		if err!=nil{ return}
 		fsf.File=f
		return
	}
    fi,err:=f.Stat()
    if err!=nil{ return}
	if !fi.IsDir(){return os.ErrNotExist}
	err=removeContents(f)
    if err!=nil{ return}
    fsf.File=f
    return
}

// flag value for a directory, pre-existing, emptied.
type OverwriteDirValue struct{
    fileValue
}

func (fsf *OverwriteDirValue) Set(v string) (err error) {
	f,err:=os.Open(v)
    if err!=nil{ return}
    fi,err:=f.Stat()
    if err!=nil{ return}
	if !fi.IsDir(){return os.ErrNotExist}
	err=removeContents(f)
    if err!=nil{ return}
    fsf.File=f
    return
}

// flag value for a directory, creates if needed, any pre-existing files inside it are erased, (but not directories).
type NewOverwriteFilesDirValue struct{
    fileValue
}

func (fsf *NewOverwriteFilesDirValue) Set(v string) (err error) {
	f,err:=os.Open(v)
	if os.IsNotExist(err){
		err=os.Mkdir(v,0777)
 		if err!=nil{ return}
		f,err=os.Open(v)
 		if err!=nil{ return}
 		fsf.File=f
		return
	}
    fi,err:=f.Stat()
    if err!=nil{ return}
	if !fi.IsDir(){return os.ErrNotExist}
	err=removeFileContents(f)
    if err!=nil{ return}
    fsf.File=f
    return
}

// flag value for a directory, pre-existing, any pre-existing files inside it are erased, (but not directories).
type OverwriteFilesDirValue struct{
    fileValue
}

func (fsf *OverwriteFilesDirValue) Set(v string) (err error) {
	f,err:=os.Open(v)
    if err!=nil{ return}
    fi,err:=f.Stat()
    if err!=nil{ return}
	if !fi.IsDir(){return os.ErrNotExist}
	err=removeFileContents(f)
    if err!=nil{ return}
    fsf.File=f
    return
}


// flag value for a directory, creates if needed, any pre-existing directories inside it are erased, (but not files).
type NewOverwriteSubdirsDirValue struct{
    fileValue
}

func (fsf *NewOverwriteSubdirsDirValue) Set(v string) (err error) {
	f,err:=os.Open(v)
	if os.IsNotExist(err){
		err=os.Mkdir(v,0777)
 		if err!=nil{ return}
		f,err=os.Open(v)
 		if err!=nil{ return}
 		fsf.File=f
		return
	}
    fi,err:=f.Stat()
    if err!=nil{ return}
	if !fi.IsDir(){return os.ErrNotExist}
	err=removeDirContents(f)
    if err!=nil{ return}
    fsf.File=f
    return
}

// flag value for a directory, pre-existing, any pre-existing directories inside it are erased, (but not files).
type OverwriteSubdirsDirValue struct{
    fileValue
}

func (fsf *OverwriteSubdirsDirValue) Set(v string) (err error) {
	f,err:=os.Open(v)
    if err!=nil{ return}
    fi,err:=f.Stat()
    if err!=nil{ return}
	if !fi.IsDir(){return os.ErrNotExist}
	err=removeDirContents(f)
    if err!=nil{ return}
    fsf.File=f
    return
}

// flag value for a directory, not pre-existing.
type MakeDirValue struct{
    fileValue
}

func (fsf *MakeDirValue) Set(v string) (err error) {
	err=os.Mkdir(v,0777)
    if err!=nil{ return}
    fsf.File,err=os.Open(v)
    return
}

// flag value for a directory, not pre-existing, possibly multiple levels down.
type MakeDirAllValue struct{
    fileValue
}

func (fsf *MakeDirAllValue) Set(v string) (err error) {
	err=os.MkdirAll(v,0777)
    if err!=nil{ return}
    fsf.File,err=os.Open(v)
   return
}

// flag value for a directory, possibly down multiple levels. if pre-existing erased.
type MakeDirOverwriteAllValue struct{
    fileValue
}

func (fsf *MakeDirOverwriteAllValue) Set(v string) (err error) {
	err=os.RemoveAll(v)
    if err!=nil{ return}
	err=os.MkdirAll(v,0777)
    fsf.File,err=os.Open(v)
   return
}

// flag value for a new directory at this level. if pre-existing erased.
type MakeDirOverwriteValue struct{
    fileValue
}

func (fsf *MakeDirOverwriteValue) Set(v string) (err error) {
	err=os.RemoveAll(v)
    if err!=nil{ return}
	err=os.Mkdir(v,0777)
    fsf.File,err=os.Open(v)
   return
}

func removeContents(d *os.File) error {
	finfos, err := d.Readdir(-1)
    if err != nil {
        return err
    }
	defer  changeWorkingDirReset(d)()
    for _, finfo := range finfos {
    	if finfo.IsDir(){
			err=os.RemoveAll(finfo.Name())	
    	}else{	
			err=os.Remove(finfo.Name())	
		}
	    if err != nil {
	        return err
	    }
    }
    return nil
}

func removeFileContents(d *os.File) error {
	finfos, err := d.Readdir(-1)
    if err != nil {
        return err
    }
	defer  changeWorkingDirReset(d)()
    for _, finfo := range finfos {
    	if !finfo.IsDir(){
			err=os.Remove(finfo.Name())	
		}
	    if err != nil {
	        return err
	    }
    }
    return nil
}

func removeDirContents(d *os.File) error {
	finfos, err := d.Readdir(-1)
    if err != nil {
        return err
    }
	defer  changeWorkingDirReset(d)()
    for _, finfo := range finfos {
    	if finfo.IsDir(){
			err=os.RemoveAll(finfo.Name())	
		}
	    if err != nil {
	        return err
	    }
    }
    return nil
}


func changeWorkingDirReset(dir *os.File) (fn func()) {
	currentDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	err = dir.Chdir()
	if err == nil {
		return func() { os.Chdir(currentDir) }
	}
	return
}
