const obj = {
  someProperties: {
    'foo': 'bar'
  },
  someMethod: (prop) => {
    return obj.someProperties[prop]
  }
}

let item =  obj.someProperties['foo']
console.log(item)

item = obj.someMethod('foo')
console.log(item)
