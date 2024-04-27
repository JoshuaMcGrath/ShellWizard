package pathbehaviour

import (
	"fmt"
	"shell_wizard/input_handler"
	"shell_wizard/tree"
)

type PathBehaviour struct {
}

func (sb *PathBehaviour) executionBehaviour(n *tree.Node) tree.NodeState {
	fmt.Println("Which Path will you take?")
	for _, childNode := range n.Children {
		//add either descritpion of some kind or file names?
		fmt.Println("What you may encounter!", childNode)
	}
	command := input_handler.RequestInput()
	checkForChangeDirectory(command)
	return tree.SUCCESS
}

func checkForChangeDirectory(cmd *input_handler.InputCommand) {
	var cdHit bool = false
	for !cdHit {
		if checkIfSliceContainsString(cmd, "cd") {
			input_handler.PipeResolver(cmd)
			cdHit = true
		} else {
			output := input_handler.PipeResolver(cmd)
			fmt.Println(output)
			fmt.Println("Now which path will you take?")
			cmd = input_handler.RequestInput()
		}
	}
}

func checkIfSliceContainsString(cmd *input_handler.InputCommand, val string) bool {
	for _, perPipeCmd := range cmd.Command {
		if perPipeCmd == val {
			return true
		}
	}
	return false
}
