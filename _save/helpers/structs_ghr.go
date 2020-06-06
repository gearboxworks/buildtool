package helpers


//const githubRelease = "github-release"
//
//type TypeGhr struct {
//	Args  []string
//
//	owner string
//	repo  string
//
//	debug bool
//	State *ux.State
//}
//
//
//func NewGitHubRelease(debugMode bool) *TypeGhr {
//	var ghr TypeGhr
//	ghr.State = ux.NewState("foo", true)
//
//	for range OnlyOnce {
//		ghr.State.SetPackage("")
//		ghr.State.SetFunctionCaller()
//
//		// A quick hack. To bring in github-release into this code.
//		// Later it'll be merged properly.
//		ghr.Args = []string{githubRelease}
//	}
//
//	return &ghr
//}
//
//
//func (ghr *TypeGhr) IsNil() *ux.State {
//	if state := ux.IfNilReturnError(ghr); state.IsError() {
//		return state
//	}
//	ghr.State = ghr.State.EnsureNotNil()
//	return ghr.State
//}
//
//
//func (ghr *TypeGhr) SetOwner(owner string) *ux.State {
//	ghr.owner = owner
//	return ghr.State
//}
//
//
//func (ghr *TypeGhr) SetRepo(repo string) *ux.State {
//	ghr.repo = repo
//	return ghr.State
//}
//
//
//func (ghr *TypeGhr) Run(cmd string, args ...string) *ux.State {
//
//	for range OnlyOnce {
//		ghr.State.SetPackage("")
//		ghr.State.SetFunction("")
//
//		// A quick hack. To bring in github-release into this code.
//		// Later it'll be merged properly.
//		os.Args = []string{githubRelease, cmd, "-u", ghr.owner, "-r", ghr.repo}
//		os.Args = append(os.Args, args...)
//
//		err := githubrelease.Main()
//		if err != nil {
//			ghr.State.SetError(err)
//		}
//	}
//
//	return ghr.State
//}
