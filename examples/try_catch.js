function foo(fail) {
  if (fail) {
    throw Error('my error')
  }
}

function main() {
  try {
    foo(true)
  } catch(err) {
    console.log(`caught error: ${err.message}`)
  }
}

main()
