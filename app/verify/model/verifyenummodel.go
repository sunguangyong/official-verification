package model

import (
	"fmt"

	"context"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ VerifyEnumModel = (*customVerifyEnumModel)(nil)

type (
	// VerifyEnumModel is an interface to be customized, add more methods here,
	// and implement the added methods in customVerifyEnumModel.
	VerifyEnumModel interface {
		verifyEnumModel
		CommonFind(ctx context.Context, querySql, orderSql, limitSql string) ([]*VerifyEnum, error)
	}

	customVerifyEnumModel struct {
		*defaultVerifyEnumModel
	}
)

// NewVerifyEnumModel returns a model for the database table.
func NewVerifyEnumModel(conn sqlx.SqlConn) VerifyEnumModel {
	return &customVerifyEnumModel{
		defaultVerifyEnumModel: newVerifyEnumModel(conn),
	}
}

func (m *defaultVerifyEnumModel) CommonFind(ctx context.Context, querySql, orderSql, limitSql string) ([]*VerifyEnum, error) {
	query := fmt.Sprintf("select %s from %s %s %s %s", verifyEnumRows, m.table, querySql, orderSql, limitSql)
	var resp []*VerifyEnum
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
