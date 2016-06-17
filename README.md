# go-youtube-url [![Build Status](https://travis-ci.org/frozzare/go-youtube-url.svg?branch=master)](https://travis-ci.org/frozzare/go-youtube-url)

Test if url is valid YouTube url.

View the [docs](http://godoc.org/github.com/frozzare/go-youtube-url).

## Installation

```
$ go get github.com/frozzare/go-youtube-url
```

## Example

```go
package main

import (
	"github.com/frozzare/go-youtube-url"
)

func main() {
	youtubeUrl.Valid("https://www.youtube.com/watch?v=Xq7z6WpeB0w")
	//=> true
	
	youtubeUrl.Valid("https://www.example.com/watch?v=Xq7z6WpeB0w")
	//=> false
}
```

# License

MIT © [Fredrik Forsmo](https://github.com/frozzare)
