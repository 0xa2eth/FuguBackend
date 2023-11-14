package secret

import (
	"FuguBackend/app/pkg/core"
	"FuguBackend/app/repository/mysql"
	"FuguBackend/app/repository/mysql/secrets"
)

func (s *service) GetPostsByAuthorID(c core.Context, id int) ([]*secrets.Secrets, error) {
	sqb := secrets.NewQueryBuilder()
	sqb.WhereAuthorId(mysql.EqualPredicate, int64(id))
	list, err := sqb.QueryAll(s.db.GetDbR().WithContext(c.RequestContext()))
	//count, err := sqb.Count(s.db.GetDbR().WithContext(c.RequestContext()))
	if err != nil {
		return nil, err
	}

	return list, err
}
