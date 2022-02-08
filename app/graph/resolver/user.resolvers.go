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

func (r *userResolver) Service(ctx context.Context, obj *model.User) (*model.Service, error) {
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

func (r *userResolver) PersonalTags(ctx context.Context, obj *model.User) (*model.PersonalTagPagination, error) {
	// あるユーザーが作成したパーソナルタグを取得する関数
	var personalTagPagination model.PersonalTagPagination

	decID, err := base64.StdEncoding.DecodeString(obj.ID)
	if err != nil {
		log.Fatal(err)
	}
	// :以下の数字を取得
	searchID := string(decID[strings.Index(string(decID), ":")+1:])
	if err := r.DB.Where("user_id=?", searchID).Find(&personalTagPagination.Nodes).Error; err != nil {
		return nil, err
	}
	for _, personalTag := range personalTagPagination.Nodes {
		personalTagID := []byte("PersonalTag:" + personalTag.ID) // プライマリキーを型名とセットで記述
		// エンコードする
		encPersonalTagID := base64.StdEncoding.EncodeToString(personalTagID)
		personalTag.ID = encPersonalTagID
	}

	return &personalTagPagination, nil
}

func (r *userResolver) ReviewTags(ctx context.Context, obj *model.User) (*model.ReviewTagPagination, error) {
	// あるユーザーが作成したレビュータグを取得する関数
	var reviewTagPagination model.ReviewTagPagination

	decID, err := base64.StdEncoding.DecodeString(obj.ID)
	if err != nil {
		log.Fatal(err)
	}
	// :以下の数字を取得
	searchID := string(decID[strings.Index(string(decID), ":")+1:])

	if err := r.DB.Where("user_id=?", searchID).Find(&reviewTagPagination.Nodes).Error; err != nil {
		return nil, err
	}
	for _, reviewTag := range reviewTagPagination.Nodes {
		reviewTagID := []byte("ReviewTag:" + reviewTag.ID) // プライマリキーを型名とセットで記述
		// エンコードする
		encReviewTagID := base64.StdEncoding.EncodeToString(reviewTagID)
		reviewTag.ID = encReviewTagID
	}
	return &reviewTagPagination, nil
}

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type userResolver struct{ *Resolver }
