/* 🦆

©️Chance Tudor, 2021
Licensed under GPL v3.0 -- https://www.gnu.org/licenses/gpl-3.0.en.html
View the code, edit the code, run the code

 */
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"golang.org/x/net/html"
	"log"
	"net/http"
	"net/url"
	"strings"
)


// this var stores the site to query when fed the -s flag
var Site string

// TODO IMPLEMENT
func printResults(results []string) {}

func search(url *url.URL) /*[]soup.Root*/ {
	fmt.Println("in search")
	resp, err := http.Get(url.String())
	if err != nil {
		fmt.Println("ERROR : ", err.Error())
		log.Fatal(err)
	}
	resp.Close = true
	defer resp.Body.Close()

	htmlTokens := html.NewTokenizer(resp.Body)
	loop:
	for {
		tt := htmlTokens.Next()
		// t := htmlTokens.Token()
		// fmt.Println(t.Data)
		switch tt {
		case html.ErrorToken:
			fmt.Println("End")
			break loop
		case html.StartTagToken:
			linkToken := htmlTokens.Token()
			if linkToken.Data == "a" && linkToken.Attr[0].Val == "result__snippet" {
				fmt.Println("START link!")
				fmt.Println(linkToken.Attr[1].Val)
				fmt.Println(linkToken.String())
				fmt.Println(htmlTokens.Raw())

				/*
				t := htmlTokens.Next()
				_ = t.String()
				t2 := htmlTokens.Token()
				fmt.Println(t2.Data)

				t3 := htmlTokens.Next()
				_ = t3.String()
				t4 := htmlTokens.Token()
				fmt.Println(t4.Data)
				 */
			}
		//case html.TextToken:
		//	t := htmlTokens.Token()
		//	fmt.Println(t.Data)
		}
	}




	// return results
}

/* TODO complete
func searchExact(query string) {}

func searchTitle(query string) {}

func searchURL(query string)  {}

func searchNews(query string) {}

func searchMaps(query string) {}
*/


// parses the query given as a string
func parseQuery(args []string) string {
	query := fmt.Sprintf(strings.Join(args[:], " "))

	return query
}


// generates a URL from the query given
func generateURL(query string) *url.URL {
	baseUrl, _ 				:= url.Parse("https://html.duckduckgo.com")
	baseUrl.Path += "/html"
	params 					:= url.Values{}
	params.Add("q", query)
	baseUrl.RawQuery 		= params.Encode()

	return baseUrl
}

// inits search command (for Cobra) and adds flags to fine-tune search command
// each flag corresponds to standard DDG search syntax
func init() {
	rootCmd.AddCommand(searchCmd)
	searchCmd.Flags().BoolP("exact", "e", false, "results for exact query, e.g. cats --exact")
	searchCmd.Flags().BoolP("title", "t", false, "page title includes the query, e.g. cats --title")
	searchCmd.Flags().BoolP("url", "u", false, "page URL includes the query, e.g. cats --url")
	searchCmd.Flags().BoolP("news", "n", false, "returns news about the query, e.g. cincinnati --news")
	searchCmd.Flags().BoolP("map", "m", false, "returns map results about the query, e.g. cincinnati --map")
	searchCmd.Flags().StringVarP(&Site, "site", "s", "", "site to query directly, e.g. cats --site aspca.org")
}

// searchCmd represents the search command
// the main function for this file
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Use this command to supply a query",
	Long: `This command is what you will use to search DuckDuckGo. Simply enter what you wish to search following the search command and duck will query for you. Output will be the first 10 link results, structured as such:
[0] PAGE TITLE : LINK
To visit a link, simply press the number corresponding with the link.
There are a number of flags available to fine-tune search results.`,
	Run: func(cmd *cobra.Command, args []string) {
		/*
		exactSwitch, _ 		:= cmd.Flags().GetBool("exact")
		siteSwitch, _ 		:= cmd.Flags().GetString("site")
		titleSwitch, _ 		:= cmd.Flags().GetBool("title")
		urlSwitch, _ 		:= cmd.Flags().GetBool("url")
		newsSwitch, _ 		:= cmd.Flags().GetBool("news")
		mapSwitch, _ 		:= cmd.Flags().GetBool("map")
		*/
		var query string

		/*
		switch {
		case exactSwitch:
			fmt.Println("exact")
		case titleSwitch:
			fmt.Println("title")
		case urlSwitch:
			fmt.Println("url")
		case newsSwitch:
			fmt.Println("news")
		case mapSwitch:
			fmt.Println("map")
		case siteSwitch != "":
				fmt.Println(Site)
		}
		*/

		query 				= parseQuery(args)
		url 				:= generateURL(query)
		fmt.Println(url.String())
		search(url)
		// results 			:= search(url)
		// printResults(results)
	},
}
