package model

import (
	"fmt"

	"context"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ SocialEnumModel = (*customSocialEnumModel)(nil)

type (
	// SocialEnumModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSocialEnumModel.
	SocialEnumModel interface {
		socialEnumModel
		CommonFind(ctx context.Context, querySql, orderSql, limitSql string) ([]*SocialEnum, error)
	}

	customSocialEnumModel struct {
		*defaultSocialEnumModel
	}
)

// NewSocialEnumModel returns a model for the database table.
func NewSocialEnumModel(conn sqlx.SqlConn) SocialEnumModel {
	return &customSocialEnumModel{
		defaultSocialEnumModel: newSocialEnumModel(conn),
	}
}

func (m *defaultSocialEnumModel) CommonFind(ctx context.Context, querySql, orderSql, limitSql string) ([]*SocialEnum, error) {
	query := fmt.Sprintf("select %s from %s %s %s %s", socialEnumRows, m.table, querySql, orderSql, limitSql)
	var resp []*SocialEnum
	err := m.conn.QueryRowsCtx(ctx, &resp, query)
	switch err {
	case nil:
		return resp, nil
	case sqlx.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
