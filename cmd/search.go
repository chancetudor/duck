package cmd

import (
	"fmt"
	"github.com/anaskhan96/soup"
	"github.com/spf13/cobra"
	"log"
	"net/url"
	"strings"
)

var Site string

func init() {
	rootCmd.AddCommand(searchCmd)
	// flags to fine-tune search command
	// each flag corresponds to standard DDG search syntax
	searchCmd.Flags().BoolP("exact", "e", false, "results for exact query, e.g. cats --exact")
	searchCmd.Flags().BoolP("title", "t", false, "page title includes the query, e.g. cats --title")
	searchCmd.Flags().BoolP("url", "u", false, "page URL includes the query, e.g. cats --url")
	searchCmd.Flags().BoolP("news", "n", false, "returns news about the query, e.g. cincinnati --news")
	searchCmd.Flags().BoolP("map", "m", false, "returns map results about the query, e.g. cincinnati --map")
	searchCmd.Flags().StringVarP(&Site, "site", "s", "", "site to query directly, e.g. cats --site aspca.org")
}

// searchCmd represents the search command
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
		results 			:= search(url)
		printResults(results)
	},
}

func printResults(results []soup.Root) {
	for i := 0; i < len(results); i = i + 1 {
		fmt.Printf("[%d] ", i)
		fmt.Println(results[i].Text(), " :", results[i].Attrs()["href"])
	}
}

func search(url *url.URL) []soup.Root {
	resp, err 				:= soup.Get(url.String())
	if err != nil {
		fmt.Println("An error has occurred fetching the results. Exiting.")
		log.Fatal(err)
	}
	doc 					:= soup.HTMLParse(resp)
	results 				:= doc.FindAll("div", "class", "links_main links_deep result__body")

	return results
}

/* TODO complete
func searchExact(query string) {}

func searchTitle(query string) {}

func searchURL(query string)  {}

func searchNews(query string) {}

func searchMaps(query string) {}
 */

func parseQuery(args []string) string {
	query := fmt.Sprintf(strings.Join(args[:], " "))

	return query
}

func generateURL(query string) *url.URL {
	baseUrl, _ 				:= url.Parse("https://html.duckduckgo.com")
	baseUrl.Path += "/html"
	params 					:= url.Values{}
	params.Add("q", query)
	baseUrl.RawQuery 		= params.Encode()

	return baseUrl
}
