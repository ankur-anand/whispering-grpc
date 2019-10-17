type getUserClient struct {
	cc *grpc.ClientConn
}

func NewGetUserClient(cc *grpc.ClientConn) GetUserClient {
	return &getUserClient{cc}
}

func (c *getUserClient) List(ctx context.Context, in *Void, opts ...grpc.CallOption) (*Users, error) {
	out := new(Users)
	err := c.cc.Invoke(ctx, "/model.GetUser/List", in, out, opts...)
	...
	return out, nil
}

func (c *getUserClient) Add(ctx context.Context, in *User, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/model.GetUser/Add", in, out, opts...)
	...
	return out, nil
}