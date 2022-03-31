package completion

const (
	BashCompletionFlag        = `generate-taoctl-completion`
	defaultCompletionFilename = "taoctl_autocomplete"
)

const (
	magic = 1 << iota
	flagZsh
	flagBash
)
