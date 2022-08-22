const dns = require('dns')

dns.resolveNs('google.com', (err, ns) => {
  if (err) {
    console.error(err)
  }

  console.log(ns)
})

dns.resolve4('google.com', (err, ips) => {
  if (err) {
    console.error(err)
  }

  console.log(ips)
})

dns.resolveMx('google.com', (err, mx) => {
  if (err) {
    console.error(err)
  }

  console.log(mx)
})

dns.resolveTxt('google.com', (err, txt) => {
  if (err) {
    console.error(err)
  }

  console.log(txt)
})

setTimeout(() => {
  dns.setServers(['1.1.1.1'])
  console.log(dns.getServers())

  dns.resolveNs('google.com', (err, ns) => {
    if (err) {
      console.error(err)
    }

    console.log(ns)
  })
}, 100)
