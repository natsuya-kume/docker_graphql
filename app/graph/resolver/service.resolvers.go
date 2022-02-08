package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"encoding/base64"
	"log"
	"strings"

	"github.com/natsuya-kume/docker_graphql/app/graph/generated"
	"github.com/natsuya-kume/docker_graphql/app/graph/model"
)

func (r *serviceResolver) ServiceAccounts(ctx context.Context, obj *model.Service) (*model.ServiceAccountPagination, error) {
	var serviceAccountPagination model.ServiceAccountPagination

	decID, err := base64.StdEncoding.DecodeString(obj.ID)
	if err != nil {
		log.Fatal(err)
	}
	// :以下の数字を取得
	searchID := string(decID[strings.Index(string(decID), ":")+1:])
	// あるサービス(楽天を想定)から作成されたサービスアカウントを取得する関数
	if err := r.DB.Where("service_id=?", searchID).Find(&serviceAccountPagination.Nodes).Error; err != nil {
		return nil, err
	}
	for _, serviceAccount := range serviceAccountPagination.Nodes {
		serviceAccountID := []byte("ServiceAccount:" + serviceAccount.ID) // プライマリキーを型名とセットで記述
		// エンコードする
		encServiceAccountID := base64.StdEncoding.EncodeToString(serviceAccountID)
		serviceAccount.ID = encServiceAccountID
	}
	return &serviceAccountPagination, nil
}

func (r *serviceResolver) Users(ctx context.Context, obj *model.Service) (*model.UserPagination, error) {
	var userPagination model.UserPagination

	decID, err := base64.StdEncoding.DecodeString(obj.ID)
	if err != nil {
		log.Fatal(err)
	}
	// :以下の数字を取得
	searchID := string(decID[strings.Index(string(decID), ":")+1:])
	// あるサービス(楽天を想定)から作成されたユーザーを取得する関数
	if err := r.DB.Where("service_id=?", searchID).Find(&userPagination.Nodes).Error; err != nil {
		return nil, err
	}
	for _, user := range userPagination.Nodes {
		userID := []byte("User:" + user.ID) // プライマリキーを型名とセットで記述
		// エンコードする
		encUserID := base64.StdEncoding.EncodeToString(userID)
		user.ID = encUserID
	}
	return &userPagination, nil
}

// Service returns generated.ServiceResolver implementation.
func (r *Resolver) Service() generated.ServiceResolver { return &serviceResolver{r} }

type serviceResolver struct{ *Resolver }
