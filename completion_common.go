package cobra

// Annotations for Bash completion.
const (
	BashCompOneRequiredFlag = "cobra_annotation_bash_completion_one_required_flag"
)

// ShellCompDirective is a bit map representing the different behaviors the shell
// can be instructed to have once completions have been provided.
type ShellCompDirective int

const (
	// ShellCompDirectiveError indicates an error occurred and completions should be ignored.
	ShellCompDirectiveError ShellCompDirective = 1 << iota

	// ShellCompDirectiveNoSpace indicates that the shell should not add a space
	// after the completion even if there is a single completion provided.
	ShellCompDirectiveNoSpace

	// ShellCompDirectiveNoFileComp indicates that the shell should not provide
	// file completion even when no completion is provided.
	ShellCompDirectiveNoFileComp

	// ShellCompDirectiveFilterFileExt indicates that the provided completions
	// should be used as file extension filters.
	// For flags, using Command.MarkFlagFilename() and Command.MarkPersistentFlagFilename()
	// is a shortcut to using this directive explicitly.  The BashCompFilenameExt
	// annotation can also be used to obtain the same behavior for flags.
	ShellCompDirectiveFilterFileExt

	// ShellCompDirectiveFilterDirs indicates that only directory names should
	// be provided in file completion.  To request directory names within another
	// directory, the returned completions should specify the directory within
	// which to search.  The BashCompSubdirsInDir annotation can be used to
	// obtain the same behavior but only for flags.
	ShellCompDirectiveFilterDirs

	// ===========================================================================

	// All directives using iota should be above this one.
	// For internal use.
	shellCompDirectiveMaxValue

	// ShellCompDirectiveDefault indicates to let the shell perform its default
	// behavior after completions have been provided.
	// This one must be last to avoid messing up the iota count.
	ShellCompDirectiveDefault ShellCompDirective = 0
)

// CompletionOptions are the options to control shell completion
type CompletionOptions struct {
	// DisableDefaultCmd prevents Cobra from creating a default 'completion' command
	DisableDefaultCmd bool
	// DisableNoDescFlag prevents Cobra from creating the '--no-descriptions' flag
	// for shells that support completion descriptions
	DisableNoDescFlag bool
	// DisableDescriptions turns off all completion descriptions for shells
	// that support them
	DisableDescriptions bool
}

// NoFileCompletions can be used to disable file completion for commands that should
// not trigger file completions.
func NoFileCompletions(cmd *Command, args []string, toComplete string) ([]string, ShellCompDirective) {
	return nil, ShellCompDirectiveNoFileComp
}
