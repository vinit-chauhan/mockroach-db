package main

import "github.com/vinit-chauhan/http-client/cmd"

func init() {
	cmd.RootCmd.AddCommand(cmd.GetCmd)
}

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		panic(err)
	}
}
