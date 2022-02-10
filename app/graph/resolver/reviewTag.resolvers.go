package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/natsuya-kume/docker_graphql/app/graph/generated"
	"github.com/natsuya-kume/docker_graphql/app/graph/model"
	"github.com/natsuya-kume/docker_graphql/app/utils"
)

func (r *reviewTagResolver) User(ctx context.Context, obj *model.ReviewTag) (*model.User, error) {
	// 1つのレビュータグを作成したユーザーを取得する関数
	var user model.User
	if err := r.DB.Where("id=?", obj.UserID).Find(&user).Error; err != nil { //一旦1人目のユーザーだけをとってくる
		return nil, err
	}
	user.ID = utils.Encode("User:", user.ID)
	return &user, nil
}

// ReviewTag returns generated.ReviewTagResolver implementation.
func (r *Resolver) ReviewTag() generated.ReviewTagResolver { return &reviewTagResolver{r} }

type reviewTagResolver struct{ *Resolver }
