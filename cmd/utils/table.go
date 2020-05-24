package utils

import (
	"fmt"
	"github.com/fredrikssongustav/avanza-cli/avanza/operations"
	"github.com/olekukonko/tablewriter"
	"os"
)

func MakePositionTable(arr []operations.Position) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"AccountName", "Volume", "ProfitPercent", "AcquiredValue", "Value", "Currency", "LastPrice", "Name"})

	data := [][]string{}

	for _, v := range arr {
		data = append(data, []string{v.AccountName,
			fmt.Sprintf("%f", v.Volume),
			fmt.Sprintf("%f", v.ProfitPercent),
			fmt.Sprintf("%f", v.AcquiredValue),
			fmt.Sprintf("%f", v.Value),
			v.Currency,
			fmt.Sprintf("%f", v.LastPrice),
			v.Name})
	}

	table.AppendBulk(data) // Add Bulk Data
	table.Render()
}
