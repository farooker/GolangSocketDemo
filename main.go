package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	utility "simulate_client/Utility"

	//"time"

	"github.com/sacOO7/gowebsocket"
)

func main(){
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	utility.InitializeConfigServer()
	fmt.Printf("ip   :%v \n",utility.ConfigServer.Ip)
	fmt.Printf("port :%v \n",utility.ConfigServer.Port)

	socket := gowebsocket.New("wss://"+utility.ConfigServer.Ip+":"+utility.ConfigServer.Port+"/")
	socket.OnConnected = func(socket gowebsocket.Socket) {
		fmt.Println("Connected to server");
	};
	// socket.ConnectionOptions = gowebsocket.ConnectionOptions {
    //  UseSSL:true,
    //  UseCompression:true,
    //  Subprotocols: [] string{"chat","superchat"},
    // }
	socket.OnConnectError = func(err error, socket gowebsocket.Socket) {
		fmt.Println("Recieved connect error ", err)
	};
	socket.OnTextMessage = func(message string, socket gowebsocket.Socket) {
		fmt.Println("Recieved message " + message)
	};
	
	// socket.OnBinaryMessage = func(data [] byte, socket gowebsocket.Socket) {
	// 	fmt.Println("Recieved binary data ", data)
	// };
	
	// socket.OnPingReceived = func(data string, socket gowebsocket.Socket) {
	// 	fmt.Println("Recieved ping " + data)
	// };
	
	socket.OnPongReceived = func(data string, socket gowebsocket.Socket) {
		fmt.Println("Recieved pong " + data)
	};
	
	socket.OnDisconnected = func(err error, socket gowebsocket.Socket) {
		fmt.Println("Disconnected from server ")
		return
	};
   socket.Connect()
      a:=map[string]string{
	"userId" : "11",
	"orgScopeId" : "",
	"nonce" : "",
	"scope" : "Fake",
	"name" : "",
	"accessToken" : "",
	"profileImgURL" : "",
	"clientVersion": "", 
	"nameOverride": "false",
}
	mapa, _ := json.Marshal(a)
	
    MessageCmd := map[string]string{
	"cmd" : "altAuthenticate",
	"accessId" : "1",
	"data" :string(mapa),
	"requestId" : "2",
	}
	mapMessageCmd, _ := json.Marshal(MessageCmd)
	main := map[string]string{
		"key":"xkauthenticate",
		"data":string(mapMessageCmd),
	}
	mapMain, _ := json.Marshal(main)
	socket.SendText(string(mapMain))
	fmt.Println("-------------------------")
	fmt.Printf("send  : %v\n", string(mapMain))
   for {
	select {
	case <-interrupt:
		log.Println("interrupt")
		socket.Close()
		return
	}
}
//    a:=map[string]string{
// 	"userId" : "",
// 	"orgScopeId" : "",
// 	"nonce" : "",
// 	"scope" : "",
// 	"name" : "",
// 	"accessToken" : "",
// 	"profileImgURL" : "",
// 	"clientVersion": "", 
// 	"nameOverride": "",
// }
	
// 	mapa, _ := json.Marshal(a)
// 	mapD := map[string]string{
// 		"key":"xkauthenticate",
// 		"data":string(mapa),
// 	}
// 	mapB, _ := json.Marshal(mapD)
// 	fmt.Printf("%v",string(mapB))
// 	for {
// 		socket.SendText(string(mapB))
// 	    time.Sleep(1*time.Second)
// 	}
}
