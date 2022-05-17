package display

import (
	"github.com/alexeyco/simpletable"
	"github.com/toffernator/todo-cli/task"
)

const (
	UNCHECKEDBOX = "[ ]"
	CHECKEDBOX   = "[x]"
)

func Table(tasks []task.Task) string {
	table := simpletable.New()
	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "Name"},
			{Align: simpletable.AlignCenter, Text: "Important"},
			{Align: simpletable.AlignCenter, Text: "Urgent"},
			{Align: simpletable.AlignCenter, Text: "Deadline"},
			{Align: simpletable.AlignCenter, Text: "Overdue"},
		},
	}

	formattedTasks := tasksToRows(tasks)
	for _, t := range formattedTasks {
		row := []*simpletable.Cell{
			{Align: simpletable.AlignRight, Text: t[0]},
			{Align: simpletable.AlignCenter, Text: t[1]},
			{Align: simpletable.AlignCenter, Text: t[2]},
			{Align: simpletable.AlignRight, Text: t[3]},
			{Align: simpletable.AlignCenter, Text: t[4]},
		}

		table.Body.Cells = append(table.Body.Cells, row)
	}

	table.SetStyle(simpletable.StyleMarkdown)
	return table.String()
}

func tasksToRows(tasks []task.Task) [][]string {
	rows := make([][]string, 0)
	for _, t := range tasks {
		row := taskToCells(t)
		rows = append(rows, row)
	}

	return rows
}

func taskToCells(t task.Task) []string {
	importantCheckbox := asCheckbox(t.IsImportant)
	urgentCheckbox := asCheckbox(t.IsUrgent)
	deadline := t.Deadline.Format("Mon, 02-01-2006")
	overdueCheckbox := asCheckbox(t.IsOverdue())

	return []string{t.Name, importantCheckbox, urgentCheckbox, deadline, overdueCheckbox}
}

func asCheckbox(b bool) string {
	checkbox := UNCHECKEDBOX
	if b {
		checkbox = CHECKEDBOX
	}

	return checkbox
}
