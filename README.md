# quickPerm
The [quickperm](http://www.quickperm.org/) algorithm implemented with Go

## Example

```` go
for permutation := range quickPerm.GeneratePermutationsString([]string{"a","b","c","d","e","f","g","h"}) {
    fmt.Println(permutation)
}
````

## Badges

[![GoDoc](https://godoc.org/github.com/Ramshackle-Jamathon/go-quickPerm?status.svg)](https://godoc.org/github.com/Ramshackle-Jamathon/go-quickPerm)
![](https://img.shields.io/badge/license-MIT-blue.svg)
![](https://img.shields.io/badge/status-stable-green.svg)
