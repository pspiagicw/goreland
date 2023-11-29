package goreland

import (
	"fmt"
	"os"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
)

var VERBOSITY = 1

func LogSuccess(format string, v ...interface{}) {
	style := lipgloss.NewStyle().Foreground(lipgloss.Color("2"))
	content := lipgloss.NewStyle().Faint(true).Render(fmt.Sprintf(format, v...))
	fmt.Println(style.Render(" [:)] ") + content)
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
	style := lipgloss.NewStyle().Foreground(lipgloss.Color("3")).Faint(true)
	fmt.Println(style.Render(" [=!] ") + fmt.Sprintf(format, v...))
}

func LogBullet(format string, v ...interface{}) {
	style := lipgloss.NewStyle().Foreground(lipgloss.Color("5"))
	fmt.Println(style.Render(" [*] ") + fmt.Sprintf(format, v...))

}

func LogExecSimple(format string, v ...interface{}) {
	style := lipgloss.NewStyle().Foreground(lipgloss.Color("6"))
	fmt.Println(style.Render(" [exec] ") + fmt.Sprintf(format, v...))
}
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
