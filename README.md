# GoNatSort
A GoLang library to sort slices of strings in Natural Sort order and based on the numerical weightage.

[![Go
Reference](https://pkg.go.dev/badge/github.com/wasim-nihal/gonatsort.svg)](https://pkg.go.dev/github.com/wasim-nihal/gonatsort)

## Usage

1. Get the package:

    ```bash
    go get github.com/wasim-nihal/gonatsort
    ```

2. Import the package in the code

    ```
    package main
    
    import (
    	"fmt"
    	"sort"
    	"github.com/wasim-nihal/gonatsort"
    )
    
    func main() {
    	arr1 := []string{"abc", "aaa", "aaa123", "aaa2", "xyz"}
    	sort.Strings(gonatsort.NatSorter(arr1))
    	// Output: [aaa aaa123 aaa2 abc xyz]
    	fmt.Println(x)
    }
    ```
    
## License
This project is licensed under the Apache-2.0 license License. See the [LICENSE](https://github.com/wasim-nihal/gonatsort/blob/main/LICENSE) file for details.
