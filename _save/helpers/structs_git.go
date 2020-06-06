package helpers

//
//
//type TypeGit struct {
//	Url          string
//	Dir          *TypeArgFile
//
//	GitConfig    *gitcmd.Config
//	GitOptions   []string
//
//	skipDirCheck bool
//
//	client       gitcmd.Client
//	repository   *git.Repository
//
//	debug        bool
//	State        *ux.State
//}
//
//
//func NewGit(config *gitcmd.Config, debugMode bool, path ...string) *TypeGit {
//	var p TypeGit
//	p.State = ux.NewState("foo", true)
//
//	for range OnlyOnce {
//		p.State.SetPackage("")
//		p.State.SetFunctionCaller()
//
//		//if config == nil {
//		//	config = &gitcmd.Config{}
//		//}
//		//p.GitConfig = config
//		//p.GitOptions = []string{}
//		//p.Dir = NewArgFile(debugMode)
//		//p.debug = debugMode
//		//
//		//p.State = p.Dir.SetPath(path...)
//		//if p.State.IsNotOk() {
//		//	break
//		//}
//		//
//		//p.client = gitcmd.New(config)
//		//if p.client == nil {
//		//	p.State.SetError("Git error: %s")
//		//	break
//		//}
//
//		p.Url = ""
//		p.repository = nil
//		//p.Cmd = NewExecCommand(debugMode)
//	}
//
//	return &p
//}
//
//
//func (g *TypeGit) IsNil() *ux.State {
//	if state := ux.IfNilReturnError(g); state.IsError() {
//		return state
//	}
//	g.State = g.State.EnsureNotNil()
//	return g.State
//}
//
//
//func (g *TypeGit) IsAvailable() bool {
//	ok := false
//
//	for range OnlyOnce {
//		if g.client == nil {
//			g.State.SetError("git is nil")
//			break
//		}
//
//		err := g.client.CanExec()
//		if err != nil {
//			g.State.SetError("`git` does not exist or its command file is not executable: %s", err)
//			break
//		}
//		g.State.SetOk()
//		ok = true
//	}
//
//	return ok
//}
//func (g *TypeGit) IsNotAvailable() bool {
//	return !g.IsAvailable()
//}
//
//
//func (g *TypeGit) SetConfig(config gitcmd.Config) *ux.State {
//	if state := g.IsNil(); state.IsError() {
//		return state
//	}
//	g.State.SetFunction("")
//
//	for range OnlyOnce {
//		g.GitConfig = &config
//		g.client = gitcmd.New(&config)
//	}
//
//	return g.State
//}
//
//
//func (g *TypeGit) Exec(cmd string, args ...string) *ux.State {
//	if state := g.IsNil(); state.IsError() {
//		return state
//	}
//	g.State.SetFunction("")
//
//	for range OnlyOnce {
//		if g.IsNotAvailable() {
//			break
//		}
//
//		a := g.GitOptions
//		a = append(a, args...)
//
//		//g.State = g.Dir.ChDir()
//		//if g.State.IsNotOk() {
//		//	break
//		//}
//
//		//out, err := g.client.Exec(g.Cmd.GetExe(), g.Cmd.GetArgs()...)
//		out, err := g.client.Exec(cmd, a...)
//		g.State.SetOutput(out)
//		g.State.OutputTrim()
//		g.State.SetError(err)
//		if g.State.IsError() {
//			g.State.SetExitCode(1) // Fake an exit code.
//			break
//		}
//
//		g.State.SetOk()
//	}
//
//	return g.State
//}
//
//
////func (g *TypeGit) FileChanges() []string {
////	var changes []string
////	if state := g.IsNil(); state.IsError() {
////		return changes
////	}
////	g.State.SetFunction("")
////
////	for range OnlyOnce {
////		ux.PrintflnBlue("Checking files in repo...")
////		g.State = g.Exec("status", "--porcelain")
////		if g.State.IsNotOk() {
////			break
////		}
////
////		for _, fp := range g.State.OutputArray {
////			s := strings.Fields(fp)
////			if len(s) > 1 {
////				changes = append(changes, fp)
////			}
////		}
////	}
////
////	return changes
////}
//
//
//func (g *TypeGit) Push(comment string, args ...interface{}) *ux.State {
//	if state := g.IsNil(); state.IsError() {
//		return state
//	}
//	g.State.SetFunction("")
//
//	for range OnlyOnce {
//		c := fmt.Sprintf(comment, args...)
//		if c == "" {
//			g.State.SetError("Missing comment to git commit.")
//			break
//		}
//
//		ux.PrintflnBlue("Adding files to repo...")
//		g.State = g.Exec("add", ".")
//		if g.State.IsNotOk() {
//			break
//		}
//
//		//changes := g.FileChanges()
//		//if len(changes) > 0 {
//		//	ux.PrintflnBlue("Committing files to repo...")
//		//	g.State = g.Exec("commit", "-m", c, ".")
//		//	if g.State.IsNotOk() {
//		//		break
//		//	}
//		//}
//
//		ux.PrintflnBlue("Pushing repo...")
//		g.State = g.Exec("push")
//		if g.State.IsNotOk() {
//			break
//		}
//	}
//
//	return g.State
//}
//
//
//func (g *TypeGit) AddTag(version string, comment string, args ...interface{}) *ux.State {
//	if state := g.IsNil(); state.IsError() {
//		return state
//	}
//	g.State.SetFunction("")
//
//	for range OnlyOnce {
//		if version == "" {
//			g.State.SetError("Missing tag version.")
//			break
//		}
//		if !strings.HasPrefix(version, "v") {
//			version = "v" + version
//		}
//
//		if g.IsTagExisting(version) {
//			g.State.SetOk()
//			break
//		}
//
//
//		c := fmt.Sprintf(comment, args...)
//		if c == "" {
//			c = fmt.Sprintf("Release %s", version)
//		}
//
//		ux.PrintflnBlue("Tagging version %s in repo...", version)
//		g.State = g.Exec("tag", "-a", version, "-m", c)
//		if g.State.IsNotOk() {
//			break
//		}
//
//		ux.PrintflnBlue("Pushing to origin...")
//		g.State = g.Exec("push", "origin", version)
//		if g.State.IsNotOk() {
//			break
//		}
//	}
//
//	return g.State
//}
//
//
//func (g *TypeGit) DelTag(version string) *ux.State {
//	if state := g.IsNil(); state.IsError() {
//		return state
//	}
//	g.State.SetFunction("")
//
//	for range OnlyOnce {
//		if version == "" {
//			g.State.SetError("Missing tag version.")
//			break
//		}
//		if !strings.HasPrefix(version, "v") {
//			version = "v" + version
//		}
//
//		if !g.IsTagExisting(version) {
//			g.State.SetOk()
//			break
//		}
//
//		ux.PrintflnBlue("Removing version tag in repo...")
//		g.State = g.Exec("tag", "-d", version)
//		if g.State.IsNotOk() {
//			break
//		}
//
//		ux.PrintflnBlue("Pushing to origin...")
//		g.State = g.Exec("push", "--delete", "origin", version)
//		if g.State.IsNotOk() {
//			break
//		}
//
//		g.State.SetOk()
//	}
//
//	return g.State
//}
//
//
//func (g *TypeGit) IsTagExisting(version string) bool {
//	var ok bool
//	if state := g.IsNil(); state.IsError() {
//		return false
//	}
//	g.State.SetFunction("")
//
//	for range OnlyOnce {
//		//ux.PrintflnBlue("Checking tag %s in repo...", version)
//		g.State = g.Exec("tag", "-l", version)
//		if g.State.IsNotOk() {
//			break
//		}
//
//		if !strings.HasPrefix(version, "v") {
//			version = "v" + version
//		}
//
//		ok = false
//		for _, t := range g.State.OutputArray {
//			if t == version {
//				ok = true
//				break
//			}
//		}
//	}
//
//	return ok
//}
