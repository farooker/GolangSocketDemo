package service

import (
	"fmt"
	utility "simulate_client/Utility"

	"github.com/sacOO7/gowebsocket"
)

type Socket struct
{
  socket interface{}

}
func SocketConnect() {
	socket := gowebsocket.New("wss://" + utility.ConfigServer.Ip + ":" + utility.ConfigServer.Port + "/")
	socket.OnConnected = func(socket gowebsocket.Socket) {
		fmt.Println("Connected to server")
	}
	// socket.ConnectionOptions = gowebsocket.ConnectionOptions {
	//     UseSSL:true,
	//     UseCompression:true,
	//     Subprotocols: [] string{"chat","superchat"},
	// }
	socket.OnConnectError = func(err error, socket gowebsocket.Socket) {
		fmt.Println("Recieved connect error ", err)
	}
	socket.OnTextMessage = func(message string, socket gowebsocket.Socket) {
		fmt.Println("Recieved message " + message)
	}

	socket.OnBinaryMessage = func(data []byte, socket gowebsocket.Socket) {
		fmt.Println("Recieved binary data ", data)
	}

	socket.OnPingReceived = func(data string, socket gowebsocket.Socket) {
		fmt.Println("Recieved ping " + data)
	}

	socket.OnPongReceived = func(data string, socket gowebsocket.Socket) {
		fmt.Println("Recieved pong " + data)
	}

	socket.OnDisconnected = func(err error, socket gowebsocket.Socket) {
		fmt.Println("Disconnected from server ")
		return
	}
	socket.Connect()
}
func(se *Socket)SocketSend(){


}