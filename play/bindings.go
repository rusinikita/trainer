package play

import "github.com/charmbracelet/bubbles/key"

type keyBindings struct {
	Input key.Binding
	Up    key.Binding
	Down  key.Binding
	Left  key.Binding
	Right key.Binding
	Next  key.Binding
	Learn key.Binding
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
		Up:    key.NewBinding(key.WithKeys("up"), key.WithHelp("↑", "up")),
		Down:  key.NewBinding(key.WithKeys("down"), key.WithHelp("↓", "down")),
		Right: key.NewBinding(key.WithKeys("right"), key.WithHelp("→", "right")),
		Input: key.NewBinding(key.WithKeys("enter", " "), key.WithHelp("⮐", "select")),
		Next:  key.NewBinding(key.WithKeys("n"), key.WithHelp("N", "next")),
		Learn: key.NewBinding(key.WithKeys("l"), key.WithHelp("L", "learn")),
	}
}
