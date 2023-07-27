package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var ErrNotFound = sqlx.ErrNotFound

// MultipartUploadInfo : 初始化信息
type MultipartUploadInfo struct {
	FileSha1   string // 文件Sha1
	FileSize   int64  // 文件大小
	UploadID   string // 上传id
	ChunkSize  int    // 分块大小
	ChunkCount int    // 分块数量
}
