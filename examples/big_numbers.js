const BN = require('bn.js')

let bn = new BN(75)
console.log(bn.toString(10))

bn = new BN('75')
console.log(bn.toString(10))

bn = new BN(0x4b, 'hex')
console.log(bn.toString(10))

bn = new BN('4b', 'hex')
console.log(bn.toString(10))

bn = new BN(Buffer.from('4b', 'hex'))
console.log(bn.toString(10))
console.log(bn.toNumber(10))
console.log(bn.toString('hex'))

let bn2 = new BN(5)
let isEqual = bn.cmp(bn2) == 0
console.log(isEqual)

bn2 = new BN('4b', 'hex')
isEqual = bn.cmp(bn2) == 0
console.log(isEqual)
