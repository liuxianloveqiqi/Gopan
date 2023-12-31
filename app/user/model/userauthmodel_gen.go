// Code generated by goctl. DO NOT EDIT.

package model

import (
	"context"
	"database/sql"
	"fmt"
	"gorm.io/gorm"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	userAuthFieldNames          = builder.RawFieldNames(&UserAuth{})
	userAuthRows                = strings.Join(userAuthFieldNames, ",")
	userAuthRowsExpectAutoSet   = strings.Join(stringx.Remove(userAuthFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	userAuthRowsWithPlaceHolder = strings.Join(stringx.Remove(userAuthFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheUserAuthIdPrefix = "cache:userAuth:id:"
)

type (
	userAuthModel interface {
		Insert(ctx context.Context, data *UserAuth) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*UserAuth, error)
		Update(ctx context.Context, data *UserAuth) error
		Delete(ctx context.Context, id int64) error
	}

	defaultUserAuthModel struct {
		sqlc.CachedConn
		table string
	}

	UserAuth struct {
		Id         int64          `db:"id" gorm:"column:id"`
		CreateTime time.Time      `gorm:"column:create_time"` // 创建时间
		UpdateTime time.Time      `gorm:"column:update_time"` // 更新时间
		DeleteTime gorm.DeletedAt `gorm:"index " db:"delete_time"`
		UserId     int64          `db:"user_id" gorm:"column:user_id"`
		ProviderId string         `db:"provider_id" gorm:"column:provider_id"` // 平台唯一id
		Provider   string         `db:"provider" gorm:"column:provider"`       // 平台名称
	}
)

func newUserAuthModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultUserAuthModel {
	return &defaultUserAuthModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`user_auth`",
	}
}

func (m *defaultUserAuthModel) Delete(ctx context.Context, id int64) error {
	userAuthIdKey := fmt.Sprintf("%s%v", cacheUserAuthIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, userAuthIdKey)
	return err
}

func (m *defaultUserAuthModel) FindOne(ctx context.Context, id int64) (*UserAuth, error) {
	userAuthIdKey := fmt.Sprintf("%s%v", cacheUserAuthIdPrefix, id)
	var resp UserAuth
	err := m.QueryRowCtx(ctx, &resp, userAuthIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", userAuthRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUserAuthModel) Insert(ctx context.Context, data *UserAuth) (sql.Result, error) {
	userAuthIdKey := fmt.Sprintf("%s%v", cacheUserAuthIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?)", m.table, userAuthRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.DeleteTime, data.UserId, data.ProviderId, data.Provider)
	}, userAuthIdKey)
	return ret, err
}

func (m *defaultUserAuthModel) Update(ctx context.Context, data *UserAuth) error {
	userAuthIdKey := fmt.Sprintf("%s%v", cacheUserAuthIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, userAuthRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.DeleteTime, data.UserId, data.ProviderId, data.Provider, data.Id)
	}, userAuthIdKey)
	return err
}

func (m *defaultUserAuthModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheUserAuthIdPrefix, primary)
}

func (m *defaultUserAuthModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", userAuthRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultUserAuthModel) tableName() string {
	return m.table
}
