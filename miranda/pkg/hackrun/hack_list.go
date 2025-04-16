package hackrun

import (
	"fmt"
	"io"
	"strings"

	"github.com/beecorrea/weaves/sun"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// List static components
var (
	title = "Which hack do you want to run?"
)

// List styling
var (
	listHeight        = 10
	defaultWidth      = len(title) * 2
	titleStyle        = lipgloss.NewStyle().MarginLeft(2)
	itemStyle         = lipgloss.NewStyle().PaddingLeft(4)
	selectedItemStyle = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("170"))
	helpStyle         = list.DefaultStyles().HelpStyle.PaddingLeft(4)
)

type Item sun.Hack

func (i Item) FilterValue() string {
	return i.Path
}

type ItemDelegate struct{}

func (d ItemDelegate) Height() int                             { return 1 }
func (d ItemDelegate) Spacing() int                            { return 0 }
func (d ItemDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }
func (d ItemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	hack, ok := listItem.(Item)
	if !ok {
		return
	}

	li := fmt.Sprintf("%d) %s", index+1, hack.Name)

	fn := itemStyle.Render
	if index == m.Index() {
		fn = func(s ...string) string {
			return selectedItemStyle.Render("> " + strings.Join(s, " "))
		}
	}

	fmt.Fprint(w, fn(li))
}

func NewHackList(hacks []*sun.Hack) list.Model {
	items := make([]list.Item, 0)
	for _, h := range hacks {
		items = append(items, Item(*h))
	}

	l := list.New(items, ItemDelegate{}, defaultWidth, listHeight)

	l.Title = title
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	l.Styles.Title = titleStyle
	l.Styles.HelpStyle = helpStyle

	return l
}
