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
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

// this var stores the site to query when fed the -s flag
var Site string

// this is a struct to store a result containing a Title and Link to print to stdio
type Result struct {
	Title string
	Link string
}

func openBrowser(url string) {
	var cmd string
	args := []string{url}

	switch runtime.GOOS {
	case "darwin":
		cmd = "open"
	case "windows":
		cmd = "rundll32"
		args = []string{"url.dll,FileProtocolHandler", url}
	default: // "linux", "freebsd", "openbsd", "netbsd"
		cmd = "xdg-open"
	}
	err := exec.Command(cmd, args...).Start()
	if err == nil {
		return
	}
}

func getUserInput(results []Result) {
	var response int
	fmt.Println("Enter a result's number to be taken to the result in your browser")
	_, _ = fmt.Scanln(&response)
	openBrowser(results[response-1].Link)
}

func printResults(results []Result) {
	for i := range results {
		// if/else for formatting purposes
		if i == len(results) - 1 {
			fmt.Printf("[%d] ", i + 1)
			fmt.Println(results[i].Title, " <|> ", results[i].Link)
		} else {
			fmt.Printf("[%d] ", i + 1)
			fmt.Println(results[i].Title, " <|> ", results[i].Link, "\n")
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

func buildSiteQuery(args []string, site string) string {
	query := fmt.Sprintf(strings.Join(args[:], " ") + " site:" + site)

	return query
}

func buildMapQuery(args []string) string {
	query := fmt.Sprintf(strings.Join(args[:], " ") + " map")

	return query
}

func buildNewsQuery(args []string) string {
	query := fmt.Sprintf(strings.Join(args[:], " ") + " news")

	return query
}

func buildInURLQuery(args []string) string {
	query := fmt.Sprintf("inurl:" + strings.Join(args[:], " "))

	return query
}

func buildInTitleQuery(args []string) string {
	query := fmt.Sprintf("intitle:" + strings.Join(args[:], " "))

	return query
}

func buildExactQuery(args []string) string {
	query := fmt.Sprintf(strconv.Quote(strings.Join(args[:], " ")))

	return query
}

// parses the query given as a string
func buildDefaultQuery(args []string) string {
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
[0] PAGE TITLE ||| LINK
To visit a link, simply press the number corresponding with the link.
There are a number of flags available to fine-tune search results.
Seeing as you are searching DuckDuckGo, duck provides no assurance that your search results will be accurate.`,
	Run: func(cmd *cobra.Command, args []string) {
		exactSwitch, _ 		:= cmd.Flags().GetBool("exact")
		titleSwitch, _ 		:= cmd.Flags().GetBool("title")
		urlSwitch, _ 		:= cmd.Flags().GetBool("url")
		newsSwitch, _ 		:= cmd.Flags().GetBool("news")
		mapSwitch, _ 		:= cmd.Flags().GetBool("map")
		siteSwitch, _ 		:= cmd.Flags().GetString("site")

		var query string

		switch {
		case exactSwitch:
			query = buildExactQuery(args)
		case titleSwitch:
			query = buildInTitleQuery(args)
		case urlSwitch:
			query = buildInURLQuery(args)
		case newsSwitch:
			query = buildNewsQuery(args)
		case mapSwitch:
			query = buildMapQuery(args)
		case siteSwitch != "":
			query = buildSiteQuery(args, Site)
		default:
			query = buildDefaultQuery(args)
		}

		url := generateURL(query)
		results := search(url.String())
		printResults(results)
		getUserInput(results)
	},
}
