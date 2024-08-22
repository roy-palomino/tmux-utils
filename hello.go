package main

import (
    "bufio"
    "fmt"
    "os"
    "os/exec"
    "strings"
    "github.com/fatih/color"
)

func main() {
    reader := bufio.NewReader(os.Stdin)
    c := color.New(color.FgHiGreen)
    c.EnableColor()
    c.Print("New tmux session name: ")
    sessionName, _ := reader.ReadString('\n')
    sessionName = strings.TrimSpace(sessionName)

    cmd := exec.Command("tmux", "new-session", "-ds", sessionName)
    err := cmd.Run()

    if err != nil {
        fmt.Println("Error trying to create session: ", &err)
        return
    }

    c.Println("Nice!")
}
