function foo() {
  throw new Error('failed')
}

try {
  foo()
} catch(err) {
  console.trace(err)
}

