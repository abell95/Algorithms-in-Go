package main

import (
	"fmt"
	"strings"
)

func simplifyPath(path string) string {
    outPath := "/"
    pathElements := strings.Split(path, "/")
    for _, x := range pathElements {
        if x == "" {
            continue
        } else if x == "." {
            continue
        } else if x == ".." {
            // strip last dir element
            strippedPath := outPath
            for i := len(strippedPath) - 1; i >= 0; i-- {
                if string(strippedPath[i]) == "/" {
                    // remove trailing slash, so long as it's not the root
                    if len(strippedPath) > 1 {
                        strippedPath = strippedPath[:(len(strippedPath) - 1)]
                    }
                    outPath = strippedPath
                    break
                } else {
                    // strip it back until we're at the last slash
                    strippedPath = strippedPath[:(len(strippedPath) - 1)]
                }
            }
        } else {
            if outPath == "/" {
                outPath += x
            } else {
                outPath += ("/" + x)
            }
        }
    }
    return outPath
}

func main() {
	fmt.Println(simplifyPath("/../"))
	fmt.Println(simplifyPath("/"))
	fmt.Println(simplifyPath("/a/../b/././c///../d///e"))
	fmt.Println(simplifyPath("/home/back/dir/../../././../././user"))
}
