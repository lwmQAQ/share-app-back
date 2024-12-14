package postserver

import (
	"context"
	"resource-server/internal/constants"
	"resource-server/internal/ecode"
	"resource-server/internal/models"
	"resource-server/internal/models/adapter"
	"resource-server/internal/rpcclient/userclient"
	"resource-server/internal/svc"
	"resource-server/internal/types"
	"resource-server/utils"
	"time"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PostServer struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPostServer(ctx context.Context, svcCtx *svc.ServiceContext) *PostServer {
	return &PostServer{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (s *PostServer) CreatePost(req *types.CreatePostReq) types.Response {

	//1.获取rpc作者信息
	addr, err := s.svcCtx.EtcdUtil.GetServiceInstance("UserServer")
	if err != nil {
		s.svcCtx.Logger.Errorf("获取rpc服务失败 %v", err)
		return types.Error(ecode.ErrSystemError)
	}
	rpcreq := &userclient.GetUserInfoReq{
		Id: req.UserID,
	}

	resp, err := s.svcCtx.UserRpcClient.GetUserInfo(context.Background(), rpcreq, addr)
	if err != nil {
		s.svcCtx.Logger.Errorf("rpc服务出错 %v", err)
		return types.Error(ecode.ErrSystemError)
	}
	author := &models.User{
		Username: resp.Username,
		Avatar:   resp.Avatar,
		ID:       resp.Id,
	}

	//2.TODO异步增加用户经验
	//3.插入mongo文档
	link, err := s.svcCtx.UrlUtil.CreateShortLink(req.Link)
	if err != nil {
		return types.Error(ecode.ErrSystemError)
	}
	req.Link = link
	post := adapter.BuildInsertPost(author, req)
	postid, err := s.svcCtx.MongoUtil.InsertDocument("Post", post)
	if err != nil {
		s.svcCtx.Logger.Errorf("mongodb出错 %v", err)
		return types.Error(ecode.ErrSystemError)
	}
	// 检查 postid 类型并转换为字符串
	postidStr, ok := postid.(primitive.ObjectID)
	if !ok {
		return types.Error(ecode.ErrSystemError)
	}

	postidAsString := postidStr.Hex()
	//4. 异步更新es
	go func(id string, client *utils.ESClient, logger *logrus.Logger) {
		resource := &models.Resource{
			Title:       post.Title,
			Tags:        post.Tags,
			LikeCount:   0,
			Publisher:   author.Username,
			PublishTime: time.Now(),
		}
		err := client.InsertDocument(constants.ResourceIndex, resource, id)
		if err != nil {
			logger.Errorf("es插入出错 %v", err)
		}
	}(postidAsString, s.svcCtx.ESClient, s.svcCtx.Logger)
	return types.Success(&types.CreatePostResp{})
}

func (s *PostServer) GetPostById(req string) types.Response {
	//1.TODO 异步更新es
	// 将 string 转换为 primitive.ObjectID
	objectID, err := primitive.ObjectIDFromHex(req)
	if err != nil {
		s.svcCtx.Logger.Errorf("转换失败: %v", err)
	}
	var post models.Post
	err = s.svcCtx.MongoUtil.SearchDocumentByID("Post", objectID, &post)
	if err != nil {
		s.svcCtx.Logger.Errorf("查询mongodb失败 %v", err)
		return types.Error(ecode.ErrSystemError)
	}

	return types.Success(post)
}

func (s *PostServer) LikePost(postId string) types.Response {
	// 修改mongo数据库
	objectID, err := primitive.ObjectIDFromHex(postId)
	if err != nil {
		s.svcCtx.Logger.Errorf("转换失败: %v", err)
	}
	// 定义更新条件和更新内容
	filter := bson.M{"_id": objectID}
	update := bson.M{
		"$inc": bson.M{"likes": 1}, // 示例：将 likes 字段加 1
	}

	// 执行更新操作
	_, err = s.svcCtx.MongoUtil.UpdateDocument("Post", filter, update)
	if err != nil {
		s.svcCtx.Logger.Errorf("更新失败: %v", err)
		return types.Error(ecode.ErrSystemError)
	}

	//1.TODO 异步更新es
	return types.Success(nil)
}
