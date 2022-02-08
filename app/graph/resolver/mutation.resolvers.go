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

var personalTags []*model.PersonalTag = make([]*model.PersonalTag, 0)
var reviewTags []*model.ReviewTag = make([]*model.ReviewTag, 0)

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

func (r *mutationResolver) CreateService(ctx context.Context, input model.CreateServiceInput) (*model.Service, error) {
	var serviceModel model.Service
	// 最後のレコードを取得
	r.DB.Last(&serviceModel)
	i, _ := strconv.Atoi(serviceModel.ID)
	serviceIDInt := i + 1
	serviceIDStr := strconv.Itoa(serviceIDInt)

	service := model.Service{
		ID:       serviceIDStr,
		Name:     input.Name,
		Password: input.Password,
		Email:    input.Email,
	}
	r.DB.Create(&service)
	serviceID := []byte("Service:" + serviceIDStr) // プライマリキーを型名とセットで記述

	// エンコードする
	encServiceID := base64.StdEncoding.EncodeToString(serviceID)
	service.ID = encServiceID
	return &service, nil
}

func (r *mutationResolver) UpdateService(ctx context.Context, input model.UpdateServiceInput) (*model.Service, error) {
	var service model.Service

	// 入力されたidをデコードする
	decID, err := base64.StdEncoding.DecodeString(input.ID)
	if err != nil {
		log.Fatal(err)
	}

	// :以下の数字を取得
	searchID := string(decID[strings.Index(string(decID), ":")+1:])

	r.DB.First(&service, searchID)
	service.ID = searchID
	service.Name = *input.Name
	service.Email = *input.Email
	service.Password = *input.Password
	r.DB.Save(&service)

	service.ID = input.ID

	return &service, nil
}

func (r *mutationResolver) DeleteService(ctx context.Context, id string) (bool, error) {
	var service model.Service

	// 入力されたidをデコードする
	decID, err := base64.StdEncoding.DecodeString(id)
	if err != nil {
		log.Fatal(err)
	}

	// :以下の数字を取得
	searchID := string(decID[strings.Index(string(decID), ":")+1:])

	r.DB.Where("id = ?", searchID).Delete(&service)
	return true, nil
}

func (r *mutationResolver) CreateServiceAccount(ctx context.Context, input model.CreateServiceAccountInput) (*model.ServiceAccount, error) {
	var serviceAccountModel model.ServiceAccount

	// 最後のレコードを取得
	r.DB.Last(&serviceAccountModel)
	i, _ := strconv.Atoi(serviceAccountModel.ID)
	serviceAccountIDInt := i + 1
	serviceAccountIDStr := strconv.Itoa(serviceAccountIDInt)

	serviceAccount := model.ServiceAccount{
		ID:        serviceAccountIDStr,
		Name:      input.Name,
		Password:  input.Password,
		Email:     input.Email,
		ServiceID: "1", //楽天を想定しているためとりあえず1で固定
	}
	r.DB.Create(&serviceAccount)
	serviceAccountID := []byte("ServiceAccount:" + serviceAccountIDStr) // プライマリキーを型名とセットで記述

	// エンコードする
	encServiceAccountID := base64.StdEncoding.EncodeToString(serviceAccountID)
	serviceAccount.ID = encServiceAccountID

	return &serviceAccount, nil
}

func (r *mutationResolver) UpdateServiceAccount(ctx context.Context, input model.UpdateServiceAccountInput) (*model.ServiceAccount, error) {
	var serviceAccount model.ServiceAccount

	// 入力されたidをデコードする
	decID, err := base64.StdEncoding.DecodeString(input.ID)
	if err != nil {
		log.Fatal(err)
	}

	// :以下の数字を取得
	searchID := string(decID[strings.Index(string(decID), ":")+1:])

	r.DB.First(&serviceAccount, searchID)
	serviceAccount.ID = searchID
	serviceAccount.Name = *input.Name
	serviceAccount.Email = *input.Email
	serviceAccount.Password = *input.Password
	r.DB.Save(&serviceAccount)

	serviceAccount.ID = input.ID

	return &serviceAccount, nil
}

