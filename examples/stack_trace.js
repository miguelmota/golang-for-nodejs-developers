function foo() {
  console.trace(new Error('failed'))
}

foo()
