function greet(name = 'stranger') {
  return `hello ${name}`
}

let message = greet()
console.log(message)

message = greet('bob')
console.log(message)
