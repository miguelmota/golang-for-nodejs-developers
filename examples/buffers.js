const buf = Buffer.alloc(6)

let value = 0x1234567890ab
let offset = 0
let byteLength = 6

buf.writeUIntBE(value, offset, byteLength)

let hexstr = buf.toString('hex')
console.log(hexstr)

const buf2 = Buffer.alloc(6)

value = 0x1234567890ab
offset = 0
byteLength = 6

buf2.writeUIntLE(value, offset, byteLength)

hexstr = buf2.toString('hex')
console.log(hexstr)
