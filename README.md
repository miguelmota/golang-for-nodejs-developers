# Golang for Node.js Developers

> [Node.js](https://nodejs.org/) vs [Golang](https://golang.org/) examples

This guide full of examples is intended for people learning Go that are coming from Node.js, although the vice versa can work too. This is not meant to be a complete guide and it is assumed that you've gone through the [Tour of Go](https://tour.golang.org/) tutorial. This guide is meant to be barely good enough to help you at a high level understand how to do *X* in *Y* and doing further research on your own is of course required.

## Contents

- [Examples](#examples)
  - [comments](#comments)
  - [print](#print)
  - [if/else](#ifelse)
    - [ternary](#ifelse)
  - [for](#for)
  - [while](#while)
  - [switch](#switch)
  - [arrays](#arrays)
    - [slicing](#arrays)
    - [copying](#arrays)
    - [appending](#arrays)
  - [uint8 arrays](#uint8-arrays)
  - [array iteration](#array-iteration)
    - [looping](#array-iteration)
    - [mapping](#array-iteration)
    - [filtering](#array-iteration)
    - [reducing](#array-iteration)
  - [buffers](#buffers)
    - [allocate](#buffers)
    - [big endian](#buffers)
    - [little endian](#buffers)
    - [hex](#buffers)
    - [equals](#buffers)
  - [maps](#maps)
  - [objects](#objects)
  - [destructuring](#destructuring)
  - [spread](#spread)
  - [rest](#rest)
  - [classes](#classes)
    - [constructors](#classes)
    - [instantiation](#classes)
    - ["this"](#classes)
  - [timeout](#timeout)
  - [interval](#interval)
  - [IIFE](#iife)
  <!--
  - [files](#files)
    - [creating](#files)
    - [opening](#files)
    - [reading](#files)
    - [writing](#files)
    - [closing](#files)
    - [file descriptors](#files)
  -->
  - [json](#json)
    - [parse](#json)
    - [stringify](#json)
  <!--
  - [big numbers](#big-numbers)
  - [async/await](#async-await)
  - [try/catch](#try-catch)
  - [concurrency](#concurrency)
  - [message passing](#message-passing)
  - [event emitter](#event-emitter)
  - [first-class functions](#first-class-functions)
  - [errors](#errors)
  -->
  - [exec (sync)](#exec-sync)
  - [exec (async)](#exec-async)
  <!--
  - [tcp server](#tcp-server)
  -->
  - [http server](#http-server)
  - [url parse](#url-parse)
  <!--
  - [gzip](#gzip)
  - [dns](#dns)
  - [stdin](#stdin)
  -->
  - [env vars](#env-vars)
  - [cli args](#cli-args)
  - [modules](#modules)
  <!--
  - [stack trace](#stack-trace)
  - [tty](#tty)
  - [crypto](#crypto)
  -->
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

### if/else
---

#### Node.js

```node
const array = [1, 2]

if (array) {
  console.log('array exists')
}

if (array.length === 2) {
  console.log('length is 2')
} else if (array.length === 1) {
  console.log('length is 1')
} else {
  console.log('length is other')
}

const isOddLength = array.length % 2 == 1 ? 'yes' : 'no'

console.log(isOddLength)
```

Output

```bash
array exists
length is 2
no
```

#### Go

(there is no ternary operator in Go)

```go
package main

import "fmt"

func main() {
	array := []byte{1, 2}

	if array != nil {
		fmt.Println("array exists")
	}

	if len(array) == 2 {
		fmt.Println("length is 2")
	} else if len(array) == 1 {
		fmt.Println("length is 1")
	} else {
		fmt.Println("length is other")
	}

	var isOddLength string
	if len(array)%2 == 1 {
		isOddLength = "yes"
	} else {
		isOddLength = "no"
	}

	fmt.Println(isOddLength)
}
```

Output

```bash
array exists
length is 2
no
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

(there is no *while* in Go)

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

### buffers
---

Examples of how to allocate a buffer, write in big or little endian format, encode to a hex string, and check if buffers are equal.

#### Node.js

```node
const buf = Buffer.alloc(6)

let value = 0x1234567890ab
let offset = 0
let byteLength = 6

buf.writeUIntBE(value, offset, byteLength)

let hexstr = buf.toString('hex')
console.log(hexstr)

const buf2 = Buffer.alloc(6)

value = 0x1234567890ab
offset = 0
byteLength = 6

buf2.writeUIntLE(value, offset, byteLength)

hexstr = buf2.toString('hex')
console.log(hexstr)

let isEqual = Buffer.compare(buf, buf2) === 0
console.log(isEqual)

isEqual = Buffer.compare(buf, buf) === 0
console.log(isEqual)
```

Output

```bash
1234567890ab
ab9078563412
false
true
```

#### Go

```go
package main

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"
	"reflect"
)

func writeUIntBE(buffer []byte, value, offset, byteLength int64) {
	slice := make([]byte, byteLength)
	val := new(big.Int)
	val.SetUint64(uint64(value))
	valBytes := val.Bytes()

	buf := bytes.NewBuffer(slice)
	err := binary.Write(buf, binary.BigEndian, &valBytes)
	if err != nil {
		log.Fatal(err)
	}

	slice = buf.Bytes()
	slice = slice[int64(len(slice))-byteLength : len(slice)]

	copy(buffer[offset:], slice)
}

func writeUIntLE(buffer []byte, value, offset, byteLength int64) {
	slice := make([]byte, byteLength)
	val := new(big.Int)
	val.SetUint64(uint64(value))
	valBytes := val.Bytes()

	tmp := make([]byte, len(valBytes))
	for i := range valBytes {
		tmp[i] = valBytes[len(valBytes)-1-i]
	}
	copy(slice, tmp)
	copy(buffer[offset:], slice)
}

func main() {
	buf := make([]byte, 6)
	writeUIntBE(buf, 0x1234567890ab, 0, 6)

	fmt.Println(hex.EncodeToString(buf))

	buf2 := make([]byte, 6)
	writeUIntLE(buf2, 0x1234567890ab, 0, 6)

	fmt.Println(hex.EncodeToString(buf2))

	isEqual := reflect.DeepEqual(buf, buf2)
	fmt.Println(isEqual)

	isEqual = reflect.DeepEqual(buf, buf)
	fmt.Println(isEqual)
}
```

Output

```bash
1234567890ab
ab9078563412
false
true
```

### maps
---

#### Node.js

```node
const map = new Map()
map.set('foo', 'bar')

let found = map.has('foo')
console.log(found)

let item = map.get('foo')
console.log(item)

map.delete('foo')

found = map.has('foo')
console.log(found)

item = map.get('foo')
console.log(item)

const map2 = {}
map2['foo'] = 'bar'
item = map2['foo']
delete map2['foo']
```

Output

```bash
true
bar
false
undefined
```

#### Go

```go
package main

import "fmt"

func main() {
	myMap := make(map[string]string)

	myMap["foo"] = "bar"

	item, found := myMap["foo"]
	fmt.Println(found)
	fmt.Println(item)

	delete(myMap, "foo")

	item, found = myMap["foo"]
	fmt.Println(found)
	fmt.Println(item)
}
```

Output

```bash
true
bar
false

```

### objects
---

#### Node.js

```node
const obj = {
  someProperty: 'bar',
  someMethod: (prop) => {
    return obj[prop]
  }
}

let item =  obj.someProperty
console.log(item)

item = obj.someMethod('someProperty')
console.log(item)
```

Output

```bash
bar
bar
```

#### Go

```go
package main

import "fmt"

type Obj struct {
	SomeProperty string
}

func NewObj(someProperty string) *Obj {
	return &Obj{
		SomeProperty: someProperty,
	}
}

func (o *Obj) SomeMethod(prop string) string {
	if prop == "SomeProperty" {
		return o.SomeProperty
	}

	return ""
}

func main() {
	obj := NewObj("bar")

	item := obj.SomeProperty
	fmt.Println(item)

	item = obj.SomeMethod("SomeProperty")
	fmt.Println(item)
}
```

Output

```bash
bar
bar
```

### destructuring
---

#### Node.js

```node
const obj = { key: 'foo', value: 'bar' }

const { key, value } = obj
console.log(key, value)
```

Output

```bash
foo bar
```

#### Go

(closest thing to destructuring is using multiple return values)

```go
package main

import "fmt"

type Obj struct {
	Key   string
	Value string
}

func (o *Obj) Read() (string, string) {
	return o.Key, o.Value
}

func main() {
	obj := Obj{
		Key:   "foo",
		Value: "bar",
	}

	key, value := obj.Read()

	fmt.Println(key, value)
}
```

Output

```bash
foo bar
```

### spread
---

#### Node.js

```node
const array = [1, 2, 3, 4, 5]

console.log(...array)
```

Output

```bash
1 2 3 4 5
```

#### Go

```go
package main

import "fmt"

func main() {
	array := []byte{1, 2, 3, 4, 5}

	var i []interface{}
	for _, value := range array {
		i = append(i, value)
	}

	fmt.Println(i...)
}
```

Output

```bash
1 2 3 4 5
```

### rest
---

#### Node.js

```node
function sum(...nums) {
	let t = 0

	for (let n of nums) {
		t += n
	}

	return t
}

const total = sum(1, 2, 3, 4, 5)
console.log(total)
```

Output

```bash
15
```

#### Go

```go
package main

import "fmt"

func sum(nums ...int) int {
	var t int
	for _, n := range nums {
		t += n
	}

	return t
}

func main() {
	total := sum(1, 2, 3, 4, 5)
	fmt.Println(total)
}
```

Output

```bash
15
```

### classes
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

### timeout
---

#### Node.js

```node
setTimeout(callback, 1e3)

function callback() {
  console.log('called')
}
```

Output

```bash
called
```

#### Go

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func callback() {
	defer wg.Done()
	fmt.Println("called")
}

func main() {
	wg.Add(1)
	time.AfterFunc(1*time.Second, callback)
	wg.Wait()
}
```

Output

```bash
called
```

### interval
---

#### Node.js

```node
let i = 0

const id = setInterval(callback, 1e3)

function callback() {
  console.log('called', i)

  if (i === 3) {
    clearInterval(id)
  }

  i++
}
```

Output

```bash
called 0
called 1
called 2
called 3
```

#### Go

```go
package main

import (
	"fmt"
	"time"
)

func callback(i int) {
	fmt.Println("called", i)
}

func main() {
	ticker := time.NewTicker(1 * time.Second)

	i := 0
	for range ticker.C {
		callback(i)

		if i == 3 {
			ticker.Stop()
			break
		}

		i++
	}
}
```

Output

```bash
called 0
called 1
called 2
called 3
```

### IIFE
---

Immediately invoked function expression

#### Node.js

```node
(function(name) {
  console.log('hello', name)
})('bob')
```

Output

```bash
hello bob
```

#### Go

```go
package main

import "fmt"

func main() {
	func(name string) {
		fmt.Println("hello", name)
	}("bob")
}
```

Output

```bash
hello bob
```

### json
---

Examples of how to parse and stringify (marshal) JSON.

#### Node.js

```node
let jsonstr = '{"foo":"bar"}'

let parsed = JSON.parse(jsonstr)
console.log(parsed)

jsonstr = JSON.stringify(parsed)
console.log(jsonstr)
```

Output

```bash
{ foo: 'bar' }
{"foo":"bar"}
```

#### Go

```go
package main

import (
	"encoding/json"
	"fmt"
)

type T struct {
	Foo string `json:"foo"`
}

func main() {
	jsonstr := `{"foo":"bar"}`

	t := new(T)
	err := json.Unmarshal([]byte(jsonstr), t)
	if err != nil {
		panic(err)
	}

	fmt.Println(t)

	marshalled, err := json.Marshal(t)
	jsonstr = string(marshalled)
	fmt.Println(jsonstr)
}
```

Output

```bash
&{bar}
{"foo":"bar"}
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

### url parse
---

#### Node.js

```node
const url = require('url')
const qs = require('querystring')

const urlstr = 'http://bob:secret@sub.example.com:8080/somepath?foo=bar'

const parsed = url.parse(urlstr)
console.log(parsed.protocol)
console.log(parsed.auth)
console.log(parsed.port)
console.log(parsed.hostname)
console.log(parsed.pathname)
console.log(qs.parse(parsed.search.substr(1)))
```

Output

```bash
http:
bob:secret
8080
sub.example.com
/somepath
{ foo: 'bar' }
```

#### Go

```go
package main

import (
	"fmt"
	"net/url"
)

func main() {
	urlstr := "http://bob:secret@sub.example.com:8080/somepath?foo=bar"

	u, err := url.Parse(urlstr)
	if err != nil {
		panic(err)
	}

	fmt.Println(u.Scheme)
	fmt.Println(u.User)
	fmt.Println(u.Port())
	fmt.Println(u.Hostname())
	fmt.Println(u.Path)
	fmt.Println(u.Query())
}
```

Output

```bash
http
bob:secret
8080
sub.example.com
/somepath
map[foo:[bar]]
```

### env vars
---

#### Node.js

```node
const key = process.env['API_KEY']

console.log(key)
```

Output

```bash
$ API_KEY=foobar node examples/env_vars.js
foobar
```

#### Go

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	key := os.Getenv("API_KEY")

	fmt.Println(key)
}
```

Output

```bash
$ API_KEY=foobar go run examples/env_vars.go
foobar
```

### cli args
---

#### Node.js

```node
const args = process.argv.slice(2)

console.log(args)
```

Output

```bash
$ node examples/cli_args.js foo bar qux
[ 'foo', 'bar', 'qux' ]
```

#### Go

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	fmt.Println(args)
}
```

Output

```bash
$ go run examples/cli_args.go foo bar qux
[foo bar qux]
```

### modules
---

#### Node.js

```node
const moment = require('moment')

const now = moment().unix()
console.log(now)
```

Output

```bash
1546595748
```

Setup

```bash
npm init
npm install moment --save
```

#### Go

```go
package main

import (
	"fmt"

	"github.com/go-shadow/moment"
)

func main() {
	now := moment.New().Now().Unix()
	fmt.Println(now)
}
```

Output

```bash
1546595748
```

Setup

```bash
go mod init
go mod vendor
```

<!--
### title
---

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