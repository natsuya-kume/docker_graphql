package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"strings"

	"github.com/natsuya-kume/docker_graphql/app/graph/generated"
	"github.com/natsuya-kume/docker_graphql/app/graph/model"
	"github.com/natsuya-kume/docker_graphql/app/utils"
)

func (r *userResolver) Service(ctx context.Context, obj *model.User) (*model.Service, error) {
	var service model.Service
	if err := r.DB.Where("id=?", obj.ServiceID).Find(&service).Error; err != nil { //一旦楽天だけをとってくる
		return nil, err
	}
	service.ID = utils.Encode("Service:", service.ID)
	return &service, nil
}

func (r *userResolver) PersonalTags(ctx context.Context, obj *model.User) (*model.PersonalTagPagination, error) {
	// あるユーザーが作成したパーソナルタグを取得する関数
	var personalTagPagination model.PersonalTagPagination

	decID := utils.Decode(obj.ID)
	// :以下の数字を取得
	searchID := string(decID[strings.Index(string(decID), ":")+1:])
	if err := r.DB.Where("user_id=?", searchID).Find(&personalTagPagination.Nodes).Error; err != nil {
		return nil, err
	}
	for _, personalTag := range personalTagPagination.Nodes {
		personalTag.ID = utils.Encode("PersonalTag:", personalTag.ID)
	}

	return &personalTagPagination, nil
}

func (r *userResolver) ReviewTags(ctx context.Context, obj *model.User) (*model.ReviewTagPagination, error) {
	// あるユーザーが作成したレビュータグを取得する関数
	var reviewTagPagination model.ReviewTagPagination

	decID := utils.Decode(obj.ID)
	// :以下の数字を取得
	searchID := string(decID[strings.Index(string(decID), ":")+1:])

	if err := r.DB.Where("user_id=?", searchID).Find(&reviewTagPagination.Nodes).Error; err != nil {
		return nil, err
	}
	for _, reviewTag := range reviewTagPagination.Nodes {
		reviewTag.ID = utils.Encode("ReviewTag:", reviewTag.ID)

	}
	return &reviewTagPagination, nil
}

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type userResolver struct{ *Resolver }
