# Rabbit For Go Language

We are using regex2 instead of regex because regex doesn't support the Positive lookahead.

## Install Regex2

```
go get github.com/dlclark/regexp2/...
```

Make sure , you already add the Go Path in your bash profile.

## Usages

First install Rabbit
```
go get github.com/Rabbit-Converter/Rabbit-Go
```

And then you can use like

```go
// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"github.com/Rabbit-Converter/Rabbit-Go"
)

func main() {
  fmt.Println(rabbit.Zg2uni("ျမန္မာစာသည္တို႔စာ"))
  fmt.Println(rabbit.Uni2zg("မြန်မာစာသည်တို့စာ"))
}
```