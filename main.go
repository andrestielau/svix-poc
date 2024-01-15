package main

import (
	"svix-poc/cmd"

	"github.com/samber/lo"
)

func main() { lo.Must0(cmd.Root.Execute()) }
