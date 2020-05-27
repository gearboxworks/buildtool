package cmd

import (
	"github.com/newclarity/buildtool/ux"
	"github.com/tsuyoshiwada/go-gitcmd"
	"gopkg.in/src-d/go-git.v4"
)


type TypeGit struct {
	Url          string
	Dir          *TypeArgFile

	GitConfig    *gitcmd.Config
	GitOptions   []string

	skipDirCheck bool

	client       gitcmd.Client
	repository   *git.Repository

	debug        bool
	State        *ux.State
}


func NewGit(config *gitcmd.Config, debugMode bool, path ...string) *TypeGit {
	var p TypeGit
	p.State = ux.NewState(debugMode)

	for range OnlyOnce {
		p.State.SetPackage("")
		p.State.SetFunctionCaller()

		if config == nil {
			config = &gitcmd.Config{}
		}
		p.GitConfig = config
		p.GitOptions = []string{}
		p.Dir = NewArgFile(debugMode)
		p.debug = debugMode

		p.State = p.Dir.SetPath(path...)
		if p.State.IsNotOk() {
			break
		}

		p.client = gitcmd.New(config)
		if p.client == nil {
			p.State.SetError("Git error: %s")
			break
		}

		p.Url = ""
		p.repository = nil
		//p.Cmd = NewExecCommand(debugMode)
	}

	return &p
}


func (g *TypeGit) IsNil() *ux.State {
	if state := ux.IfNilReturnError(g); state.IsError() {
		return state
	}
	g.State = g.State.EnsureNotNil()
	return g.State
}


func (g *TypeGit) IsAvailable() bool {
	ok := false

	for range OnlyOnce {
		if g.client == nil {
			g.State.SetError("git is nil")
			break
		}

		err := g.client.CanExec()
		if err != nil {
			g.State.SetError("`git` does not exist or its command file is not executable: %s", err)
			break
		}
		g.State.SetOk()
		ok = true
	}

	return ok
}
func (g *TypeGit) IsNotAvailable() bool {
	return !g.IsAvailable()
}


func (g *TypeGit) Open() *ux.State {
	if state := g.IsNil(); state.IsError() {
		return state
	}
	g.State.SetFunction("")

	for range OnlyOnce {
		if g.IsNotAvailable() {
			break
		}

		g.State = g.Dir.ChDir()
		if g.State.IsError() {
			break
		}

		g.State = g.Exec("rev-parse", "--is-inside-work-tree")
		if !g.State.OutputEquals("true") {
			if g.State.IsError() {
				g.State.SetError("current directory does not contain a valid .Git repository: %s", g.State.GetError())
				break
			}

			g.State.SetError("current directory does not contain a valid Git repository")
			break
		}

		var err error
		g.repository, err = git.PlainOpen(g.Dir.Path)
		if err != nil {
			g.State.SetError(err)
			break
		}

		c, _ := g.repository.Config()
		g.Url = c.Remotes["origin"].URLs[0]

		g.State.SetOk("Opened directory %s.\nRemote origin is set to %s\n", g.Dir.Path, g.Url)
	}

	return g.State
}


func (g *TypeGit) SetConfig(config gitcmd.Config) *ux.State {
	if state := g.IsNil(); state.IsError() {
		return state
	}
	g.State.SetFunction("")

	for range OnlyOnce {
		g.GitConfig = &config
		g.client = gitcmd.New(&config)
	}

	return g.State
}


func (g *TypeGit) Exec(cmd string, args ...string) *ux.State {
	if state := g.IsNil(); state.IsError() {
		return state
	}
	g.State.SetFunction("")

	for range OnlyOnce {
		if g.IsNotAvailable() {
			break
		}

		a := g.GitOptions
		a = append(a, args...)

		g.State = g.Dir.ChDir()
		if g.State.IsNotOk() {
			break
		}

		//out, err := g.client.Exec(g.Cmd.GetExe(), g.Cmd.GetArgs()...)
		out, err := g.client.Exec(cmd, a...)
		g.State.SetOutput(out)
		g.State.OutputTrim()
		g.State.SetError(err)
		if g.State.IsError() {
			g.State.SetExitCode(1) // Fake an exit code.
			break
		}

		g.State.SetOk()
	}

	return g.State
}
