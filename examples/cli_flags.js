const yargs = require('yargs')

const { foo='default value', qux=false } = yargs.argv
console.log('foo:', foo)
console.log('qux:', qux)
