// GetUserClient is the client API for GetUser service.
//
type GetUserClient interface {
	List(ctx context.Context, in *Void, opts ...grpc.CallOption) (*Users, error)
	Add(ctx context.Context, in *User, opts ...grpc.CallOption) (*Void, error)
}