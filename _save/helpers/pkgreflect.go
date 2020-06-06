package helpers

//const pkgReflect = "pkgreflect"
//
//
//type TypePkgReflect struct {
//	Args  []string
//
//	dir  *TypeArgFile
//
//	debug bool
//	State *ux.State
//}
//
//
//func NewPkgReflect(debugMode bool) *TypePkgReflect {
//	var ghr TypePkgReflect
//	ghr.State = ux.NewState(debugMode)
//
//	for range OnlyOnce {
//		ghr.State.SetPackage("")
//		ghr.State.SetFunctionCaller()
//
//		// A quick hack. To bring in github-release into this code.
//		// Later it'll be merged properly.
//		ghr.Args = []string{"--notests"}
//		ghr.dir = NewArgFile(debugMode)
//	}
//
//	return &ghr
//}
//
//
//func (ghr *TypePkgReflect) IsNil() *ux.State {
//	if state := ux.IfNilReturnError(ghr); state.IsError() {
//		return state
//	}
//	ghr.State = ghr.State.EnsureNotNil()
//	return ghr.State
//}
//
//
//func (ghr *TypePkgReflect) SetPath(path ...string) *ux.State {
//	if len(path) == 0 {
//		ghr.State.SetOk()
//		return ghr.State
//	}
//
//	ghr.dir = NewArgFile(ghr.debug)
//	ghr.State = ghr.dir.SetPath(path...)
//	return ghr.State
//}
//
//
//func (ghr *TypePkgReflect) SetArgs(args ...string) *ux.State {
//	//ghr.Args = []string{pkgReflect}
//	ghr.Args = []string{}
//	ghr.Args = append(ghr.Args, args...)
//	ghr.State.SetOk()
//	return ghr.State
//}
//
//
//func (ghr *TypePkgReflect) AppendArgs(args ...string) *ux.State {
//	ghr.Args = append(ghr.Args, args...)
//	ghr.State.SetOk()
//	return ghr.State
//}
//
//
//func (ghr *TypePkgReflect) Run(paths []string) *ux.State {
//
//	for range OnlyOnce {
//		ghr.State.SetPackage("")
//		ghr.State.SetFunction("")
//
//		for _, p := range paths {
//			// A quick hack. To bring in github-release into this code.
//			// Later it'll be merged properly.
//			if p == "" {
//				continue
//			}
//
//			ghr.SetPath(p)
//			os.Args = []string{pkgReflect}
//			os.Args = append(os.Args, ghr.Args...)
//			os.Args = append(os.Args, ghr.dir.Dir)
//
//			//err := scribe.PackageReflect()
//			//if err != nil {
//			//	ghr.State.SetError(err)
//			//}
//		}
//	}
//
//	return ghr.State
//}
