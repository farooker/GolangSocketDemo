package service

import (
	"fmt"
	// utility "simulate_client/Utility"

	"github.com/sacOO7/gowebsocket"
)
 
var Socket gowebsocket.Socket;

func SocketConnect() {
	Socket = gowebsocket.New("ws://127.0.0.1:8080/websocket")
	Socket.ConnectionOptions = gowebsocket.ConnectionOptions {
		UseSSL:true,
        UseCompression:true,
	 }
	Socket.OnConnected = func(socket gowebsocket.Socket) {
		fmt.Println("Connected to server")
	}
	Socket.OnConnectError = func(err error, socket gowebsocket.Socket) {
		fmt.Println("Recieved connect error ", err)
	}
	Socket.OnTextMessage = func(message string, socket gowebsocket.Socket) {
		fmt.Println("Recieved message " + message)
	}

	Socket.OnBinaryMessage = func(data []byte, socket gowebsocket.Socket) {
		fmt.Println("Recieved binary data ", data)
	}

	Socket.OnPingReceived = func(data string, socket gowebsocket.Socket) {
		fmt.Println("Recieved ping " + data)
	}

	Socket.OnPongReceived = func(data string, socket gowebsocket.Socket) {
		fmt.Println("Recieved pong " + data)
	}

	Socket.OnDisconnected = func(err error, socket gowebsocket.Socket) {
		fmt.Println("Disconnected from server ")
		return
	}
	//socket.Connect()
}
func SocketSendText(req string){
	Socket.SendText(req)
}
func SocketSendBinary(req []byte){
	Socket.SendBinary(req)
}
/*
   //Send text message
   socket.SendText("Hi there, this is my sample test message")

   //Send binary data
   token := make([]byte, 4)
   // rand.Read(token) putting some random value in token
   socket.SendBinary(token)
*/