// Code generated by mockery v2.22.1. DO NOT EDIT.

package mocks

import (
	context "context"

	pgconn "github.com/jackc/pgconn"
	mock "github.com/stretchr/testify/mock"

	pgx "github.com/jackc/pgx/v4"
)

// DBTxMock is an autogenerated mock type for the Tx type
type DBTxMock struct {
	mock.Mock
}

// Begin provides a mock function with given fields: ctx
func (_m *DBTxMock) Begin(ctx context.Context) (pgx.Tx, error) {
	ret := _m.Called(ctx)

	var r0 pgx.Tx
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (pgx.Tx, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) pgx.Tx); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(pgx.Tx)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BeginFunc provides a mock function with given fields: ctx, f
func (_m *DBTxMock) BeginFunc(ctx context.Context, f func(pgx.Tx) error) error {
	ret := _m.Called(ctx, f)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, func(pgx.Tx) error) error); ok {
		r0 = rf(ctx, f)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Commit provides a mock function with given fields: ctx
func (_m *DBTxMock) Commit(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Conn provides a mock function with given fields:
func (_m *DBTxMock) Conn() *pgx.Conn {
	ret := _m.Called()

	var r0 *pgx.Conn
	if rf, ok := ret.Get(0).(func() *pgx.Conn); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pgx.Conn)
		}
	}

	return r0
}

// CopyFrom provides a mock function with given fields: ctx, tableName, columnNames, rowSrc
func (_m *DBTxMock) CopyFrom(ctx context.Context, tableName pgx.Identifier, columnNames []string, rowSrc pgx.CopyFromSource) (int64, error) {
	ret := _m.Called(ctx, tableName, columnNames, rowSrc)

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, pgx.Identifier, []string, pgx.CopyFromSource) (int64, error)); ok {
		return rf(ctx, tableName, columnNames, rowSrc)
	}
	if rf, ok := ret.Get(0).(func(context.Context, pgx.Identifier, []string, pgx.CopyFromSource) int64); ok {
		r0 = rf(ctx, tableName, columnNames, rowSrc)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, pgx.Identifier, []string, pgx.CopyFromSource) error); ok {
		r1 = rf(ctx, tableName, columnNames, rowSrc)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Exec provides a mock function with given fields: ctx, sql, arguments
func (_m *DBTxMock) Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error) {
	var _ca []interface{}
	_ca = append(_ca, ctx, sql)
	_ca = append(_ca, arguments...)
	ret := _m.Called(_ca...)

	var r0 pgconn.CommandTag
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, ...interface{}) (pgconn.CommandTag, error)); ok {
		return rf(ctx, sql, arguments...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, ...interface{}) pgconn.CommandTag); ok {
		r0 = rf(ctx, sql, arguments...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(pgconn.CommandTag)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, ...interface{}) error); ok {
		r1 = rf(ctx, sql, arguments...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LargeObjects provides a mock function with given fields:
func (_m *DBTxMock) LargeObjects() pgx.LargeObjects {
	ret := _m.Called()

	var r0 pgx.LargeObjects
	if rf, ok := ret.Get(0).(func() pgx.LargeObjects); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(pgx.LargeObjects)
	}

	return r0
}

// Prepare provides a mock function with given fields: ctx, name, sql
func (_m *DBTxMock) Prepare(ctx context.Context, name string, sql string) (*pgconn.StatementDescription, error) {
	ret := _m.Called(ctx, name, sql)

	var r0 *pgconn.StatementDescription
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) (*pgconn.StatementDescription, error)); ok {
		return rf(ctx, name, sql)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *pgconn.StatementDescription); ok {
		r0 = rf(ctx, name, sql)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pgconn.StatementDescription)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, name, sql)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Query provides a mock function with given fields: ctx, sql, args
func (_m *DBTxMock) Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error) {
	var _ca []interface{}
	_ca = append(_ca, ctx, sql)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	var r0 pgx.Rows
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, ...interface{}) (pgx.Rows, error)); ok {
		return rf(ctx, sql, args...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, ...interface{}) pgx.Rows); ok {
		r0 = rf(ctx, sql, args...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(pgx.Rows)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, ...interface{}) error); ok {
		r1 = rf(ctx, sql, args...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// QueryFunc provides a mock function with given fields: ctx, sql, args, scans, f
func (_m *DBTxMock) QueryFunc(ctx context.Context, sql string, args []interface{}, scans []interface{}, f func(pgx.QueryFuncRow) error) (pgconn.CommandTag, error) {
	ret := _m.Called(ctx, sql, args, scans, f)

	var r0 pgconn.CommandTag
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, []interface{}, []interface{}, func(pgx.QueryFuncRow) error) (pgconn.CommandTag, error)); ok {
		return rf(ctx, sql, args, scans, f)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, []interface{}, []interface{}, func(pgx.QueryFuncRow) error) pgconn.CommandTag); ok {
		r0 = rf(ctx, sql, args, scans, f)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(pgconn.CommandTag)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, []interface{}, []interface{}, func(pgx.QueryFuncRow) error) error); ok {
		r1 = rf(ctx, sql, args, scans, f)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// QueryRow provides a mock function with given fields: ctx, sql, args
func (_m *DBTxMock) QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row {
	var _ca []interface{}
	_ca = append(_ca, ctx, sql)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	var r0 pgx.Row
	if rf, ok := ret.Get(0).(func(context.Context, string, ...interface{}) pgx.Row); ok {
		r0 = rf(ctx, sql, args...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(pgx.Row)
		}
	}

	return r0
}

// Rollback provides a mock function with given fields: ctx
func (_m *DBTxMock) Rollback(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SendBatch provides a mock function with given fields: ctx, b
func (_m *DBTxMock) SendBatch(ctx context.Context, b *pgx.Batch) pgx.BatchResults {
	ret := _m.Called(ctx, b)

	var r0 pgx.BatchResults
	if rf, ok := ret.Get(0).(func(context.Context, *pgx.Batch) pgx.BatchResults); ok {
		r0 = rf(ctx, b)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(pgx.BatchResults)
		}
	}

	return r0
}

type mockConstructorTestingTNewDBTxMock interface {
	mock.TestingT
	Cleanup(func())
}

// NewDBTxMock creates a new instance of DBTxMock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewDBTxMock(t mockConstructorTestingTNewDBTxMock) *DBTxMock {
	mock := &DBTxMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
