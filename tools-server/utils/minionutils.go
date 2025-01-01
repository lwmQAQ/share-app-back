package utils

import (
	"context"
	"fmt"
	"time"
	"tools-back/config"
	"tools-back/internal/types"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type MinioClient struct {
	Client *minio.Client
	config *config.MinioConfig
}

func NewMinioClient(config *config.MinioConfig) *MinioClient {
	/*
		Secure: true:
		启用 HTTPS（SSL/TLS）加密传输。这意味着所有的数据传输都是加密的，适合在需要保护数据传输的场景下使用，比如生产环境。
		Secure: false:
		使用 HTTP 进行非加密的通信。这适合在开发环境或内部网络中使用，尤其是在不涉及敏感数据的情况下。
	*/
	// 根据实际情况选择是否使用 SSL
	useSSL := false // 在开发环境下通常设为 false
	minioClient, err := minio.New(config.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(config.AccessKeyID, config.SecretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {

		return nil
	}
	client := &MinioClient{
		Client: minioClient,
		config: config,
	}
	return client
}

// ext 文件后缀
func (c *MinioClient) GetUploadUrl(ext string, bucketName string, uuid string) (*types.OssUploadResp, error) {
	absolutePath, err := GenerateFileName(uuid, ext)
	if err != nil {
		return nil, err
	}
	// 设置有效期为1天
	expires := time.Hour * 24
	// 生成 PUT 请求的预签名 URL
	preSignedUrl, err := c.Client.PresignedPutObject(context.Background(), bucketName, absolutePath, expires)
	if err != nil {
		return nil, fmt.Errorf("failed to get presigned URL: %w", err)
	}

	downloadUrl := getDownloadUrl(c.config.Endpoint, bucketName, absolutePath)

	return &types.OssUploadResp{
		UploadUrl:   preSignedUrl.String(),
		DownLoadUrl: downloadUrl,
	}, nil
}

func GenerateFileName(uuid string, ext string) (string, error) {
	return fmt.Sprintf("%s/%s.%s", GetCurrentYearMonth(), uuid, ext), nil
}
func GetCurrentYearMonth() string {
	now := time.Now()
	return fmt.Sprintf("%d-%02d-%02d", now.Year(), now.Month(), now.Day())
}
func getDownloadUrl(endpoint string, bucketName string, absolutePath string) string {
	return fmt.Sprintf("http://%s/%s/%s", endpoint, bucketName, absolutePath)
}
