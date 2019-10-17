 // GetUserServer is the server API for GetUser service.
 type GetUserServer interface {
	List(context.Context, *Void) (*Users, error)
	Add(context.Context, *User) (*Void, error)
 }

 func RegisterGetUserServer(s *grpc.Server, srv GetUserServer) {
	s.RegisterService(&_GetUser_serviceDesc, srv)
}