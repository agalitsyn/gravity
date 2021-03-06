/*
Copyright 2018 Gravitational, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package common

import (
	"fmt"
	"io"
	"strings"

	"github.com/fatih/color"
	"github.com/gravitational/trace"
)

// PrintError prints the red error message to the console
func PrintError(err error) {
	color.Red("[ERROR]: %v\n", trace.UserMessage(err))
}

// PrintHeader formats the provided string as a header and prints it to the console
func PrintHeader(val string) {
	fmt.Printf("\n[%v]\n%v\n", val, strings.Repeat("-", len(val)+2))
}

// PrintTableHeader prints header of a table
func PrintTableHeader(w io.Writer, cols []string) {
	dots := make([]string, len(cols))
	for i := range dots {
		dots[i] = strings.Repeat("-", len(cols[i]))
	}
	fmt.Fprint(w, strings.Join(cols, "\t")+"\n")
	fmt.Fprint(w, strings.Join(dots, "\t")+"\n")
}
