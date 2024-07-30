package cmd

type Command struct {
	Name     string
	Descp    string
	Callback func() error
}
