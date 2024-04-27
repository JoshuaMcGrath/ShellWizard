package input_handler

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type InputCommand struct {
	Command []string
	args    [][]string
}

func ArgInput(inputLine []string) *InputCommand {
	cmd := new(InputCommand)
	for _, fullCommand := range inputLine {
		inputSlice := ReadInput(fullCommand)
		command := inputSlice[0]
		args := inputSlice[1:]
		cmd.Command = append(cmd.Command, command)
		cmd.args = append(cmd.args, args)
	}
	return cmd
}

func ReadInput(input string) []string {
	return strings.Split(input, " ")
}

func RequestInput() *InputCommand {
	fmt.Println("Cast your next spell!")
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("what did you just try??")
	}
	line = strings.TrimSpace(line)
	//check each pipe and split into array of command and args and run input with each array
	return ArgInput(SplitByPipe(line))
}

func SplitByPipe(line string) []string {
	return strings.Split(line, " | ")
}

func PipeResolver(lineCommand *InputCommand) string {
	var output string = ""
	fmt.Println(len(lineCommand.Command))
	for indx, shCommand := range lineCommand.Command {
		if indx == 0 {
			if lineCommand.args[indx] == nil {
				output = RunInput(shCommand, []string{}, "")
			} else {
				output = RunInput(shCommand, lineCommand.args[indx], "")
			}
		} else {
			if lineCommand.args[indx] == nil {
				output = RunInput(shCommand, []string{}, output)
			} else {
				output = RunInput(shCommand, lineCommand.args[indx], output)
			}
		}
	}
	return output
}

func RunInput(inputCommand string, inputArgs []string, prevInput string) string {
	if inputCommand == "cd" {
		if len(inputArgs) != 1 {
			fmt.Printf("cd command requires exactly one argument")
			return ""
		}
		err := os.Chdir(inputArgs[0])
		if err != nil {
			fmt.Printf("Failed to change directory: %s", err)
			return ""
		}
		pwd, _ := os.Getwd()
		fmt.Println("Directory Changed to " + pwd)
		return ""
	}
	cmd := exec.Command(inputCommand, inputArgs...)
	fmt.Println("previnput:", prevInput)
	if prevInput != "" {
		cmd.Stdin = strings.NewReader(prevInput)
	}
	fmt.Println("full command:", cmd)
	//newcmd := exec.Command("bash", "-c", inputCommand)
	//newcmd.Run()
	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		fmt.Printf("This Spell %s is not in your book! Check again young greybeard\n", inputCommand)
		fmt.Println(err)
		return ""
	}
	return out.String()
}

func TryBash() (string, error) {
	fmt.Println("Cast your next spell!")
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("what did you just try??")
		return "", err
	}
	line = strings.TrimSpace(line)
	newcmd := exec.Command("bash", "-c", line)

	var out bytes.Buffer
	newcmd.Stdout = &out
	if err := newcmd.Run(); err != nil {
		fmt.Printf("This Spell %s is not in your book! Check again young greybeard\n", newcmd)
		return "", err
	}
	return out.String(), nil
}
