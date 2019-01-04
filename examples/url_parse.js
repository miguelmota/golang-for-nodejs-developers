const url = require('url')
const qs = require('querystring')

const urlstr = 'http://bob:secret@sub.example.com:8080/somepath?foo=bar'

const parsed = url.parse(urlstr)
console.log(parsed.protocol)
console.log(parsed.auth)
console.log(parsed.port)
console.log(parsed.hostname)
console.log(parsed.pathname)
console.log(qs.parse(parsed.search.substr(1)))
