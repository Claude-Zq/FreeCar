// Code generated by Kitex v0.4.4. DO NOT EDIT.

package authservice

import (
	"context"
	auth "github.com/CyanAsterisk/FreeCar/server/shared/kitex_gen/auth"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

func serviceInfo() *kitex.ServiceInfo {
	return authServiceServiceInfo
}

var authServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "AuthService"
	handlerType := (*auth.AuthService)(nil)
	methods := map[string]kitex.MethodInfo{
		"Login":               kitex.NewMethodInfo(loginHandler, newAuthServiceLoginArgs, newAuthServiceLoginResult, false),
		"AdminLogin":          kitex.NewMethodInfo(adminLoginHandler, newAuthServiceAdminLoginArgs, newAuthServiceAdminLoginResult, false),
		"ChangeAdminPassword": kitex.NewMethodInfo(changeAdminPasswordHandler, newAuthServiceChangeAdminPasswordArgs, newAuthServiceChangeAdminPasswordResult, false),
		"UploadAvatar":        kitex.NewMethodInfo(uploadAvatarHandler, newAuthServiceUploadAvatarArgs, newAuthServiceUploadAvatarResult, false),
		"GetUser":             kitex.NewMethodInfo(getUserHandler, newAuthServiceGetUserArgs, newAuthServiceGetUserResult, false),
		"AddUser":             kitex.NewMethodInfo(addUserHandler, newAuthServiceAddUserArgs, newAuthServiceAddUserResult, false),
		"DeleteUser":          kitex.NewMethodInfo(deleteUserHandler, newAuthServiceDeleteUserArgs, newAuthServiceDeleteUserResult, false),
		"UpdateUser":          kitex.NewMethodInfo(updateUserHandler, newAuthServiceUpdateUserArgs, newAuthServiceUpdateUserResult, false),
		"GetUsers":            kitex.NewMethodInfo(getUsersHandler, newAuthServiceGetUsersArgs, newAuthServiceGetUsersResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "auth",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.4.4",
		Extra:           extra,
	}
	return svcInfo
}

func loginHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*auth.AuthServiceLoginArgs)
	realResult := result.(*auth.AuthServiceLoginResult)
	success, err := handler.(auth.AuthService).Login(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newAuthServiceLoginArgs() interface{} {
	return auth.NewAuthServiceLoginArgs()
}

func newAuthServiceLoginResult() interface{} {
	return auth.NewAuthServiceLoginResult()
}

func adminLoginHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*auth.AuthServiceAdminLoginArgs)
	realResult := result.(*auth.AuthServiceAdminLoginResult)
	success, err := handler.(auth.AuthService).AdminLogin(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newAuthServiceAdminLoginArgs() interface{} {
	return auth.NewAuthServiceAdminLoginArgs()
}

func newAuthServiceAdminLoginResult() interface{} {
	return auth.NewAuthServiceAdminLoginResult()
}

func changeAdminPasswordHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*auth.AuthServiceChangeAdminPasswordArgs)
	realResult := result.(*auth.AuthServiceChangeAdminPasswordResult)
	success, err := handler.(auth.AuthService).ChangeAdminPassword(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newAuthServiceChangeAdminPasswordArgs() interface{} {
	return auth.NewAuthServiceChangeAdminPasswordArgs()
}

func newAuthServiceChangeAdminPasswordResult() interface{} {
	return auth.NewAuthServiceChangeAdminPasswordResult()
}

func uploadAvatarHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*auth.AuthServiceUploadAvatarArgs)
	realResult := result.(*auth.AuthServiceUploadAvatarResult)
	success, err := handler.(auth.AuthService).UploadAvatar(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newAuthServiceUploadAvatarArgs() interface{} {
	return auth.NewAuthServiceUploadAvatarArgs()
}

func newAuthServiceUploadAvatarResult() interface{} {
	return auth.NewAuthServiceUploadAvatarResult()
}

func getUserHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*auth.AuthServiceGetUserArgs)
	realResult := result.(*auth.AuthServiceGetUserResult)
	success, err := handler.(auth.AuthService).GetUser(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newAuthServiceGetUserArgs() interface{} {
	return auth.NewAuthServiceGetUserArgs()
}

func newAuthServiceGetUserResult() interface{} {
	return auth.NewAuthServiceGetUserResult()
}

func addUserHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*auth.AuthServiceAddUserArgs)
	realResult := result.(*auth.AuthServiceAddUserResult)
	success, err := handler.(auth.AuthService).AddUser(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newAuthServiceAddUserArgs() interface{} {
	return auth.NewAuthServiceAddUserArgs()
}

func newAuthServiceAddUserResult() interface{} {
	return auth.NewAuthServiceAddUserResult()
}

func deleteUserHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*auth.AuthServiceDeleteUserArgs)
	realResult := result.(*auth.AuthServiceDeleteUserResult)
	success, err := handler.(auth.AuthService).DeleteUser(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newAuthServiceDeleteUserArgs() interface{} {
	return auth.NewAuthServiceDeleteUserArgs()
}

func newAuthServiceDeleteUserResult() interface{} {
	return auth.NewAuthServiceDeleteUserResult()
}

func updateUserHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*auth.AuthServiceUpdateUserArgs)
	realResult := result.(*auth.AuthServiceUpdateUserResult)
	success, err := handler.(auth.AuthService).UpdateUser(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newAuthServiceUpdateUserArgs() interface{} {
	return auth.NewAuthServiceUpdateUserArgs()
}

func newAuthServiceUpdateUserResult() interface{} {
	return auth.NewAuthServiceUpdateUserResult()
}

func getUsersHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*auth.AuthServiceGetUsersArgs)
	realResult := result.(*auth.AuthServiceGetUsersResult)
	success, err := handler.(auth.AuthService).GetUsers(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newAuthServiceGetUsersArgs() interface{} {
	return auth.NewAuthServiceGetUsersArgs()
}

func newAuthServiceGetUsersResult() interface{} {
	return auth.NewAuthServiceGetUsersResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) Login(ctx context.Context, req *auth.LoginRequest) (r *auth.LoginResponse, err error) {
	var _args auth.AuthServiceLoginArgs
	_args.Req = req
	var _result auth.AuthServiceLoginResult
	if err = p.c.Call(ctx, "Login", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) AdminLogin(ctx context.Context, req *auth.AdminLoginRequest) (r *auth.AdminLoginResponse, err error) {
	var _args auth.AuthServiceAdminLoginArgs
	_args.Req = req
	var _result auth.AuthServiceAdminLoginResult
	if err = p.c.Call(ctx, "AdminLogin", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) ChangeAdminPassword(ctx context.Context, req *auth.ChangeAdminPasswordRequest) (r *auth.ChangeAdminPasswordResponse, err error) {
	var _args auth.AuthServiceChangeAdminPasswordArgs
	_args.Req = req
	var _result auth.AuthServiceChangeAdminPasswordResult
	if err = p.c.Call(ctx, "ChangeAdminPassword", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UploadAvatar(ctx context.Context, req *auth.UploadAvatarRequset) (r *auth.UploadAvatarResponse, err error) {
	var _args auth.AuthServiceUploadAvatarArgs
	_args.Req = req
	var _result auth.AuthServiceUploadAvatarResult
	if err = p.c.Call(ctx, "UploadAvatar", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetUser(ctx context.Context, req *auth.GetUserRequest) (r *auth.UserInfo, err error) {
	var _args auth.AuthServiceGetUserArgs
	_args.Req = req
	var _result auth.AuthServiceGetUserResult
	if err = p.c.Call(ctx, "GetUser", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) AddUser(ctx context.Context, req *auth.AddUserRequest) (r *auth.AddUserResponse, err error) {
	var _args auth.AuthServiceAddUserArgs
	_args.Req = req
	var _result auth.AuthServiceAddUserResult
	if err = p.c.Call(ctx, "AddUser", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) DeleteUser(ctx context.Context, req *auth.DeleteUserRequest) (r *auth.DeleteUserResponse, err error) {
	var _args auth.AuthServiceDeleteUserArgs
	_args.Req = req
	var _result auth.AuthServiceDeleteUserResult
	if err = p.c.Call(ctx, "DeleteUser", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UpdateUser(ctx context.Context, req *auth.UpdateUserRequest) (r *auth.UpdateUserResponse, err error) {
	var _args auth.AuthServiceUpdateUserArgs
	_args.Req = req
	var _result auth.AuthServiceUpdateUserResult
	if err = p.c.Call(ctx, "UpdateUser", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetUsers(ctx context.Context, req *auth.GetUsersRequest) (r *auth.GetUsersResponse, err error) {
	var _args auth.AuthServiceGetUsersArgs
	_args.Req = req
	var _result auth.AuthServiceGetUsersResult
	if err = p.c.Call(ctx, "GetUsers", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
