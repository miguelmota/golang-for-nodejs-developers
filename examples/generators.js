function *generator() {
  yield 'hello'
  yield 'world'
}

let gen = generator()

while (true) {
  let { value, done } = gen.next()
  console.log(value, done)

  if (done) {
    break
  }
}

// alternatively
for (let value of generator()) {
  console.log(value)
}
