package goreland

import (
	"fmt"
	"os"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
)

var VERBOSITY = 1

var successIcon = lipgloss.NewStyle().Foreground(lipgloss.Color("2")).Render(" [:)] ")
var errorIcon = lipgloss.NewStyle().Foreground(lipgloss.Color("1")).Render(" [!!] ")
var infoIcon = lipgloss.NewStyle().Foreground(lipgloss.Color("3")).Render(" [=!] ")
var execIcon = lipgloss.NewStyle().Foreground(lipgloss.Color("6")).Render(" [exec] ")

var successContentStyle = lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{Light: "#909090", Dark: "#626262"})
var errorContentStyle = lipgloss.NewStyle().Bold(true).Underline(true)
var infoContentStyle = lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{Light: "#B2B2B2", Dark: "#4A4A4A"})

func LogSuccess(format string, v ...interface{}) {
	content := fmt.Sprintf(format, v...)
	fmt.Println(successIcon + successContentStyle.Render(content))
}

func LogError(format string, v ...interface{}) {
	content := fmt.Sprintf(format, v...)
	fmt.Println(errorIcon + errorContentStyle.Render(content))
}

func LogFatal(format string, v ...interface{}) {
	LogError(format, v...)
	os.Exit(1)
}

func LogInfo(format string, v ...interface{}) {
	content := fmt.Sprintf(format, v...)
	fmt.Println(infoIcon + infoContentStyle.Render(content))
}

func LogBullet(format string, v ...interface{}) {
	style := lipgloss.NewStyle().Foreground(lipgloss.Color("5"))
	fmt.Println(style.Render(" [*] ") + fmt.Sprintf(format, v...))

}

func LogExecSimple(format string, v ...interface{}) {
	content := fmt.Sprintf(format, v...)
	fmt.Println(execIcon + successContentStyle.Render(content))
}

// Deprecated: Use LogExecSimple instead
func LogExec(format string, v ...interface{}) {
	LogExecSimple(format, v...)
}
func LogTable(headers []string, rows [][]string) {

	headerStyle := lipgloss.NewStyle().Bold(true).Align(lipgloss.Center).PaddingLeft(2).PaddingRight(2)
	rowStyle := lipgloss.NewStyle().Align(lipgloss.Left).PaddingLeft(2).PaddingRight(2)

	t := table.New().Border(lipgloss.NormalBorder()).BorderStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("99"))).StyleFunc(func(row, col int) lipgloss.Style {
		if row == 0 {
			return headerStyle
		}
		return rowStyle
	}).Headers(headers...).Rows(rows...)

	fmt.Println(t)
}
