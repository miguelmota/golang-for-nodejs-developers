const yargs = require('yargs')

const { foo, qux }= yargs.argv
console.log('foo:', foo)
console.log('qux:', qux)
