package dao

import "felix.bs.com/felix/BeStrongerInGO/Gin-BlogService/internal/model"

func (d *Dao) GetAuth(appKey, appSecret string) (model.Auth, error) {
	auth := model.Auth{
		AppKey:    appKey,
		AppSecret: appSecret,
	}

	return auth.Get(d.engine)
}
