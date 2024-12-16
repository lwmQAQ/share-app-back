package userserver

import (
	"context"
	"user-server/internal/constant"
	"user-server/internal/ecode"
	"user-server/internal/models/adapter"
	"user-server/internal/svc"
	"user-server/internal/types"
	"user-server/utils"

	"strings"

	"github.com/sirupsen/logrus"
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

func (s *UserServer) LoginByCode(req *types.LoginCodeReq, ip string) types.Response {
	code, err := s.svcCtx.CodeCache.Get(req.Email)
	if err != nil {
		return types.Error(ecode.ErrCodeExtError)
	}
	if code != req.Code {
		return types.Error(ecode.ErrCodeError)
	}
	//检验通过
	user, err := s.svcCtx.UserDao.GetUserByEmail(req.Email)
	if err != nil {
		return types.Error(ecode.ErrUserNotFound)
	}
	// ip异常发送短信警告 解析ip
	go func(log *logrus.Logger) {
		if ok := utils.ContainString(strings.Split(user.IPInfo, ","), ip); !ok { //异常ip
			err := s.svcCtx.Emailutil.SendIPChangeEmail(req.Email)
			if err != nil {
				log.Errorf("发送邮件失败%v", err)
			}

		}
	}(s.svcCtx.Logger)
	//将token存入redis
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

func (s *UserServer) Login(req *types.LoginReq, ip string) types.Response {

	user, err := s.svcCtx.UserDao.GetUserByEmail(req.Email)
	if err != nil {
		return types.Error(ecode.ErrUserNotFound)
	}
	//验证密码
	if err = utils.Verify(user.Password, req.Password); err != nil {
		return types.Error(ecode.ErrPassWordError)
	}
	// ip异常发送短信警告 解析ip
	go func(log *logrus.Logger) {
		if ok := utils.ContainString(strings.Split(user.IPInfo, ","), ip); !ok { //异常ip
			err := s.svcCtx.Emailutil.SendIPChangeEmail(req.Email)
			if err != nil {
				log.Errorf("发送邮件失败%v", err)
			}

		}
	}(s.svcCtx.Logger)
	//将token存入redis
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

func (s *UserServer) SendCode(Email string) types.Response {
	code := utils.CreateCode()
	//异步发送邮件
	go func(emailUtils *utils.EmailUtil, logger *logrus.Logger) {
		//增加重试机制
		err := emailUtils.SendCode(Email, code)
		if err != nil {
			logger.Errorf("发送邮箱工具失败：%v", err)
		}
	}(s.svcCtx.Emailutil, s.svcCtx.Logger)

	//写入redis
	err := s.svcCtx.CodeCache.Put(code, Email)
	if err != nil {
		s.svcCtx.Logger.Errorf("redis 出错：%v", err)
		return types.Error(ecode.ErrSystemError)
	}
	return types.Success(nil)
}

func (s *UserServer) Register(req *types.RegisterReq, ip string) types.Response {
	user, err := adapter.BuildInsertUser(req, ip)
	if err != nil {
		s.svcCtx.Logger.Errorf("密码加密出错：%v", err)
		return types.Error(ecode.ErrSystemError)
	}
	err = s.svcCtx.UserDao.CreateUser(user)

	if err != nil {
		return types.Error(ecode.ErrSystemError)
	}
	newtoken, err := s.svcCtx.JWTUtil.GenerateToken(user.ID)
	if err != nil {
		s.svcCtx.Logger.Errorf("生成token失败：%v", err)
		return types.Error(ecode.ErrSystemError)
	}
	s.svcCtx.UserTokenCache.Set(user.ID, newtoken, constant.USER_TOKEN_EX)
	return types.Success(types.RegisterResp{
		ID:    user.ID,
		Token: newtoken,
	})
}

func (s *UserServer) GetUserinfo(userId uint64) types.Response {

	user, err := s.svcCtx.UserInfoCache.Get(userId)

	if err != nil {
		//加入缓存
		s.svcCtx.Logger.Infof("redis中不存在key%d", userId)
		user, err = s.svcCtx.UserInfoCache.LoadCache(userId)
		if err != nil {
			return types.Error(ecode.ErrUserNotExist)
		}
	}
	return types.Success(types.GetUserInfoResp{
		ID:         userId,
		Name:       user.Name,
		Avatar:     user.Avatar,
		Sex:        user.Sex,
		Bio:        user.Bio,
		Level:      user.Level,
		Experience: user.Experience,
	})
}

func (s *UserServer) UpdateUser(update *types.UpdateUserReq) types.Response {
	usermap := adapter.BuildUpdateUser(update)
	err := s.svcCtx.UserDao.UserUpdate(usermap, update.ID)
	if err != nil {
		return types.ErrorMsg(err.Error())
	}
	err = s.svcCtx.UserInfoCache.Delete(update.ID)
	if err != nil {
		s.svcCtx.Logger.Errorf("缓存出错 %v", err)
	}
	return types.Success(types.UpdateUserResp{})
}
