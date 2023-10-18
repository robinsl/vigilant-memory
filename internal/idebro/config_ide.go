package idebro

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type ConfigIde struct {
  textInput textinput.Model
  currentIde string
  err error
}

type errMsg error

func NewConfigIde(currentIde string) ConfigIde {
  ti := textinput.New()
  ti.Placeholder = "Ide/Editor name..."
  ti.Focus()

  defaultIde := "vscode"
  if currentIde == "" {
    currentIde = defaultIde
  }

  return ConfigIde{
    textInput: ti,
    currentIde: currentIde,
    err: nil,
  }
}

func (m ConfigIde) Init() tea.Cmd {
  return textinput.Blink
}

func (m ConfigIde) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
  var cmd tea.Cmd

  switch msg := msg.(type) {
  case tea.KeyMsg:  
    switch msg.Type {
    case tea.KeyEnter:
      var err error
      m.currentIde, err = SaveIde(m.textInput.Value())
      if err != nil {
        m.err = err
        return m, nil
      }
      return m, tea.Quit
    case tea.KeyCtrlC, tea.KeyEsc:
      return m, tea.Quit
    }
  case errMsg:
    m.err = msg
    return m, nil
  }

  m.textInput, cmd = m.textInput.Update(msg)

  return m, cmd
}

func (m ConfigIde) View() string {
  return fmt.Sprintf(
    "Your Current ide is: %s\nSelect your new ide/editor:\n\n%s\n\n%s",
    m.currentIde,
    m.textInput.View(),
    "Press esc to exit.",
  ) + "\n"
}

func SaveIde(ide string) (string, error) {
  fmt.Println("Saving ide: ", ide)
  return ide, nil
}
