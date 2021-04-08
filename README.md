# duck
## CLI to search DuckDuckGo. Written in Go (for practice and punniness).

### Usage:
	- ./duck [command] 

### Available Commands:
	help : Help about any command, like so:
    	./duck help
	
    	Use "duck [command] --help" for more information about a command, like so:
      		./duck search help
		
  	search : Use this command to supply a query, like so:
    	./duck search cats

### Global Flags:
	-h, --help : help for duck

### Search Flags:
	-e, --exact : results for exact query, e.g. cats --exact
	-h, --help : help for search
	-m, --map : returns map results about the query, e.g. cincinnati --map
	-n, --news : returns news about the query, e.g. cincinnati --news
	-s, --site string : site to query directly, e.g. cats --site aspca.org
	-t, --title : page title includes the query, e.g. cats --title
	-u, --url : page URL includes the query, e.g. cats --url
