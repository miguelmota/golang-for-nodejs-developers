function asyncMethod(value) {
  return new Promise((resolve, reject) => {
    setTimeout(() => {
      resolve('resolved: ' + value)
    }, 1e3)
  })
}

function main() {
  asyncMethod('foo')
    .then(result => console.log(result))
    .catch(err => console.error(err))

  Promise.all([
    asyncMethod('A'),
    asyncMethod('B'),
    asyncMethod('C')
  ])
  .then(result => console.log(result))
  .catch(err => console.error(err))
}

main()
