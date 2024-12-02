package utils

import (
	"bytes"
	"context"
	"fmt"
	"oss-server/config"

	"time"

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

// 定义常量映射
var sceneMap = map[int32]string{
	1: "user-avatar",
	2: "analyse-file",
	3: "task-result",
}

// 上传文件并返回一个文件下载地址
func (client *MinioClient) UploadFile(flie []byte, uuid string, ext string, bucketName string, scene int32) (string, error) {
	// 上传文件到 MinIO
	absolutePath, err := GenerateFileName(uuid, ext, scene)
	if err != nil {
		return "", err
	}
	// 将字节数组转换为 *bytes.Reader
	byteReader := bytes.NewReader(flie)
	// 使用 PutObject 上传文件
	_, err = client.Client.PutObject(context.Background(), bucketName, absolutePath, byteReader, int64(len(flie)), minio.PutObjectOptions{})
	if err != nil {
		return "", err
	}

	// 生成文件的下载链接
	DownLoadURL := getDownloadUrl(client.config.Endpoint, bucketName, absolutePath)

	return DownLoadURL, nil // 返回下载链接
}

// 获取一个带签名的put连接上传 URL 和下载 URL
func (minioClient *MinioClient) GetUploadUrl(bucketName string, uuid string, Scene int32, ext string) (string, string, error) {
	fmt.Println(bucketName)
	absolutePath, err := GenerateFileName(uuid, ext, Scene)
	if err != nil {
		return "", "", err
	}
	// 设置有效期为1天
	expires := time.Hour
	// 生成 PUT 请求的预签名 URL
	preSignedUrl, err := minioClient.Client.PresignedPutObject(context.Background(), bucketName, absolutePath, expires)
	if err != nil {
		return "", "", fmt.Errorf("failed to get presigned URL: %w", err)
	}

	downloadUrl := getDownloadUrl(minioClient.config.Endpoint, bucketName, absolutePath)
	return preSignedUrl.String(), downloadUrl, nil
}

// GenerateRandomFileName 生成带有UUID的随机文件名
func GenerateFileName(uuid string, ext string, Scene int32) (string, error) {
	value, exists := sceneMap[Scene]
	if !exists {
		return "", fmt.Errorf("场景有误")
	}
	// 拼接成文件名，带有文件后缀
	return fmt.Sprintf("%s/%s/%s%s", value, GetCurrentYearMonth(), uuid, ext), nil
}
func GetCurrentYearMonth() string {
	now := time.Now()
	return fmt.Sprintf("%d-%02d", now.Year(), now.Month())
}

// 获取下载 URL 的示例函数（替换为实际的下载 URL 生成逻辑）
func getDownloadUrl(endpoint string, bucketName string, filename string) string {
	return fmt.Sprintf("http://%s/%s/%s", endpoint, bucketName, filename)
}
