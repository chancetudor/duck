/* 🦆

©️Chance Tudor, 2021
Licensed under GPL v3.0 -- https://www.gnu.org/licenses/gpl-3.0.en.html
View the code, edit the code, run the code

 */
package cmd

import (
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command {
	Use:   "duck",
	Short: "Search DuckDuckGo from the command line",
	Long: `Duck is a CLI that allows you to search DuckDuckGo from the command line. Privacy with ease of use -- why isn't this mainstream yet?`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
