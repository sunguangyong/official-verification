package model

import (
	"fmt"

	"context"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ReportRecordModel = (*customReportRecordModel)(nil)

type (
	// ReportRecordModel is an interface to be customized, add more methods here,
	// and implement the added methods in customReportRecordModel.
	ReportRecordModel interface {
		reportRecordModel
		FindNewsCount(ctx context.Context, paramSql string) (count int64, err error)
		CommonFind(ctx context.Context, querySql, orderSql, limitSql string) ([]*OfficialVerify, error)
	}

	customReportRecordModel struct {
		*defaultReportRecordModel
	}
)

// NewReportRecordModel returns a model for the database table.
func NewReportRecordModel(conn sqlx.SqlConn) ReportRecordModel {
	return &customReportRecordModel{
		defaultReportRecordModel: newReportRecordModel(conn),
	}
}

func (m *defaultReportRecordModel) CommonFind(ctx context.Context, querySql, orderSql, limitSql string) ([]*OfficialVerify, error) {
	query := fmt.Sprintf("select %s from %s %s %s %s", reportRecordRows, m.table, querySql, orderSql, limitSql)
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

func (m *defaultReportRecordModel) FindNewsCount(ctx context.Context, paramSql string) (count int64, err error) {
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
