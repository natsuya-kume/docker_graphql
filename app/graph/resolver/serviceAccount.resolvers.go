package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"encoding/base64"

	"github.com/natsuya-kume/docker_graphql/app/graph/generated"
	"github.com/natsuya-kume/docker_graphql/app/graph/model"
)

func (r *serviceAccountResolver) Service(ctx context.Context, obj *model.ServiceAccount) (*model.Service, error) {
	// サービスアカウントを作成したサービスを取得する関数
	var service model.Service
	if err := r.DB.Where("id=?", obj.ServiceID).Find(&service).Error; err != nil { //一旦楽天だけをとってくる
		return nil, err
	}
	serviceID := []byte("Service:" + service.ID) // プライマリキーを型名とセットで記述
	// エンコードする
	encServiceID := base64.StdEncoding.EncodeToString(serviceID)
	service.ID = encServiceID
	return &service, nil
}

// ServiceAccount returns generated.ServiceAccountResolver implementation.
func (r *Resolver) ServiceAccount() generated.ServiceAccountResolver {
	return &serviceAccountResolver{r}
}

type serviceAccountResolver struct{ *Resolver }
