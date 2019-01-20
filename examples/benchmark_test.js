const Benchmark = require('benchmark')

const suite = new Benchmark.Suite
suite.add('fib#recursion', () => {
  fibRec(10)
})
.add('fib#loop', () => {
  fibLoop(10)
})
.on('complete', () => {
  console.log(suite[0].toString())
  console.log(suite[1].toString())
})
.run({
  async: true
})

function fibRec(n) {
  if (n <= 1) {
    return n
  }

  return fibRec(n-1) + fibRec(n-2)
}

function fibLoop(n) {
  let f = [0, 1]
  for (let i = 2; i <= n; i++) {
    f[i] = f[i-1] + f[i-2]
  }
  return f[n]
}

