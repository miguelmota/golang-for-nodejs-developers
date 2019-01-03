const { execSync } = require('child_process')

const output = execSync(`echo 'hello world'`)

console.log(output.toString()) // hello world
