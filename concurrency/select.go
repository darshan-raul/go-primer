package main

import (
	// "time"
	"fmt"
	// "sync"
)

type Message struct {
	From string
	Payload string

}

type Server struct{
	msgch chan Message

}


func (s *Server) StartandListen(){

	for {
		select {
		case msg := <- s.msgch:
			fmt.Printf("%s",msg.From)
		
		default:
		}
	}
}

func sendMessageToServer( msgch chan Message, payload string){
	msg := Message{
		From: "fdgfgfg",
		Payload: payload,
	}
	msgch <- msg

}


func main(){

	s:= &Server{
		msgch: make(chan Message, 128),
	}
	go s.StartandListen()

	sendMessageToServer(s.msgch, "hi")
	

}



// func main(){
//  now := time.Now()
//  userID := 10

//  respch := make(chan string,128) //buffer the channel, else you will face deadlock

//  wg := &sync.WaitGroup{}

//  go fetchUserData(userID,respch,wg)
//  wg.Add(1)
//  go fetchUserRecommendations(userID,respch,wg) 
//  wg.Add(1)
//  go fetchUserLikes(userID,respch,wg)
//  wg.Add(1)

//  wg.Wait()

//  close(respch)

//  for resp := range respch{

// 	fmt.Println(resp)
//  }
 
//  fmt.Println(time.Since(now))


// }

// func fetchUserData(userID int, respch chan string, wg *sync.WaitGroup) {

// 	time.Sleep(80 * time.Millisecond)
// 	respch <- "user data"
// 	wg.Done()
// }

// func fetchUserRecommendations(userID int,respch chan string, wg *sync.WaitGroup)  {

// 	time.Sleep(120 * time.Millisecond)
// 	respch <- "user rec"
// 	wg.Done()
// }

// func fetchUserLikes(userID int, respch chan string, wg *sync.WaitGroup) {

// 	time.Sleep(50 * time.Millisecond)

// 	respch <- "userlikes"
// 	wg.Done()
// }