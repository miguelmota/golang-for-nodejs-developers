function myPromise(value) {
  return new Promise((resolve, reject) => {
    setTimeout(() => {
      resolve('resolved: ' + value)
    }, 1e3)
  })
}

function main() {
  myPromise('foo').then(res => console.log(res)).catch(err => console.err(err))

  Promise.all([
    myPromise('A'),
    myPromise('B'),
    myPromise('C')
  ])
  .then(res => console.log(res))
  .catch(err => console.error(err))
}

main()
