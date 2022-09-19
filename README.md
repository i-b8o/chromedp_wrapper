# chromedp_wrapper
Provides a wrapper around [chromedp package](https://github.com/chromedp/chromedp.git)

## Installing

Install in the usual Go way:

```sh
$ go get -u github.com/i-b8o/chromedp_wrapper
```

## Usage

```go
package main

import (
	"fmt"

	chrwr "github.com/i-b8o/chromedp_wrapper"
)

func main() {
	ctx, cancel := chrwr.Init()
	defer cancel()

	c := chrwr.NewChromeWrapper()

	_ = c.OpenURL(ctx, "https://www.google.com/search?q=parsing")

	_ = c.WaitLoaded(ctx)

	results, _ := c.GetStringsSlice(ctx, `Array.prototype.slice.apply( document.getElementsByTagName("h3") ).map((h3)=> h3.innerText)`)

	for _, result := range results {
		fmt.Println(result)
	}

	c.Click(ctx, "#pnnext")

}
```
