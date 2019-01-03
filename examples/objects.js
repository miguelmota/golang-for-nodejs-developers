const obj = {
  someProperty: 'bar',
  someMethod: (prop) => {
    return obj[prop]
  }
}

let item =  obj.someProperty
console.log(item)

item = obj.someMethod('someProperty')
console.log(item)
