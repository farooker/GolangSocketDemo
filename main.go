package main

import (
	"encoding/json"
	// "fmt"
	// "log"
	"fmt"
	"log"
	"os"
	"os/signal"
	service "simulate_client/Service"
	utility "simulate_client/Utility"
)

type MessageAuthen struct {
	UserId 		  string   `json:"userId"`
	OrgScopeId    string   `json:"orgScopeId"`
	Nonce 		  string   `json:"nonce "`
	Scope 		  string   `json:"scope"` 
	Name 		  string   `json:"name"` 
	AccessToken   string   `json:"accessToken"`
	ProfileImgURL string   `json:"profileImgURL"`
	ClientVersion string   `json:"clientVersion"`
	NameOverride  bool     `json:"nameOverride"`
}

type MessageCmd  struct{
	Cmd 	  string    `json:"cmd"`
 	AccessId  string    `json:"accessId"`
 	Data 	  string 	`json:"data"`
 	RequestId string    `json:"requestId "`
}

type SendData struct{
	Key 	string   `json:"key"`
	Data 	string   `json:"data"`
}


func main(){
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	utility.InitializeConfigServer()
	service.SocketConnect()  
	service.Socket.Connect()

	Authen := &MessageAuthen{
		UserId : "11",
		OrgScopeId : "",
		Nonce : "",
		Scope : "Fake",
		Name : "",
		AccessToken : "",
		ProfileImgURL : "",
		ClientVersion: "", 
		NameOverride: false,
	}

	AuthenStringtify, _ := json.Marshal(Authen)
	
	Cmd  :=&MessageCmd {
		Cmd : "altAuthenticate",
		AccessId : "1",
		Data :string(AuthenStringtify),
		RequestId : "2",
	}
	
	CmdStringtify, _ := json.Marshal(Cmd)
	MessageData  := &SendData{
		Key:"xkauthenticate",
 		Data:string(CmdStringtify),
	}
	MessageDataStringtify, _ := json.Marshal(MessageData)


	
	fmt.Printf("ip   :%v \n",utility.ConfigServer.Ip)
	fmt.Printf("port :%v \n",utility.ConfigServer.Port)
	fmt.Printf(string(MessageDataStringtify))
	//service.Socket.SendText(string(MessageDataStringtify))
	


	for {
		select {
		case <-interrupt:
			log.Println("interrupt")
			service.Socket.Close()
			return
		}
	}

// 	socket := gowebsocket.New("ws://localhost:7777/websocket")
// 	socket.OnConnected = func(socket gowebsocket.Socket) {
// 		fmt.Println("Connected to server");
// 	};
// 	// socket.ConnectionOptions = gowebsocket.ConnectionOptions {
//     //  UseSSL:true,
//     //  UseCompression:true,
//     //  Subprotocols: [] string{"chat","superchat"},
//     // }
// 	socket.OnConnectError = func(err error, socket gowebsocket.Socket) {
// 		fmt.Println("Recieved connect error ", err)
// 	};
// 	socket.OnTextMessage = func(message string, socket gowebsocket.Socket) {
// 		fmt.Println("Recieved message " + message)
// 	};
	
// 	socket.OnBinaryMessage = func(data [] byte, socket gowebsocket.Socket) {
// 		fmt.Println("Recieved binary data ", data)
// 	};
	
// 	socket.OnPingReceived = func(data string, socket gowebsocket.Socket) {
// 		fmt.Println("Recieved ping " + data)
// 	};
	
// 	socket.OnPongReceived = func(data string, socket gowebsocket.Socket) {
// 		fmt.Println("Recieved pong " + data)
// 	};
	
// 	socket.OnDisconnected = func(err error, socket gowebsocket.Socket) {
// 		fmt.Println("Disconnected from server ")
// 		return
// 	};
//    socket.Connect()
//       a:=map[string]string{
// 	"userId" : "11",
// 	"orgScopeId" : "",
// 	"nonce" : "",
// 	"scope" : "Fake",
// 	"name" : "",
// 	"accessToken" : "",
// 	"profileImgURL" : "",
// 	"clientVersion": "", 
// 	"nameOverride": "false",
// }
// 	mapa, _ := json.Marshal(a)
	
//     MessageCmd := map[string]string{
// 	"cmd" : "altAuthenticate",
// 	"accessId" : "1",
// 	"data" :string(mapa),
// 	"requestId" : "2",
// 	}
// 	mapMessageCmd, _ := json.Marshal(MessageCmd)
// 	main := map[string]string{
// 		"key":"xkauthenticate",
// 		"data":string(mapMessageCmd),
// 	}
// 	mapMain, _ := json.Marshal(main)
// 	socket.SendText(string(mapMain))
// 	fmt.Println("-------------------------")
// 	fmt.Printf("send  : %v\n", string(mapMain))
//    for {
// 	select {
// 	case <-interrupt:
// 		log.Println("interrupt")
// 		socket.Close()
// 		return
// 	}
// }
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
