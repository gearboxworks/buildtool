package cmd

import (
	"github.com/newclarity/scribeHelpers/toolPath"
	"github.com/newclarity/scribeHelpers/ux"
	"strings"
)

const (
	BinaryName = "BinaryName"
	BinaryVersion = "BinaryVersion"
	BinaryRepo = "BinaryRepo"
	SourceRepo = "SourceRepo"
)

var DefaultVersionFile = []string{"defaults", "version.go"}


func Get(args ...string) *ux.State {
	state := Cmd.State

	for range onlyOnce {
		if len(args) == 0 {
			state = GetHelp()
			break
		}

		if args[0] == "all" {
			args = []string{"name", "version", "src", "bin"}
		}

		for _, v := range args {
			k := ""
			v = strings.ToLower(v)
			switch {
			case v == "name":
				k = "Name"
				v = BinaryName

			case v == "ver":
				fallthrough
			case v == "version":
				k = "Version"
				v = BinaryVersion

			case v == "src":
				fallthrough
			case v == "source":
				fallthrough
			case v == "srcrepo":
				k = "Source Repo"
				v = SourceRepo

			case v == "bin":
				fallthrough
			case v == "binary":
				fallthrough
			case v == "binrepo":
				k = "Binary Repo"
				v = BinaryRepo
			}

			v, state = getValue(v, DefaultVersionFile...)
			if state.IsNotOk() {
				continue
			}

			if len(args) > 1 {
				ux.PrintfBlue("%s: ", k)
			}
			ux.PrintflnCyan("%s", v)
		}
	}

	return state
}


func GetHelp() *ux.State {
	ux.PrintflnYellow("Need to supply one of:")
	ux.PrintflnYellow("\t'all' - Show all of the below.")
	ux.PrintflnYellow("\t'name' - Show the binary name.")
	ux.PrintflnYellow("\t'version' - Show the binary version.")
	ux.PrintflnYellow("\t'src' - Show the Git source repo.")
	ux.PrintflnYellow("\t'bin' - Show the Git binary repo.")
	return Cmd.State
}


func getValue(lookfor string, path ...string) (string, *ux.State) {
	var version string
	state := Cmd.State

	for range onlyOnce {
		if len(path) == 0 {
			path = []string{"."}
		}

		grFile := toolPath.New(Cmd.Runtime)
		grFile.SetPath(path...)

		state = grFile.StatPath()
		if state.IsNotOk() {
			break
		}

		state = grFile.ReadFile()
		if state.IsNotOk() {
			break
		}

		for _, v := range grFile.GetContentArray() {
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
	}

	return version, state
}


func getBinaryVersion(path ...string) (string, *ux.State) {
	if len(path) == 0 {
		path = DefaultVersionFile
	}
	ret, state := getValue(BinaryVersion, path...)
	if !strings.HasPrefix(ret, "v") {
		ret = "v" + ret
	}
	return ret, state
}


func getBinaryName(path ...string) (string, *ux.State) {
	if len(path) == 0 {
		path = DefaultVersionFile
	}
	return getValue(BinaryName, path...)
}


func getBinaryRepo(path ...string) (string, *ux.State) {
	if len(path) == 0 {
		path = DefaultVersionFile
	}
	return getValue(BinaryRepo, path...)
}


func getSourceRepo(path ...string) (string, *ux.State) {
	if len(path) == 0 {
		path = DefaultVersionFile
	}
	return getValue(SourceRepo, path...)
}


func getSourceOwner(path ...string) (string, *ux.State) {
	var owner string
	state := Cmd.State

	for range onlyOnce {
		owner, state = getValue(SourceRepo, path...)
		owner, _ = GetRepoComponents(owner)
		state.SetOk()
	}

	return owner, state
}


func getSourceRepoName(path ...string) (string, *ux.State) {
	var name string
	state := Cmd.State

	for range onlyOnce {
		name, state = getValue(SourceRepo, path...)
		name, _ = GetRepoComponents(name)
		state.SetOk()
	}

	return name, state
}


func getBinaryOwner(path ...string) (string, *ux.State) {
	var owner string
	state := Cmd.State

	for range onlyOnce {
		owner, state = getValue(BinaryRepo, path...)
		owner, _ = GetRepoComponents(owner)
		state.SetOk()
	}

	return owner, state
}


func getBinaryRepoName(path ...string) (string, *ux.State) {
	var name string
	state := Cmd.State

	for range onlyOnce {
		name, state = getValue(BinaryRepo, path...)
		name, _ = GetRepoComponents(name)
		state.SetOk()
	}

	return name, state
}


func GetRepoComponents(url string) (string, string) {
	var owner string
	var name string

	for range onlyOnce {
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
