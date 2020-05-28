package helpers

import (
	"github.com/gearboxworks/buildtool/ux"
	"io/ioutil"
	"os"
	"path/filepath"
)

type TypeArgFile struct {
	Path       string
	Dir        string
	FileHandle *os.File
	FileInfo   os.FileInfo
	String     string

	State      *ux.State
}


func NewArgFile(debugMode bool) *TypeArgFile {

	p := TypeArgFile{
		Path:       "",

		Dir:        "",
		FileHandle: nil,
		FileInfo:   nil,
		String:     "",

		State:      ux.NewState(debugMode),
	}

	p.State.SetPackage("")
	p.State.SetFunctionCaller()

	return &p
}


func (at *TypeArgFile) IsNil() *ux.State {
	if state := ux.IfNilReturnError(at); state.IsError() {
		return state
	}
	at.State = at.State.EnsureNotNil()
	return at.State
}


func (at *TypeArgFile) IsOk() bool {
	var ok bool

	for range OnlyOnce {
		if at.Path == "" {
			at.State.SetError("No path specified.")
			break
		}

		if at.FileInfo == nil {
			at.State.SetError("Path doesn't exist.")
			break
		}

		ok = true
	}

	return ok
}


func (at *TypeArgFile) IsNotOk() bool {
	return !at.IsOk()
}


func (at *TypeArgFile) ChangeSuffix(suffix string) {
	s := filepath.Ext(at.Path)
	at.Path = at.Path[:len(at.Path) - len(s)] + suffix
}


func (at *TypeArgFile) SetPath(path ...string) *ux.State {
	for range OnlyOnce {
		var p string
		var err error

		p = filepath.Join(path...)
		if p == "" {
			at.State.SetError("No path specified.")
			break
		}

		at.Path, err = filepath.Abs(p)
		if err != nil {
			at.State.SetError("Path can't resolve - %s", err)
			break
		}

		at.FileInfo, err = os.Stat(at.Path)
		if os.IsNotExist(err) {
			at.FileInfo = nil
			at.State.SetError("Path does not exist - %s", err)
			break
		}

		if at.FileInfo.IsDir() {
			at.Dir = at.Path
		} else {
			at.Dir = filepath.Dir(at.Path)
		}

		at.State.SetOk()
	}

	return at.State
}


func (at *TypeArgFile) ChDir() *ux.State {
	for range OnlyOnce {
		var err error
		if at.IsNotOk() {
			break
		}

		err = os.Chdir(at.Dir)
		if err != nil {
			at.State.SetError(err)
			break
		}

		at.State.SetOk()
	}

	return at.State
}


func (p *TypeArgFile) Exists() bool {
	var ok bool

	for range OnlyOnce {
		if p.FileInfo == nil {
			p.State.SetError("path does not exist")
			break
		}
		p.State.SetOk("path exists")
		ok = true
	}

	return ok
}
func (p *TypeArgFile) NotExists() bool {
	return !p.Exists()
}

func (at *TypeArgFile) ReadFile() *ux.State {
	for range OnlyOnce {
		var err error
		if at.IsNotOk() {
			break
		}

		var d []byte
		d, err = ioutil.ReadFile(at.Path)
		if err != nil {
			at.State.SetError(err)
			break
		}
		at.String = string(d)

		at.State.SetOk()
	}

	return at.State
}
