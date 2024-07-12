package main

func main(){
	httpServer := NewHttpServer(":8080")
	go httpServer.Run()

	gRPCServer := NewGRPCServer(":9000")
	gRPCServer.Run()
}