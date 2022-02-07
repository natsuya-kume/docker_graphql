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

func (r *queryResolver) Admin(ctx context.Context, id string) (*model.Admin, error) {
	var admin model.Admin
	// 入力されたidをデコードする
	decID, err := base64.StdEncoding.DecodeString(id)
	if err != nil {
		log.Fatal(err)
	}

	// :以下の数字を取得
	searchID := string(decID[strings.Index(string(decID), ":")+1:])
	if err := r.DB.Find(&admin, searchID).Error; err != nil {
		return nil, err
	}
	admin.ID = id
	return &admin, nil
}

func (r *queryResolver) Admins(ctx context.Context, limit int, offset *int) (*model.AdminPagination, error) {
	var admins []*model.Admin
	if err := r.DB.Find(&admins).Error; err != nil {
		return nil, err
	}
	for _, admin := range admins {
		adminID := []byte("Admin:" + admin.ID) // プライマリキーを型名とセットで記述
		// エンコードする
		encAdminID := base64.StdEncoding.EncodeToString(adminID)
		admin.ID = encAdminID
	}
	adminPagination := &model.AdminPagination{
		Nodes: admins,
		PageInfo: &model.PaginationInfo{ //仮のデータ
			Page:             1,
			PaginationLength: 1,
			HasNextPage:      false,
			HasPreviousPage:  false,
			Count:            len(admins),
			TotalCount:       1,
		},
	}
	return adminPagination, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
