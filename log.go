package goreland

import (
	"fmt"
	"os"

	"github.com/charmbracelet/lipgloss"
)

var VERBOSITY = 1

func LogSuccess(format string, v ...interface{}) {
	style := lipgloss.NewStyle().Foreground(lipgloss.Color("2"))
	fmt.Println((style.Render(" [:)] ") + fmt.Sprintf(format, v...)))
}

func LogError(format string, v ...interface{}) {
	style := lipgloss.NewStyle().Foreground(lipgloss.Color("1"))
	fmt.Println(style.Render(" [!!] ") + fmt.Sprintf(format, v...))
}

func LogFatal(format string, v ...interface{}) {
	style := lipgloss.NewStyle().Foreground(lipgloss.Color("1"))
	fmt.Println(style.Render(" [!!] ") + fmt.Sprintf(format, v...))
	os.Exit(1)
}

func LogInfo(format string, v ...interface{}) {
	style := lipgloss.NewStyle().Foreground(lipgloss.Color("3"))
	fmt.Println(style.Render(" [=!] ") + fmt.Sprintf(format, v...))
}

func LogBullet(format string, v ...interface{}) {
	style := lipgloss.NewStyle().Foreground(lipgloss.Color("5"))
	fmt.Println(style.Render(" [@!] ") + fmt.Sprintf(format, v...))

}
func LogExec(format string, v ...interface{}) {
	style := lipgloss.NewStyle().Foreground(lipgloss.Color("6"))
	borderStyle := lipgloss.NewStyle().Border(lipgloss.NormalBorder(), true, true, true, true).PaddingLeft(1).PaddingRight(2)
	fmt.Println(borderStyle.Render(style.Render("[exec] ") + fmt.Sprintf(format, v...)))
}

func LogExecSimple(format string, v ...interface{}) {
	style := lipgloss.NewStyle().Foreground(lipgloss.Color("6"))
	fmt.Println(style.Render(" [exec] ") + fmt.Sprintf(format, v...))
}
