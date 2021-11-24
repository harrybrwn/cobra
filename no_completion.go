//go:build no_cobra_completion

package cobra

func (c *Command) initDefaultCompletionCmd() {}

func (c *Command) initCompleteCmd([]string) {}
