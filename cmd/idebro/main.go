package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/robinsl/vigilant-memory/internal/idebro"
)

func main() {
  fmt.Println(len(os.Args), os.Args)

  if len(os.Args) > 1 {
    var command = os.Args[1]
    switch command {
    case "help":
      fmt.Println("idebro help")
    case "set-ide":
      LauchProgram(LoadCmd_SetIde())
    case "set-editor":
      fmt.Println("idebro set-ide")
    case "list":
      fmt.Println("idebro list")
    case "ls":
      fmt.Println("idebro list")
    case "projects":
      fmt.Println("idebro projects")
    case "sync":
      fmt.Println("idebro sync")
    default:
      fmt.Println("idebro")
    }
  } else {
    fmt.Println("idebro No params")
  }
}

func LauchProgram(model tea.Model) {
  if _, err := tea.NewProgram(model).Run(); err != nil {
    fmt.Println("Oh no!", err)
    os.Exit(1)
  }
}

func LoadCmd_SetIde() idebro.ConfigIde {
  return idebro.NewConfigIde("nvim")
}
