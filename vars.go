package main

// String vars that are free to be changed at compile time
var (
	defaultHistDir = ""
	commonRequirePaths = "';./libs/?/init.lua;./?/init.lua;./?/?.lua'"

	prompt string
	multilinePrompt = "> "
)

// Version info
var (
	ver = "v2.3.4"
	releaseName = "Alyssum"

	gitCommit string
	gitBranch string
)

// Flags
var (
	running bool // Is a command currently running
	interactive bool
	login bool // Are we the login shell?
	noexecute bool // Should we run Lua or only report syntax errors
	initialized bool
)

