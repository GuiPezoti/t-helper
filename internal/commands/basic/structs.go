package basic

type cliCommand struct {
	name 			string
	description 	string
	Callback		func() error
}