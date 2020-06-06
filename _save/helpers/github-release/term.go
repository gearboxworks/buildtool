package githubrelease

func Mark(ok bool) string {
	if ok {
		return "✔"
	} else {
		return "✗"
	}
}
