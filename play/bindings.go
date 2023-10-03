package play

import "github.com/charmbracelet/bubbles/key"

type keyBindings struct {
	Input key.Binding
	Left  key.Binding
	Right key.Binding
	Next  key.Binding
}

func (k keyBindings) ShortHelp() []key.Binding {
	return []key.Binding{k.Left, k.Right, k.Input}
}

func (k keyBindings) FullHelp() [][]key.Binding {
	return [][]key.Binding{{k.Left, k.Right, k.Input}}
}

func newBindings() keyBindings {
	return keyBindings{
		Left:  key.NewBinding(key.WithKeys("left"), key.WithHelp("←", "left")),
		Right: key.NewBinding(key.WithKeys("right"), key.WithHelp("→", "right")),
		Input: key.NewBinding(key.WithKeys("enter", " "), key.WithHelp("⮐", "select")),
		Next:  key.NewBinding(key.WithKeys("n"), key.WithHelp("N", "next")),
	}
}
