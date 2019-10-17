func main() {
	svc := grpc.NewServer()
	var um UserModel

	// Register our service implementation with the gRPC server.
	RegisterGetUserServer(svc, um) // HL

	l, err := net.Listen("tcp", ":3002")
	if err != nil {
		log.Fatalf("could not start server on [%s]: %w", "3002", err)
	}
	err = svc.Serve(l)
	if err != nil {
		log.Fatal(err)
	}
}