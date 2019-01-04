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
