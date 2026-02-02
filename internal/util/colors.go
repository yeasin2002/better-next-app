package util

import "github.com/charmbracelet/lipgloss"

var (
	greenStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("2"))
	blueStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("4"))
	cyanStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("6"))
	yellowStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("3"))
	redStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("1"))
	boldStyle   = lipgloss.NewStyle().Bold(true)
)

func Success(s string) string { return greenStyle.Render(s) }
func Info(s string) string    { return cyanStyle.Render(s) }
func Warning(s string) string { return yellowStyle.Render(s) }
func Error(s string) string   { return redStyle.Render(s) }
func Bold(s string) string    { return boldStyle.Render(s) }
func Cyan(s string) string    { return cyanStyle.Render(s) }
func Blue(s string) string    { return blueStyle.Render(s) }
