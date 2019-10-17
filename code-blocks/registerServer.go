type UserModel struct {
}

func (u UserModel) List(ctx context.Context, void *model.Void) (*model.Users,
	error) {
		// somework
	return nil, nil
}

func (u UserModel) Add(ctx context.Context, user *model.User) (*model.Void,
	error) {
		// somework
	return nil, nil
}
