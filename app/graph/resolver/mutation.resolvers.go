package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"encoding/base64"
	"log"
	"strconv"
	"strings"

	"github.com/natsuya-kume/docker_graphql/app/graph/generated"
	"github.com/natsuya-kume/docker_graphql/app/graph/model"
)

func (r *mutationResolver) CreateAdmin(ctx context.Context, input model.CreateAdminInput) (*model.Admin, error) {
	var adminModel model.Admin

	// 最後のレコードを取得
	r.DB.Last(&adminModel)
	i, _ := strconv.Atoi(adminModel.ID)
	adminIDInt := i + 1
	adminIDStr := strconv.Itoa(adminIDInt)
	admin := model.Admin{
		ID:       adminIDStr,
		Name:     input.Name,
		Password: input.Password,
		Email:    input.Email,
	}
	// レコード作成
	r.DB.Create(&admin)
	adminID := []byte("Admin:" + adminIDStr) // プライマリキーを型名とセットで記述

	// エンコードする
	encAdminID := base64.StdEncoding.EncodeToString(adminID)
	admin.ID = encAdminID
	return &admin, nil
}

func (r *mutationResolver) UpdateAdmin(ctx context.Context, input model.UpdateAdminInput) (*model.Admin, error) {
	var admin model.Admin

	// 入力されたidをデコードする
	decID, err := base64.StdEncoding.DecodeString(input.ID)
	if err != nil {
		log.Fatal(err)
	}

	// :以下の数字を取得
	searchID := string(decID[strings.Index(string(decID), ":")+1:])

	r.DB.First(&admin, searchID)
	admin.Name = *input.Name
	admin.Email = *input.Email
	admin.Password = *input.Password
	r.DB.Save(&admin)

	admin.ID = input.ID

	return &admin, nil
}

func (r *mutationResolver) DeleteAdmin(ctx context.Context, id string) (bool, error) {
	var admin model.Admin

	// 入力されたidをデコードする
	decID, err := base64.StdEncoding.DecodeString(id)
	if err != nil {
		log.Fatal(err)
	}

	// :以下の数字を取得
	searchID := string(decID[strings.Index(string(decID), ":")+1:])

	r.DB.Where("id = ?", searchID).Delete(&admin)
	return true, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
