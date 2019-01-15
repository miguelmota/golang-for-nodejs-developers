const err1 = new Error('some error')

console.log(err1)

class FooError extends Error{
  constructor(message) {
    super(message)
    this.name = 'FooError'
    this.message = message
  }

  toString() {
    return this.message
  }
}

const err2 = new FooError('my custom error')

console.log(err2)
