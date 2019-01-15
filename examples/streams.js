const { Writable, Readable } = require('stream')

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
