let bn = 75n;
console.log(bn.toString(10))

bn = BigInt('75')
console.log(bn.toString(10))

bn = BigInt(0x4b)
console.log(bn.toString(10))

bn = BigInt('0x' + '4b')
console.log(bn.toString(10))

bn = BigInt('0x' + Buffer.from('4b', 'hex').toString('hex'))
console.log(bn.toString(10))
console.log(Number(bn))
console.log(bn.toString(16))
console.log(Buffer.from(bn.toString(16), 'hex'))

/**
 * Compare numbers and return -1 (a < b), 0 (a == b), or 1 (a > b)
 *
 * @param {BigInt} a 
 * @param {BigInt} b
 * 
 * @returns number
 */
function cmp(a, b) {
	if (a < b) {
		return -1;
	}

	if (a > b) {
		return 1;
	}

	return 0;
}

let bn2 = BigInt(5)
let isEqual = cmp(bn, bn2) == 0
console.log(isEqual)

bn2 = BigInt('0x' + '4b')
isEqual = cmp(bn, bn2) == 0
console.log(isEqual)
