<h1 align="center">
  <br />
  <img src="https://user-images.githubusercontent.com/168240/51193345-801fa380-189d-11e9-92ef-2101bf6a24ed.png" alt="logo" width="700" />
  <br />
  <br />
  <br />
</h1>

# Golang for Node.js Developers

> Examples of [Golang](https://golang.org/) examples compared to [Node.js](https://nodejs.org/) for learning

This guide full of examples is intended for people learning Go that are coming from Node.js, although the vice versa can work too. This is not meant to be a complete guide and it is assumed that you've gone through the [Tour of Go](https://tour.golang.org/) tutorial. This guide is meant to be barely good enough to help you at a high level understand how to do *X* in *Y* and doing further research on your own is of course required.

## Contents

- [Examples](#examples)
  - [comments](#comments)
  - [printing](#printing)
  - [variables](#variables)
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
    - [compare](#buffers)
    - [equals](#buffers)
  - [maps](#maps)
    - [iteration](#maps)
  - [objects](#objects)
  - [destructuring](#destructuring)
  - [spread operator](#spread-operator)
  - [rest operator](#rest-operator)
  - [classes](#classes)
    - [constructors](#classes)
    - [instantiation](#classes)
    - ["this"](#classes)
  - [timeout](#timeout)
  - [interval](#interval)
  - [IIFE](#iife)
  - [files](#files)
    - [creating](#files)
    - [opening](#files)
    - [writing](#files)
    - [reading](#files)
    - [closing](#files)
    - [deleting](#files)
    - [file descriptors](#files)
  - [json](#json)
    - [parse](#json)
    - [stringify](#json)
  - [big numbers](#big-numbers)
    - [uint](#big-numbers)
    - [string](#big-numbers)
    - [hex](#big-numbers)
    - [buffers](#big-numbers)
    - [compare](#buffers)
    - [equals](#buffers)
  - [promises](#promises)
    - [all](#promises)
  - [async/await](#async-await)
  - [streams](#streams)
    - [reading](#streams)
    - [writing](#streams)
  <!--
    - [transform](#streams)
  - [try/catch](#try-catch)
  - [concurrency](#concurrency)
    - [threads](#concurrency)
    - [forking](#concurrency)
  - [message passing](#message-passing)
  - [event emitter](#event-emitter)
  - [first-class functions](#first-class-functions)
  -->
  - [errors](#errors)
  - [exec (sync)](#exec-sync)
  - [exec (async)](#exec-async)
  - [tcp server](#tcp-server)
  - [udp server](#udp-server)
  - [http server](#http-server)
  - [url parse](#url-parse)
  - [gzip](#gzip)
    - [compress](#gzip)
    - [decompress](#gzip)
  - [dns](#dns)
    - [ip lookup](#dns)
    - [mx lookup](#dns)
    - [txt lookup](#dns)
  - [crypto](#crypto)
    - [sha256](#crypto)
  - [env vars](#env-vars)
  - [cli args](#cli-args)
  - [stdin](#stdin)
  - [modules](#modules)
  - [stack trace](#stack-trace)
  <!--
  - [exceptions](#exceptions)
  (catch panic)
  - [testing](#testing)
    - [benchmarking](#testing)
  - [tty](#tty)
  - [db](#db)
    - [postgres](#postgres)
  -->
  <!--
  - [jsdoc](#jsdoc)
  -->
- [License](#license)

## Examples

All sample code is available in [examples/](examples/)

### comments
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

### printing
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

### variables
---

#### Node.js

```node
// function scoped
var foo = 'foo'

// block scoped
let bar = 'bar'

// constant
const qux = 'qux'
```

#### Go

(variables are block scoped in Go)

```go
package main

func main() {
	// explicit
	var foo string = "foo"

	// type inferred
	var bar = "foo"

	// shorthand
	baz := "bar"

	// constant
	const qux = "qux"

	_ = foo
	_ = bar
	_ = baz
	_ = qux
}
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

	// closest thing to ternary operator
	isOddLength := "no"
	if len(array)%2 == 1 {
		isOddLength = "yes"
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

const map3 = new Map()
map3.set('foo', 100)
map3.set('bar', 200)
map3.set('baz', 300)

for (let [key, value] of map3) {
  console.log(key, value)
}
```

Output

```bash
true
bar
false
undefined
foo 100
bar 200
baz 300
```

#### Go

```go
package main

import "fmt"

func main() {
	map1 := make(map[string]string)

	map1["foo"] = "bar"

	item, found := map1["foo"]
	fmt.Println(found)
	fmt.Println(item)

	delete(map1, "foo")

	item, found = map1["foo"]
	fmt.Println(found)
	fmt.Println(item)

	map2 := make(map[string]int)
	map2["foo"] = 100
	map2["bar"] = 200
	map2["baz"] = 300

	for key, value := range map2 {
		fmt.Println(key, value)
	}
}
```

Output

```bash
true
bar
false

foo 100
bar 200
baz 300
```

### objects
---

#### Node.js

```node
const obj = {
  someProperties: {
    'foo': 'bar'
  },
  someMethod: (prop) => {
    return obj.someProperties[prop]
  }
}

let item =  obj.someProperties['foo']
console.log(item)

item = obj.someMethod('foo')
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
	SomeProperties map[string]string
}

func NewObj() *Obj {
	return &Obj{
		SomeProperties: map[string]string{
			"foo": "bar",
		},
	}
}

func (o *Obj) SomeMethod(prop string) string {
	return o.SomeProperties[prop]
}

func main() {
	obj := NewObj()

	item := obj.SomeProperties["foo"]
	fmt.Println(item)

	item = obj.SomeMethod("foo")
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

### spread operator
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

### rest operator
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

(closest thing to a class is to use a structure)

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

### files
---

Examples of creating, opening, writing, reading, closing, and deleting files.

#### Node.js

```node
const fs = require('fs')

// create file
fs.closeSync(fs.openSync('test.txt', 'w'))

// open file (returns file descriptor)
const fd = fs.openSync('test.txt', 'r+')

let wbuf = Buffer.from('hello world.')
let rbuf = Buffer.alloc(12)
let off = 0
let len = 12
let pos = 0

// write file
fs.writeSync(fd, wbuf, pos)

// read file
fs.readSync(fd, rbuf, off, len, pos)

console.log(rbuf.toString())

// close file
fs.closeSync(fd)

// delete file
fs.unlinkSync('test.txt')
```

Output

```bash
hello world.
```

#### Go

```go
package main

import (
	"fmt"
	"os"
	"syscall"
)

func main() {
	// create file
	file, err := os.Create("test.txt")
	if err != nil {
		panic(err)
	}

	// close file
	file.Close()

	// open file
	file, err = os.OpenFile("test.txt", os.O_RDWR, 0755)
	if err != nil {
		panic(err)
	}

	// file descriptor
	fd := file.Fd()

	// open file (using file descriptor)
	file = os.NewFile(fd, "test file")

	wbuf := []byte("hello world.")
	rbuf := make([]byte, 12)
	var off int64

	// write file
	if _, err := file.WriteAt(wbuf, off); err != nil {
		panic(err)
	}

	// read file
	if _, err := file.ReadAt(rbuf, off); err != nil {
		panic(err)
	}

	fmt.Println(string(rbuf))

	// close file (using file descriptor)
	if err := syscall.Close(int(fd)); err != nil {
		panic(err)
	}

	// delete file
	if err := os.Remove("test.txt"); err != nil {
		panic(err)
	}
}
```

Output

```bash
hello world.
```

### json
---

Examples of how to parse (unmarshal) and stringify (marshal) JSON.

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

### big numbers
---

Examples of creating big number types from and to uint, string, hex, and buffers.

#### Node.js

```node
const BN = require('bn.js')

let bn = new BN(75)
console.log(bn.toString(10))

bn = new BN('75')
console.log(bn.toString(10))

bn = new BN(0x4b, 'hex')
console.log(bn.toString(10))

bn = new BN('4b', 'hex')
console.log(bn.toString(10))

bn = new BN(Buffer.from('4b', 'hex'))
console.log(bn.toString(10))
console.log(bn.toNumber(10))
console.log(bn.toString('hex'))
console.log(bn.toBuffer())

let bn2 = new BN(5)
let isEqual = bn.cmp(bn2) == 0
console.log(isEqual)

bn2 = new BN('4b', 'hex')
isEqual = bn.cmp(bn2) == 0
console.log(isEqual)
```

Output

```bash
75
75
75
75
75
75
4b
<Buffer 4b>
false
true
```

#### Go

```go
package main

import (
	"encoding/hex"
	"fmt"
	"math/big"
)

func main() {
	bn := new(big.Int)
	bn.SetUint64(75)
	fmt.Println(bn.String())

	bn = new(big.Int)
	bn.SetString("75", 10)
	fmt.Println(bn.String())

	bn = new(big.Int)
	bn.SetUint64(0x4b)
	fmt.Println(bn.String())

	bn = new(big.Int)
	bn.SetString("4b", 16)
	fmt.Println(bn.String())

	bn = new(big.Int)
	bn.SetBytes([]byte{0x4b})
	fmt.Println(bn.String())
	fmt.Println(bn.Uint64())
	fmt.Println(hex.EncodeToString(bn.Bytes()))
	fmt.Println(bn.Bytes())

	bn2 := big.NewInt(5)
	isEqual := bn.Cmp(bn2) == 0
	fmt.Println(isEqual)

	bn2 = big.NewInt(75)
	isEqual = bn.Cmp(bn2) == 0
	fmt.Println(isEqual)
}
```

Output

```bash
75
75
75
75
75
75
4b
[75]
false
true
```

### promises
---

#### Node.js

```node
function myPromise(value) {
  return new Promise((resolve, reject) => {
    setTimeout(() => {
      resolve('resolved: ' + value)
    }, 1e3)
  })
}

function main() {
  myPromise('foo').then(res => console.log(res)).catch(err => console.err(err))

  Promise.all([
    myPromise('A'),
    myPromise('B'),
    myPromise('C')
  ])
  .then(res => console.log(res))
  .catch(err => console.error(err))
}

main()
```

Output

```bash
resolved: foo
[ 'resolved: A', 'resolved: B', 'resolved: C' ]
```

#### Go

(closest thing is to use channels)

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

func myPromise(value string) chan string {
	ch := make(chan string, 1)
	go func() {
		time.Sleep(1 * time.Second)
		ch <- "resolved: " + value
	}()

	return ch
}

func promiseAll(ch ...chan string) []string {
	var wg sync.WaitGroup
	res := make([]string, len(ch))
	for i, c := range ch {
		wg.Add(1)
		go func(j int, s chan string) {
			res[j] = <-s
			wg.Done()
		}(i, c)
	}

	wg.Wait()
	return res
}

func main() {
	res := <-myPromise("foo")
	fmt.Println(res)

	all := promiseAll(
		myPromise("A"),
		myPromise("B"),
		myPromise("C"),
	)

	fmt.Println(all)
}
```

Output

```bash
resolved: foo
[resolved: A resolved: B resolved: C]
```

### async/await
---

#### Node.js

```node
function hello(name) {
  return new Promise((resolve, reject) => {
    setTimeout(() => {
      if (name === 'fail') {
        reject(new Error('failed'))
      } else {
        resolve('hello ' + name)
      }
    }, 1e3)
  })
}

async function main() {
  try {
    let output = await hello('bob')
    console.log(output)

    output = await hello('fail')
    console.log(output)
  } catch(err) {
    console.log(err.message)
  }
}

main()
```

Output

```bash
hello bob
failed
```

#### Go

(closest thing is to use channels)

```go
package main

import (
	"errors"
	"fmt"
	"time"
)

func hello(name string) chan interface{} {
	ch := make(chan interface{}, 1)
	go func() {
		time.Sleep(1 * time.Second)
		if name == "fail" {
			ch <- errors.New("failed")
		} else {
			ch <- "hello " + name
		}
	}()

	return ch
}

func await(ch chan interface{}) (string, error) {
	res := <-ch
	switch v := res.(type) {
	case string:
		return v, nil
	case error:
		return "", v
	default:
		return "", errors.New("unkown")
	}
}

func main() {
	result, err := await(hello("bob"))
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(result)

	result, err = await(hello("fail"))
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(result)
}
```

Output

```bash
hello bob
failed
```

### streams
---

Examples of reading and writing streams

#### Node.js

```node
const { Readable, Writable } = require('stream')

const inStream = new Readable()

inStream.push(Buffer.from('foo'))
inStream.push(Buffer.from('bar'))
inStream.push(null) // end stream
inStream.pipe(process.stdout)

const outStream = new Writable({
  write(chunk, encoding, callback) {
    console.log('received: ' + chunk.toString('utf8'))
    callback()
  }
})

outStream.write(Buffer.from('abc'))
outStream.write(Buffer.from('xyz'))
outStream.end()
```

Output

```bash
foobar
received: abc
received: xyz
```

#### Go

```go
package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"runtime"
)

func main() {
	inStream := new(bytes.Buffer)
	w := bufio.NewWriter(inStream)
	_, err := w.Write([]byte("foo"))
	if err != nil {
		panic(err)
	}
	_, err = w.Write([]byte("bar"))
	if err != nil {
		panic(err)
	}
	err = w.Flush()
	if err != nil {
		panic(err)
	}

	inStream.WriteTo(os.Stdout)
	fmt.Print("\n")

	outStream := new(bytes.Buffer)
	outStream.Write([]byte("abc\n"))
	outStream.Write([]byte("xyc\n"))
	piper, pipew := io.Pipe()

	go func() {
		sc := bufio.NewScanner(piper)
		for sc.Scan() {
			fmt.Println("received: " + sc.Text())
		}
		if err := sc.Err(); err != nil {
			panic(err)
		}

		os.Exit(0)
	}()

	go func() {
		defer pipew.Close()
		io.Copy(pipew, outStream)
	}()

	defer runtime.Goexit()
}
```

Output

```bash
foobar
received: abc
received: xyc
```

### errors
---

#### Node.js

```node
const err1 = new Error('some error')

console.log(err1)

class FooError extends Error{
  constructor(message) {
    super(message)
    this.name = 'FooError'
    this.message = message
  }

  toString() {
    return this.message
  }
}

const err2 = new FooError('my custom error')

console.log(err2)
```

Output

```bash
Error: some error
{ FooError: my custom error }
```

#### Go

```go
package main

import (
	"errors"
	"fmt"
)

type FooError struct {
	s string
}

func (f *FooError) Error() string {
	return f.s
}

func NewFooError(s string) error {
	return &FooError{s}
}

func main() {
	err1 := errors.New("some error")
	fmt.Println(err1)

	err2 := NewFooError("my custom error")
	fmt.Println(err2)
}
```

Output

```bash
some error
my custom error
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

### tcp server
---

#### Node.js

```node
const net = require('net')

function handler(socket) {
	socket.write('Received: ')
	socket.pipe(socket)
}

const server = net.createServer(handler)
server.listen(3000)
```

Output

```bash
$ echo 'hello' | nc localhost 3000
Received: hello
```

#### Go

```go
package main

import (
	"bufio"
	"net"
)

func handler(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)

	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			return
		}

		conn.Write([]byte("Received: "))
		conn.Write([]byte(message))
	}
}

func main() {
	listener, err := net.Listen("tcp", ":3000")
	if err != nil {
		panic(err)
	}

	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}

		go handler(conn)
	}
}
```

Output

```bash
$ echo 'hello' | nc localhost 3000
Received: hello
```

### udp server
---

#### Node.js

```node
const dgram = require('dgram')
const server = dgram.createSocket('udp4')

server.on('error', err => {
  console.error(err)
  server.close()
})

server.on('message', (msg, rinfo) => {
  const data = msg.toString('utf8').trim()
  console.log(`received: ${data} from ${rinfo.address}:${rinfo.port}`)
})

server.on('listening', () => {
  const address = server.address()
  console.log(`server listening ${address.address}:${address.port}`)
})

server.bind(3000)
```

Output

```bash
$ echo 'hello world' > /dev/udp/0.0.0.0/3000

server listening 0.0.0.0:3000
received: hello world from 127.0.0.1:51452
```

#### Go

```go
package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	conn, err := net.ListenUDP("udp", &net.UDPAddr{
		Port: 3000,
		IP:   net.ParseIP("0.0.0.0"),
	})
	if err != nil {
		panic(err)
	}

	defer conn.Close()
	fmt.Printf("server listening %s\n", conn.LocalAddr().String())

	for {
		message := make([]byte, 20)
		rlen, remote, err := conn.ReadFromUDP(message[:])
		if err != nil {
			panic(err)
		}

		data := strings.TrimSpace(string(message[:rlen]))
		fmt.Printf("received: %s from %s\n", data, remote)
	}
}
```

Output

```bash
$ echo 'hello world' > /dev/udp/0.0.0.0/3000

server listening [::]:3000
received: hello world from 127.0.0.1:50275
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

### gzip
---

#### Node.js

```node
const zlib = require('zlib')

const data = Buffer.from('hello world', 'utf-8')

zlib.gzip(data, (err, compressed) => {
  if (err) {
    console.error(err)
  }

  console.log(compressed)

  zlib.unzip(compressed, (err, decompressed) => {
    if (err) {
      console.error(err)
    }

    console.log(decompressed.toString())
  })
})

```

Output

```bash
<Buffer 1f 8b 08 00 00 00 00 00 00 13 cb 48 cd c9 c9 57 28 cf 2f ca 49 01 00 85 11 4a 0d 0b 00 00 00>
hello world
```

#### Go

```go
package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
)

func main() {
	data := []byte("hello world\n")

	compressed := new(bytes.Buffer)
	w := gzip.NewWriter(compressed)
	if _, err := w.Write(data); err != nil {
		panic(err)
	}
	if err := w.Close(); err != nil {
		panic(err)
	}

	fmt.Println(compressed.Bytes())

	decompressed := new(bytes.Buffer)
	r, err := gzip.NewReader(compressed)
	if err != nil {
		panic(err)
	}

	_, err = decompressed.ReadFrom(r)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(decompressed.Bytes()))
}
```

Output

```bash
[31 139 8 0 0 0 0 0 0 255 202 72 205 201 201 87 40 207 47 202 73 225 2 4 0 0 255 255 45 59 8 175 12 0 0 0]
hello world
```

### dns
---

TXT lookup example

#### Node.js

```node
const dns = require('dns')

dns.resolve4('google.com', (err, ips) => {
  if (err) {
    console.error(err)
  }

  console.log(ips)
})

dns.resolveMx('google.com', (err, mx) => {
  if (err) {
    console.error(err)
  }

  console.log(mx)
})

dns.resolveTxt('google.com', (err, txt) => {
  if (err) {
    console.error(err)
  }

  console.log(txt)
})
```

Output

```bash
[ '172.217.11.78' ]
[ { exchange: 'alt4.aspmx.l.google.com', priority: 50 },
  { exchange: 'alt2.aspmx.l.google.com', priority: 30 },
  { exchange: 'alt3.aspmx.l.google.com', priority: 40 },
  { exchange: 'aspmx.l.google.com', priority: 10 },
  { exchange: 'alt1.aspmx.l.google.com', priority: 20 } ]
[ [ 'v=spf1 include:_spf.google.com ~all' ],
  [ 'docusign=05958488-4752-4ef2-95eb-aa7ba8a3bd0e' ],
  [ 'facebook-domain-verification=22rm551cu4k0ab0bxsw536tlds4h95' ],
  [ 'globalsign-smime-dv=CDYX+XFHUw2wml6/Gb8+59BsH31KzUr6c1l2BPvqKX8=' ] ]
```

#### Go

```go
package main

import (
	"fmt"
	"net"
)

func main() {
	ips, err := net.LookupIP("google.com")
	if err != nil {
		panic(err)
	}

	fmt.Println(ips)

	mx, err := net.LookupMX("google.com")
	if err != nil {
		panic(err)
	}

	fmt.Println(mx)

	txt, err := net.LookupTXT("google.com")
	if err != nil {
		panic(err)
	}

	fmt.Println(txt)
}
```

Output

```bash
[172.217.5.78 2607:f8b0:4007:80d::200e]
[0xc0000ba2e0 0xc0000ba260 0xc0000ba2a0 0xc0000ba280 0xc0000ba300]
[facebook-domain-verification=22rm551cu4k0ab0bxsw536tlds4h95 docusign=05958488-4752-4ef2-95eb-aa7ba8a3bd0e v=spf1 include:_spf.google.com ~all globalsign-smime-dv=CDYX+XFHUw2wml6/Gb8+59BsH31KzUr6c1l2BPvqKX8=]
```

### crypto
---

#### Node.js

```node
const crypto = require('crypto')

const hash = crypto.createHash('sha256').update(Buffer.from('hello')).digest()

console.log(hash.toString('hex'))
```

Output

```bash
2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824
```

#### Go

```go
package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func main() {
	hash := sha256.Sum256([]byte("hello"))

	fmt.Println(hex.EncodeToString(hash[:]))
}
```

Output

```bash
2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824
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

### stdin
---

#### Node.js

```node
const stdin = process.openStdin()

process.stdout.write('Enter name: ')

stdin.addListener('data', text => {
  const name = text.toString().trim()
  console.log('Your name is: ' + name)

  stdin.pause()
})
```

Output

```bash
Enter name: bob
Your name is: bob
```

#### Go

```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter name: ")

	text, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	name := strings.TrimSpace(text)
	fmt.Printf("Your name is: %s\n", name)
}
```

Output

```bash
Enter name: bob
Your name is: bob
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

### stack trace
---

#### Node.js

```node
function foo() {
  console.trace(new Error('failed'))
}

foo()
```

Output

```bash
Trace: Error: failed
    at foo (/Users/bob/examples/stack_trace.js:2:17)
    at Object.<anonymous> (/Users/bob/examples/stack_trace.js:5:1)
    at Module._compile (internal/modules/cjs/loader.js:688:30)
    at Object.Module._extensions..js (internal/modules/cjs/loader.js:699:10)
    at Module.load (internal/modules/cjs/loader.js:598:32)
    at tryModuleLoad (internal/modules/cjs/loader.js:537:12)
    at Function.Module._load (internal/modules/cjs/loader.js:529:3)
    at Function.Module.runMain (internal/modules/cjs/loader.js:741:12)
    at startup (internal/bootstrap/node.js:285:19)
    at bootstrapNodeJSCore (internal/bootstrap/node.js:739:3)
    at foo (/Users/bob/examples/stack_trace.js:2:11)
    at Object.<anonymous> (/Users/bob/examples/stack_trace.js:5:1)
    at Module._compile (internal/modules/cjs/loader.js:688:30)
    at Object.Module._extensions..js (internal/modules/cjs/loader.js:699:10)
    at Module.load (internal/modules/cjs/loader.js:598:32)
    at tryModuleLoad (internal/modules/cjs/loader.js:537:12)
    at Function.Module._load (internal/modules/cjs/loader.js:529:3)
    at Function.Module.runMain (internal/modules/cjs/loader.js:741:12)
    at startup (internal/bootstrap/node.js:285:19)
    at bootstrapNodeJSCore (internal/bootstrap/node.js:739:3)
```

#### Go

```go
package main

import "errors"

func foo() {
	panic(errors.New("failed"))
}

func main() {
	foo()
}
```

Output

```bash
panic: failed

goroutine 1 [running]:
main.foo(...)
        /Users/bob/examples/stack_trace.go:6
main.main()
        /Users/bob/examples/stack_trace.go:10 +0x70
exit status 2
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