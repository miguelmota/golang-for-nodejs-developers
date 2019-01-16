const dgram = require('dgram')
const server = dgram.createSocket('udp4')

server.on('error', err => {
  console.error(err)
  server.close()
})

server.on('message', (msg, rinfo) => {
  const data = msg.toString('utf8').trim()
  console.log(`received: ${data} from ${rinfo.address}:${rinfo.port}`)
})

server.on('listening', () => {
  const address = server.address()
  console.log(`server listening ${address.address}:${address.port}`)
})

server.bind(3000)
