/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"github.com/fredrikssongustav/avanza-cli/avanza/operations"
	"github.com/fredrikssongustav/avanza-cli/cmd/utils"

	"github.com/spf13/cobra"
)

// getChartCmd represents the getChart command
var getChartCmd = &cobra.Command{
	Use:   "getChart",
	Short: "This command gives a chart of the desired orderbook",
	Long: "This command gives a chart of the desired orderbook",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("getChartPosition called")

		resp := operations.GetChartPosition(args[0])
		trend := "📈"
		if resp.ChangePercent < 0 {
			trend = "📉"
		}
		fmt.Println(fmt.Sprintf(`      %s – 1 WEEK – %.2f`,trend, resp.ChangePercent),"%")

		utils.Chart(resp.DataPoints)
	},
}

func init() {
	rootCmd.AddCommand(getChartCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getChartCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getChartCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
