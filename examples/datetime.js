const nowUnix = Date.now()
console.log(nowUnix)

const datestr = '2019-01-17T09:24:23+00:00'
const date = new Date(datestr)
console.log(date.getTime())
console.log(date.toString())

const futureDate = new Date(date)
futureDate.setDate(date.getDate()+14)
console.log(futureDate.toString())
