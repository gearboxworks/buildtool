package helpers

import (
	"github.com/newclarity/buildtool/ux"
	"strings"
)

const (
	BinaryName = "BinaryName"
	BinaryVersion = "BinaryVersion"
	BinaryRepo = "BinaryRepo"
	SourceRepo = "SourceRepo"
)

var DefaultVersionFile = []string{"defaults", "version.go"}


func (at *ArgTemplate) GetValue(lookfor string, path ...string) (string, *ux.State) {
	var version string

	for range OnlyOnce {
		grFile := NewArgFile(at.Debug)
		at.State = grFile.SetPath(path...)
		if grFile.NotExists() {
			at.State = grFile.State
			break
		}

		at.State = grFile.ReadFile()
		if at.State.IsNotOk() {
			break
		}

		for _, v := range strings.Split(grFile.String, "\n") {
			if !strings.Contains(v, lookfor) {
				continue
			}

			sa := strings.Split(v, "=")
			if len(sa) != 2 {
				continue
			}

			version = strings.TrimSpace(sa[1])
			version = strings.TrimPrefix(version, "\"")
			version = strings.TrimSuffix(version, "\"")
			break
		}

		//
	}

	return version, at.State
}


func (at *ArgTemplate) GetBinaryVersion(path ...string) (string, *ux.State) {
	return at.GetValue(BinaryVersion, path...)
}


func (at *ArgTemplate) GetBinaryName(path ...string) (string, *ux.State) {
	return at.GetValue(BinaryName, path...)
}


func (at *ArgTemplate) GetBinaryRepo(path ...string) (string, *ux.State) {
	return at.GetValue(BinaryRepo, path...)
}


func (at *ArgTemplate) GetSourceRepo(path ...string) (string, *ux.State) {
	return at.GetValue(SourceRepo, path...)
}


func (at *ArgTemplate) GetSourceOwner(path ...string) (string, *ux.State) {
	var owner string

	for range OnlyOnce {
		owner, at.State = at.GetValue(SourceRepo, path...)
		owner, _ = GetRepoComponents(owner)
		at.State.SetOk()
	}

	return owner, at.State
}


func (at *ArgTemplate) GetSourceRepoName(path ...string) (string, *ux.State) {
	var name string

	for range OnlyOnce {
		name, at.State = at.GetValue(SourceRepo, path...)
		name, _ = GetRepoComponents(name)
		at.State.SetOk()
	}

	return name, at.State
}


func (at *ArgTemplate) GetBinaryOwner(path ...string) (string, *ux.State) {
	var owner string

	for range OnlyOnce {
		owner, at.State = at.GetValue(BinaryRepo, path...)
		owner, _ = GetRepoComponents(owner)
		at.State.SetOk()
	}

	return owner, at.State
}


func (at *ArgTemplate) GetBinaryRepoName(path ...string) (string, *ux.State) {
	var name string

	for range OnlyOnce {
		name, at.State = at.GetValue(BinaryRepo, path...)
		name, _ = GetRepoComponents(name)
		at.State.SetOk()
	}

	return name, at.State
}


func GetRepoComponents(url string) (string, string) {
	var owner string
	var name string

	for range OnlyOnce {
		url = StripUrlPrefix(url)
		ua := strings.Split(url, "/")
		switch len(ua) {
			case 0:
				break
			case 1:
				owner = ua[0]
			case 2:
				owner = ua[0]
				name = ua[1]
			default:
				owner = ua[0]
				name = ua[1]
		}
	}

	return owner, name
}


func StripUrlPrefix(url string) string {
	url = strings.TrimPrefix(url, "https://")
	url = strings.TrimPrefix(url, "github.com/")
	url = strings.TrimSuffix(url, "/")
	url = strings.TrimSpace(url)

	return url
}
