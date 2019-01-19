let input = 'foobar'
let replaced = input.replace(/foo(.*)/i, 'qux$1')
console.log(replaced)

let match = /o{2}/i.test(input)
console.log(match)

input = '111-222-333'
let matches = input.match(/([0-9]+)/gi)
console.log(matches)
