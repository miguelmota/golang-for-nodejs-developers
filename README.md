# Golang for Node.js Developers

> [Node.js](https://nodejs.org/) vs [Golang](https://golang.org/) examples

This guide full of examples is intended for people learning Go that are coming from Node.js, although the vice versa can work too. This is not meant to be a complete guide and it is assumed that you've gone through the [Tour of Go](https://tour.golang.org/) tutorial. This guide is meant to be barely good enough to help you at a high level understand how to do *X* in *Y* and doing further research on your own is of course required.

## Contents

- [Examples](#examples)
  - [comments](#comments)
  - [print](#print)
  - [for](#for)
  - [while](#while)
  - [switch](#switch)
  - [arrays](#arrays)
    - [slice](#arrays)
    - [copy](#arrays)
    - [append](#arrays)
  - [uint8 arrays](#uint8-arrays)
  - [array iteration](#array-iteration)
    - [looping](#array-iteration)
    - [mapping](#array-iteration)
    - [filtering](#array-iteration)
    - [reducing](#array-iteration)
  - [classes](#classes)
    - [constructors](#classes)
    - [instantiation](#classes)
    - ["this"](#classes)
  - [exec (sync)](#exec-sync)
  - [exec (async)](#exec-async)
  - [http server](#http-server)
- [License](#license)

## Examples

All sample code is available in [examples/](examples/)

### Comments
---

#### Node.js

```node
// this is a line comment

/*
 this is a block comment
*/
```

#### Go

```go
package main

func main() {
	// this is a line comment

	/*
	   this is a block comment
	*/
}
```

### Print
---

#### Node.js

```node
console.log('hello world')
console.log('hello %s', 'world')
console.log('hello %d %s', 5, 'worlds')
```

Output

```bash
hello world
hello world
hello 5 worlds
```

#### Go

```go
package main

import "fmt"

func main() {
	fmt.Println("hello world")
	fmt.Printf("hello %s\n", "world")
	fmt.Printf("hello %d %s\n", 5, "worlds")
}
```

Output

```bash
hello world
hello world
hello 5 worlds
```

### for
---

#### Node.js

```node
for (let i = 0; i <= 5; i++) {
  console.log(i)
}
```

Output

```bash
0
1
2
3
4
5
```

#### Go

```go
package main

import "fmt"

func main() {
	for i := 0; i <= 5; i++ {
		fmt.Println(i)
	}
}
```

Output

```bash
0
1
2
3
4
5
```

### while
---

#### Node.js

```node
let i = 0

while (i <= 5) {
  console.log(i)

  i++
}
```

Output

```bash
0
1
2
3
4
5
```

#### Go

```go
package main

import "fmt"

func main() {
	i := 0

	for i <= 5 {
		fmt.Println(i)

		i++
	}
}
```

Output

```bash
0
1
2
3
4
5
```

### switch
---

#### Node.js

```node
const value = 'c'

switch(value) {
  case 'a':
    console.log('A')
    break
  case 'b':
    console.log('B')
    break
  case 'c':
    console.log('C')
    break
  default:
    console.log('default')
}

switch(value) {
  case 'a':
    console.log('A - falling through')
  case 'b':
    console.log('B - falling through')
  case 'c':
    console.log('C - falling through')
  default:
    console.log('default')
}
```

Output

```bash
C
C - falling through
default
```

#### Go

```go
package main

import "fmt"

func main() {
	value := "c"

	switch value {
	case "a":
		fmt.Println("A")
	case "b":
		fmt.Println("B")
	case "c":
		fmt.Println("C")
	default:
		fmt.Println("default")
	}

	switch value {
	case "a":
		fmt.Println("A - falling through")
		fallthrough
	case "b":
		fmt.Println("B - falling through")
		fallthrough
	case "c":
		fmt.Println("C - falling through")
		fallthrough
	default:
		fmt.Println("default")
	}
}
```

Output

```bash
C
C - falling through
default
```

### arrays
---

Examples of slicing, copying, and appending arrays.

#### Node.js

```node
const array = [1, 2, 3, 4, 5]
console.log(array)

const clone = array.slice(0)
console.log(clone)

const sub = array.slice(2,4)
console.log(sub)

const concatenated = clone.concat([6, 7])
console.log(concatenated)
```

Output

```bash
[ 1, 2, 3, 4, 5 ]
[ 1, 2, 3, 4, 5 ]
[ 3, 4 ]
[ 1, 2, 3, 4, 5, 6, 7 ]
```

#### Go

```go
package main

import "fmt"

func main() {
	array := []byte{1, 2, 3, 4, 5}
	fmt.Println(array)

	clone := make([]byte, len(array))
	copy(clone, array)
	fmt.Println(clone)

	sub := array[2:4]
	fmt.Println(sub)

	concatenated := append(array, []byte{6, 7}...)
	fmt.Println(concatenated)
}
```

Output

```bash
[1 2 3 4 5]
[1 2 3 4 5]
[3 4]
[1 2 3 4 5 6 7]
```

### uint8 arrays
---

#### Node.js

```javascript
const array = new Uint8Array(10)
console.log(array)

const offset = 1

array.set([1, 2, 3], offset)
console.log(array)

const sub = array.subarray(2)
console.log(sub)

const sub2 = array.subarray(2,4)
console.log(sub2)

console.log(array.byteLength)
```

Output

```bash
Uint8Array [ 0, 0, 0, 0, 0, 0, 0, 0, 0, 0 ]
Uint8Array [ 0, 1, 2, 3, 0, 0, 0, 0, 0, 0 ]
Uint8Array [ 2, 3, 0, 0, 0, 0, 0, 0 ]
Uint8Array [ 2, 3 ]
10
```

#### Go

```go
package main

import "fmt"

func main() {
	array := make([]uint8, 10)
	fmt.Println(array)

	offset := 1

	copy(array[offset:], []uint8{1, 2, 3})
	fmt.Println(array)

	sub := array[2:]
	fmt.Println(sub)

	sub2 := array[2:4]
	fmt.Println(sub2)

	fmt.Println(len(array))
}
```

Output

```bash
[0 0 0 0 0 0 0 0 0 0]
[0 1 2 3 0 0 0 0 0 0]
[2 3 0 0 0 0 0 0]
[2 3]
10
```

### array iteration
---

Examples of iterating, mapping, filtering, and reducing arrays.

#### Node.js

```node
const array = ['a', 'b', 'c']

array.forEach((value, i) => {
  console.log(i, value)
})

const mapped = array.map(value => {
  return value.toUpperCase()
})

console.log(mapped)

const filtered = array.filter((value, i) => {
  return i % 2 == 0
})

console.log(filtered)

const reduced = array.reduce((acc, value, i) => {
  if (i % 2 == 0) {
    acc.push(value.toUpperCase())
  }

  return acc
}, [])

console.log(reduced)
```

Output

```bash
0 'a'
1 'b'
2 'c'
[ 'A', 'B', 'C' ]
[ 'a', 'c' ]
[ 'A', 'C' ]
```

#### Go

```go
package main

import (
	"fmt"
	"strings"
)

func main() {
	array := []string{"a", "b", "c"}

	for i, value := range array {
		fmt.Println(i, value)
	}

	mapped := make([]string, len(array))
	for i, value := range array {
		mapped[i] = strings.ToUpper(value)
	}

	fmt.Println(mapped)

	var filtered []string
	for i, value := range array {
		if i%2 == 0 {
			filtered = append(filtered, value)
		}
	}

	fmt.Println(filtered)

	var reduced []string
	for i, value := range array {
		if i%2 == 0 {
			reduced = append(filtered, strings.ToUpper(value))
		}
	}

	fmt.Println(reduced)
}
```

Output

```bash
0 a
1 b
2 c
[A B C]
[a c]
[a c C]
```

### Classes
---

Examples of classes, constructors, instantiation, and "this" keyword.

#### Node.js

```node
class Foo {
  constructor(value) {
    this.item = value
  }

  getItem() {
    return this.item
  }

  setItem(value) {
    this.item = value
  }
}

const foo = new Foo('bar')
console.log(foo.item)

foo.setItem('qux')

const item = foo.getItem()
console.log(item)
```

Output

```bash
bar
qux
```

#### Go

(closest thing to a class)

```go
package main

import "fmt"

type Foo struct {
	Item string
}

func NewFoo(value string) *Foo {
	return &Foo{
		Item: value,
	}
}

func (f *Foo) GetItem() string {
	return f.Item
}

func (f *Foo) SetItem(value string) {
	f.Item = value
}

func main() {
	foo := NewFoo("bar")
	fmt.Println(foo.Item)

	foo.SetItem("qux")

	item := foo.GetItem()
	fmt.Println(item)
}
```

Output

```bash
bar
qux
```

### exec (sync)
---

#### Node.js

```javascript
const { execSync } = require('child_process')

const output = execSync(`echo 'hello world'`)

console.log(output.toString())
```

Output

```bash
hello world
```

#### Go

```go
package main

import (
	"fmt"
	"os/exec"
)

func main() {
	output, err := exec.Command("echo", "hello world").Output()
	if err != nil {
		panic(err)
	}

	fmt.Println(string(output))
}
```

Output

```bash
hello world
```

### exec (async)
---

#### Node.js

```javascript
const { exec } = require('child_process')

exec(`echo 'hello world'`, (error, stdout, stderr) => {
  if (error) {
    console.error(err)
  }

  if (stderr) {
    console.error(stderr)
  }

  if (stdout) {
    console.log(stdout)
  }
})
```

Output

```bash
hello world
```

#### Go

```go
package main

import (
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command("echo", "hello world")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}
```

Output

```bash
hello world
```

### http server
---

#### Node.js

```node
const http = require('http')

function handler(request, response) {
  response.writeHead(200, { 'Content-type':'text/plain' })
  response.write('hello world')
  response.end()
}

const server = http.createServer(handler)
server.listen(8080)
```

Output

```bash
$ curl http://localhost:8080
hello world
```

#### Go

```go
package main

import (
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("hello world"))
}

func main() {
	http.HandleFunc("/", handler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
```

Output

```bash
$ curl http://localhost:8080
hello world
```


<!--
#### Node.js

```node
```

Output

```bash
```

#### Go

```go
```

Output

```bash
```
-->

## License

[MIT](LICENSE)