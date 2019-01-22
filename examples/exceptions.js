function foo() {
  throw Error('my exception')
}

function main() {
  foo()
}

process.on('uncaughtException', err => {
  console.log(`caught exception: ${err.message}`)
  process.exit(1)
})

main()
