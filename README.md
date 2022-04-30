# Image-Scrapper


**A basic application to scrap images from a webpage**<br></br>
 **written in Golang using GoQuery and Grab**
 
 
 
 # Pre-requisites

## [GoQuery](https://github.com/PuerkitoBio/goquery)  
goquery brings a syntax and a set of features similar to [jQuery](http://jquery.com/) to the [Go language](http://golang.org/). It is based on Go's [net/html package](https://pkg.go.dev/golang.org/x/net/html) and the CSS Selector library [cascadia](https://github.com/andybalholm/cascadia). Since the net/html parser returns nodes, and not a full-featured DOM tree, jQuery's stateful manipulation functions (like height(), css(), detach()) have been left off.
 ### Installation
	$ go get github.com/PuerkitoBio/goquery



## [Grab](https://github.com/cavaliergopher/grab)
Grab is a Go package for downloading files from the internet with the following rad features:

-   Monitor download progress concurrently
-   Auto-resume incomplete downloads
-   Guess filename from content header or URL path
-   Safely cancel downloads using context.Context
-   Validate downloads using checksums
-   Download batches of files concurrently
### Installation
	$ go get github.com/cavaliergopher/grab/v3
  
### Importing packages

    import (
      "github.com/PuerkitoBio/goquery"

    "github.com/cavaliergopher/grab/v3"

    )





 
 
