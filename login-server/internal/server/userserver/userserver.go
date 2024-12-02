package userserver

import (
	"context"
	"login-server/internal/constant"
	"login-server/internal/ecode"
	"login-server/internal/svc"
	"login-server/internal/types"
	"login-server/utils"
)

type UserServer struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserServer(ctx context.Context, svcCtx *svc.ServiceContext) *UserServer {
	return &UserServer{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (s *UserServer) Login(req *types.LoginReq, ip string) types.Response {
	user, err := s.svcCtx.UserDao.GetUserByEmail(req.Email)
	if err != nil {
		return types.Error(ecode.ErrUserNotFound)
	}
	//验证密码
	if err = utils.Verify(user.Password, req.Password); err != nil {
		return types.Error(ecode.ErrPassWordError)
	}
	// TODO 将token存入redis
	token, err := s.svcCtx.UserTokenCache.Get(user.ID) //当前有没有token
	if err != nil {
		newtoken, err := s.svcCtx.JWTUtil.GenerateToken(user.ID)
		if err != nil {
			s.svcCtx.Logger.Errorf("生成token失败：%v", err)
			return types.Error(ecode.ErrSystemError)
		}
		s.svcCtx.UserTokenCache.Set(user.ID, newtoken, constant.USER_TOKEN_EX)
		token = newtoken
	}
	return types.Success(&types.LoginResp{
		ID:    user.ID,
		Token: token,
	})
}
