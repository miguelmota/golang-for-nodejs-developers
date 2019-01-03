# Golang for Node.js Developers

> [Node.js](https://nodejs.org/) vs [Golang](https://golang.org/) examples

## Contents

- [Examples](#examples)
  - [comments](#comments)
  - [print](#print)
  - [arrays](#arrays)
  - [uint8 arrays](#uint8-arrays)
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

### arrays
---

#### Node.js

```node
const array = [1, 2, 3, 4, 5]
console.log(array)

const clone = array.slice(0)
console.log(clone)

const sub = array.slice(2,4)
console.log(sub)
```

Output

```bash
[ 1, 2, 3, 4, 5 ]
[ 1, 2, 3, 4, 5 ]
[ 3, 4 ]
```

#### Go

```go
package main

import "fmt"

func main() {
	array := []int{1, 2, 3, 4, 5}
	fmt.Println(array)

	clone := make([]int, len(array))
	copy(clone, array)
	fmt.Println(clone)

	sub := array[2:4]
	fmt.Println(sub)
}
```

Output

```bash
[1 2 3 4 5]
[1 2 3 4 5]
[3 4]

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
	array := make([]byte, 10)
	fmt.Println(array)

	offset := 1

	copy(array[offset:], []byte{1, 2, 3})
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

### Classes
---

Examples of classes, constructors, instantiation, and "this" keyword

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