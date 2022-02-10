package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/natsuya-kume/docker_graphql/app/graph/generated"
	"github.com/natsuya-kume/docker_graphql/app/graph/model"
	"github.com/natsuya-kume/docker_graphql/app/utils"
)

func (r *personalTagResolver) User(ctx context.Context, obj *model.PersonalTag) (*model.User, error) {
	// 1つのパーソナルタグを作成したユーザーを取得する関数
	var user model.User
	if err := r.DB.Where("id=?", obj.UserID).Find(&user).Error; err != nil { //一旦1人目のユーザーだけをとってくる
		return nil, err
	}
	user.ID = utils.Encode("User:", user.ID)
	return &user, nil
}

// PersonalTag returns generated.PersonalTagResolver implementation.
func (r *Resolver) PersonalTag() generated.PersonalTagResolver { return &personalTagResolver{r} }

type personalTagResolver struct{ *Resolver }
