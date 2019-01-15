const crypto = require('crypto')

const hash = crypto.createHash('sha256').update(Buffer.from('hello')).digest()

console.log(hash.toString('hex'))
