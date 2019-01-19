const test = require('tape')

test(t => {
  const tt = [
		{a:1, b:1, ret:2},
		{a:2, b:3, ret:5},
		{a:5, b:5, ret:10}
  ]

  t.plan(tt.length)

  tt.forEach(tt => {
    t.equal(sum(tt.a, tt.b), tt.ret)
  })
})

function sum(a, b) {
	return a + b
}
