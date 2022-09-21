package model

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlc"

	"context"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ OfficialVerifyModel = (*customOfficialVerifyModel)(nil)

type (
	// OfficialVerifyModel is an interface to be customized, add more methods here,
	// and implement the added methods in customOfficialVerifyModel.
	OfficialVerifyModel interface {
		officialVerifyModel
		CommonFind(ctx context.Context, querySql, orderSql, limitSql string) ([]*OfficialVerify, error)
		FindNewsCount(ctx context.Context, paramSql string) (count int64, err error)
	}

	customOfficialVerifyModel struct {
		*defaultOfficialVerifyModel
	}
)

// NewOfficialVerifyModel returns a model for the database table.
func NewOfficialVerifyModel(conn sqlx.SqlConn) OfficialVerifyModel {
	return &customOfficialVerifyModel{
		defaultOfficialVerifyModel: newOfficialVerifyModel(conn),
	}
}

func (m *defaultOfficialVerifyModel) CommonFind(ctx context.Context, querySql, orderSql, limitSql string) ([]*OfficialVerify, error) {
	query := fmt.Sprintf("select %s from %s %s %s %s", officialVerifyRows, m.table, querySql, orderSql, limitSql)
	var resp []*OfficialVerify
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


func (m *defaultOfficialVerifyModel) FindNewsCount(ctx context.Context, paramSql string) (count int64, err error) {
	query := fmt.Sprintf(`select count(1) from %s %s`, m.table, paramSql)
	err = m.conn.QueryRowCtx(ctx, &count, query)

	switch err {
	case nil:
		return count, nil
	case sqlc.ErrNotFound:
		return 0, ErrNotFound
	default:
		return 0, err
	}
}