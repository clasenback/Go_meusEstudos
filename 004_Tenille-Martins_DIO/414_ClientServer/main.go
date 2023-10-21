package main

import (
	"bufio"
	"fmt"
	"net/http"
)

func main() {
	resp, err := http.Get("http://gobyexample.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("Qual é o status:", resp.Status)

	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan() && i < 25; i++ {
		fmt.Println(i, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

}

/*
Qual é o status: 200 OK
0 <!DOCTYPE html>
1 <html>
2   <head>
3     <meta charset="utf-8">
4     <title>Go by Example</title>
5     <link rel=stylesheet href="site.css">
6   </head>
7   <body>
8     <div id="intro">
9       <h2><a href="./">Go by Example</a></h2>
10       <p>
11         <a href="https://go.dev">Go</a> is an
12         open source programming language designed for
13         building simple, fast, and reliable software.
14         Please read the official
15         <a href="https://go.dev/doc/tutorial/getting-started">documentation</a>
16         to learn a bit about Go code, tools packages,
17         and modules.
18       </p>
19
20       <p>
21         <em>Go by Example</em> is a hands-on introduction
22         to Go using annotated example programs. Check out
23         the <a href="hello-world">first example</a> or
24         browse the full list below.
*/
