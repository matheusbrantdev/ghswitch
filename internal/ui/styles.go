package ui

import "github.com/charmbracelet/lipgloss"

var (
	Bold  = lipgloss.NewStyle().Bold(true)
	Muted = lipgloss.NewStyle().Foreground(lipgloss.Color("8"))
	Name  = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("12"))
	Email = lipgloss.NewStyle().Foreground(lipgloss.Color("10"))
	Key   = lipgloss.NewStyle().Foreground(lipgloss.Color("11"))
)
