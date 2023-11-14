package user

import (
	"FuguBackend/app/repository/mysql/users"
	"go.uber.org/zap"
)

func (s *service) FindAll() ([]*users.Users, error) {
	qb := users.NewQueryBuilder()
	all, err := qb.QueryAll(s.db.GetDbR())
	if err != nil {
		s.logger.Error("find all failed ", zap.Error(err))
		return nil, err
	}
	return all, nil

}
