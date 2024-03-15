// Copyright 2018 Syed Wasim Nihal (wasim-nihal). All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

package example_test

import (
	"fmt"
	"sort"

	"github.com/wasim-nihal/gonatsort"
)

func main() {
	a1 := []string{"abc", "def", "aab", "aaa123", "aaa2"}
	sort.Strings(gonatsort.NatSorter(s1))
	// output: [aaa123 aaa2 aab abc def]
	fmt.Println(s1)
}
