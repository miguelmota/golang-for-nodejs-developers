const array = new Uint8Array(10)
console.log(array)

const offset = 1

array.set([1, 2, 3], offset)
console.log(array)

const sub = array.subarray(2)
console.log(sub)

const sub2 = array.subarray(2,4)
console.log(sub2)

console.log(array)
const value = 9
const start = 5
const end = 10
array.fill(value, start, end)
console.log(array)

console.log(array.byteLength)
