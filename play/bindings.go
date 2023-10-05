package play

import "github.com/charmbracelet/bubbles/key"

type keyBindings struct {
	short      key.Binding
	line       key.Binding
	Up         key.Binding
	Down       key.Binding
	Left       key.Binding
	Right      key.Binding
	Input      key.Binding
	Next       key.Binding
	Help       key.Binding
	Learn      key.Binding
	CloseLearn key.Binding
	Quit       key.Binding
}

func (k keyBindings) ShortHelp() []key.Binding {
	return []key.Binding{k.short, k.line, k.CloseLearn, k.Help}
}

func (k keyBindings) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Left, k.Right, k.Up, k.Down, k.Input, k.Next},
		{
			copyKey(k.Help, "close bindings help"),
			k.Learn,
			k.CloseLearn,
			k.Quit,
		},
	}
}

func newBindings() keyBindings {
	learnB := key.NewBinding(key.WithKeys("l", "i"), key.WithHelp("l/i", "open learn info"))
	return keyBindings{
		Left:       key.NewBinding(key.WithKeys("left", "a"), key.WithHelp("←/a", "left answer")),
		Up:         key.NewBinding(key.WithKeys("up", "w"), key.WithHelp("↑/w", "move line/link up")),
		Down:       key.NewBinding(key.WithKeys("down", "s"), key.WithHelp("↓/s", "move line/link down")),
		Right:      key.NewBinding(key.WithKeys("right", "d"), key.WithHelp("→/d", "right answer")),
		Input:      key.NewBinding(key.WithKeys("enter", " ", "f"), key.WithHelp("⮐ / /f", "select answer/open link")),
		Next:       key.NewBinding(key.WithKeys("n"), key.WithHelp("N", "next answer")),
		Help:       key.NewBinding(key.WithKeys("h", "?"), key.WithHelp("h/?", "see key bindings")),
		Learn:      learnB,
		CloseLearn: copyDisabled(learnB, "close learn info"),
		Quit:       key.NewBinding(key.WithKeys("q", "esc"), key.WithHelp("q/esc", "quit app")),
		short:      key.NewBinding(key.WithKeys("c"), key.WithHelp("←/→/↑/↓/⮐ ", "select")),
	}
}

func copyKey(k key.Binding, desc string) key.Binding {
	k.SetHelp(k.Help().Key, desc)

	return k
}

func copyDisabled(k key.Binding, desc string) key.Binding {
	k.SetHelp(k.Help().Key, desc)
	k.SetEnabled(false)

	return k
}
