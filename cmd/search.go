/* 🦆

©️Chance Tudor, 2021
Licensed under GPL v3.0 -- https://www.gnu.org/licenses/gpl-3.0.en.html
View the code, edit the code, run the code

 */
package cmd

import (
	"fmt"
	"github.com/gocolly/colly/v2"
	"github.com/spf13/cobra"
	"log"
	"net/url"
	"strings"
)


// this var stores the site to query when fed the -s flag
var Site string

// this is a struct to store a result containing a Title and Link to print to stdio
type Result struct {
	Title string
	Link string
}

// TODO IMPLEMENT
func printResults(results []Result) {
	for i := range results {
		if i == len(results) - 1 {
			fmt.Println(results[i].Title, " | ", results[i].Link)
		} else {
			fmt.Println(results[i].Title, " | ", results[i].Link, "\n")
		}
	}
}

func search(url string) []Result {
	var results []Result
	collect := colly.NewCollector()
	collect.OnHTML("a[href]", func(e *colly.HTMLElement) {
		if e.Attr("class") == "result__a" {
			link := cleanLink(e.Attr("href"))
			title := e.Text
			results = append(results, Result{title, link})
		}
	})
	_ = collect.Visit(url)

	return results
}

func cleanLink(dirtyLink string) string {
	// get rid of dumb DDG additions before the good stuff
	cleanerLink := dirtyLink[25:]
	cleanLink, err := url.PathUnescape(cleanerLink)
	if err != nil {
		log.Fatal(err)
	}

	return cleanLink
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
		results 			:= search(url.String())
		printResults(results)
	},
}
