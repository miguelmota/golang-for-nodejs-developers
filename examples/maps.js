const map = new Map()
map.set('foo', 'bar')

let found = map.has('foo')
console.log(found)

let item = map.get('foo')
console.log(item)

map.delete('foo')

found = map.has('foo')
console.log(found)

item = map.get('foo')
console.log(item)

const map2 = {}
map2['foo'] = 'bar'
item = map2['foo']
delete map2['foo']

const map3 = new Map()
map3.set('foo', 100)
map3.set('bar', 200)
map3.set('baz', 300)

for (let [key, value] of map3) {
  console.log(key, value)
}
