const net = require('net')

function handler(socket) {
	socket.write('Received: ')
	socket.pipe(socket)
}

const server = net.createServer(handler)
server.listen(3000)