func (r *mutationResolver) DeleteServiceAccount(ctx context.Context, id string) (bool, error) {
	var serviceAccount model.ServiceAccount

	// 入力されたidをデコードする
	decID, err := base64.StdEncoding.DecodeString(id)
	if err != nil {
		log.Fatal(err)
	}

	// :以下の数字を取得
	searchID := string(decID[strings.Index(string(decID), ":")+1:])

	r.DB.Where("id = ?", searchID).Delete(&serviceAccount)
	return true, nil
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.CreateUserInput) (*model.User, error) {
	var userModel model.User

	// 最後のレコードを取得
	r.DB.Last(&userModel)
	i, _ := strconv.Atoi(userModel.ID)
	userIDInt := i + 1
	userIDStr := strconv.Itoa(userIDInt)

	user := model.User{
		ID:        userIDStr,
		Name:      input.Name,
		ServiceID: "1", //楽天を想定しているためとりあえず1で固定
	}
	r.DB.Create(&user)
	userID := []byte("User:" + userIDStr) // プライマリキーを型名とセットで記述

	// エンコードする
	encUserID := base64.StdEncoding.EncodeToString(userID)
	user.ID = encUserID
	return &user, nil
}

func (r *mutationResolver) UpdateUser(ctx context.Context, input model.UpdateUserInput) (*model.User, error) {
	var user model.User

	// 入力されたidをデコードする
	decID, err := base64.StdEncoding.DecodeString(input.ID)
	if err != nil {
		log.Fatal(err)
	}

	// :以下の数字を取得
	searchID := string(decID[strings.Index(string(decID), ":")+1:])

	r.DB.First(&user, searchID)
	user.ID = searchID
	user.Name = *input.Name
	r.DB.Save(&user)

	user.ID = input.ID
	return &user, nil
}

func (r *mutationResolver) DeleteUser(ctx context.Context, id string) (bool, error) {
	// 入力されたidをデコードする
	decID, err := base64.StdEncoding.DecodeString(id)
	if err != nil {
		log.Fatal(err)
	}

	// :以下の数字を取得
	searchID := string(decID[strings.Index(string(decID), ":")+1:])
	var user model.User

	r.DB.Where("id = ?", searchID).Delete(&user)
	return true, nil
}

func (r *mutationResolver) PostPersonalTag(ctx context.Context, input []*model.PostPersonalTagInput) ([]*model.PersonalTag, error) {
	// personalTagを投稿した時の関数
	for _, personalTagInfo := range input {
		var personalTagModel model.PersonalTag
		// 最後のレコードを取得
		r.DB.Last(&personalTagModel)
		i, _ := strconv.Atoi(personalTagModel.ID)
		personalTagIDInt := i + 1
		personalTagIDStr := strconv.Itoa(personalTagIDInt)

		personalTag := model.PersonalTag{
			ID:     personalTagIDStr,
			Name:   personalTagInfo.Name,
			UserID: "1", //ここは将来その時に作成しているユーザーのidを入れる 現状は1で設定
		}
		personalTags = append(personalTags, &personalTag)
		r.DB.Create(&personalTag)
		personalTagStr := strconv.Itoa(len(personalTags))
		personalTagID := []byte("PersonalTag:" + personalTagStr)
		// エンコードする
		encPersonalTagID := base64.StdEncoding.EncodeToString(personalTagID)
		personalTag.ID = encPersonalTagID
	}
	return personalTags, nil
}

func (r *mutationResolver) PostReviewTag(ctx context.Context, input []*model.PostReviewTagInput) ([]*model.ReviewTag, error) {
	// reviewTagを投稿した時の関数
	for _, reviewTagInfo := range input {
		var reviewTagModel model.ReviewTag
		// 最後のレコードを取得
		r.DB.Last(&reviewTagModel)
		i, _ := strconv.Atoi(reviewTagModel.ID)
		reviewTagIDInt := i + 1
		reviewTagIDStr := strconv.Itoa(reviewTagIDInt)
		reviewTag := model.ReviewTag{
			ID:     reviewTagIDStr,
			Name:   reviewTagInfo.Name,
			UserID: "1", //ここは将来その時に作成しているユーザーのidを入れる 現状は1で設定
		}
		reviewTags = append(reviewTags, &reviewTag)
		r.DB.Create(&reviewTag)
		reviewTagStr := strconv.Itoa(len(reviewTags))
		reviewTagID := []byte("ReviewTag:" + reviewTagStr)
		// エンコードする
		encReviewTagID := base64.StdEncoding.EncodeToString(reviewTagID)
		reviewTag.ID = encReviewTagID
	}
	return reviewTags, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
