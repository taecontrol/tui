package layout

import "github.com/charmbracelet/lipgloss"

type Layout struct {
	height int
	width  int

	header       string
	leftSidebar  string
	body         string
	rightSidebar string
	footer       string
}

func NewLayout(height, width int, body string, opts ...LayoutOption) *Layout {
	layout := &Layout{
		height: height,
		width:  width,
		body:   body,
	}

	for _, opt := range opts {
		opt(layout)
	}

	return layout
}

func (l Layout) Render() string {
	return lipgloss.Place(
		l.width,
		l.height,
		lipgloss.Center,
		lipgloss.Center,
		lipgloss.JoinVertical(
			lipgloss.Center,
			l.headerSection(),
			l.mainSection(),
			l.footerSection(),
		),
	)
}

func (l Layout) headerSection() string {
	return l.header
}

func (l Layout) mainSection() string {
	return lipgloss.Place(
		l.width,
		l.height-lipgloss.Height(l.header)-lipgloss.Height(l.footer),
		lipgloss.Center,
		lipgloss.Center,
		lipgloss.JoinHorizontal(
			lipgloss.Center,
			l.leftSidebar,
			lipgloss.PlaceHorizontal(
				l.width-lipgloss.Width(l.leftSidebar)-lipgloss.Width(l.rightSidebar),
				lipgloss.Center,
				l.body,
			),
			l.rightSidebar,
		),
	)
}

func (l Layout) footerSection() string {
	return l.footer
}

type LayoutOption func(*Layout)

func WithHeader(header string) LayoutOption {
	return func(l *Layout) {
		l.header = header
	}
}

func WithLeftSidebar(leftSidebar string) LayoutOption {
	return func(l *Layout) {
		l.leftSidebar = leftSidebar
	}
}

func WithRightSidebar(rightSidebar string) LayoutOption {
	return func(l *Layout) {
		l.rightSidebar = rightSidebar
	}
}

func WithFooter(footer string) LayoutOption {
	return func(l *Layout) {
		l.footer = footer
	}
}
