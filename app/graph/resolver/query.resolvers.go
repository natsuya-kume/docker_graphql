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

func (r *queryResolver) Service(ctx context.Context, id string) (*model.Service, error) {
	var service model.Service
	// 入力されたidをデコードする
	decID, err := base64.StdEncoding.DecodeString(id)
	if err != nil {
		log.Fatal(err)
	}

	// :以下の数字を取得
	searchID := string(decID[strings.Index(string(decID), ":")+1:])
	if err := r.DB.Find(&service, searchID).Error; err != nil {
		return nil, err
	}
	service.ID = id
	return &service, nil
}

func (r *queryResolver) Services(ctx context.Context, limit int, offset *int) (*model.ServicePagination, error) {
	var services []*model.Service
	if err := r.DB.Find(&services).Error; err != nil {
		return nil, err
	}
	for _, service := range services {
		serviceID := []byte("Service:" + service.ID) // プライマリキーを型名とセットで記述
		// エンコードする
		encServiceID := base64.StdEncoding.EncodeToString(serviceID)
		service.ID = encServiceID
	}
	servicePagination := &model.ServicePagination{
		Nodes: services,
		PageInfo: &model.PaginationInfo{ //仮のデータ
			Page:             1,
			PaginationLength: 1,
			HasNextPage:      false,
			HasPreviousPage:  false,
			Count:            len(services),
			TotalCount:       1,
		},
	}
	return servicePagination, nil
}

func (r *queryResolver) ServiceAccount(ctx context.Context, id string) (*model.ServiceAccount, error) {
	var serviceAccount model.ServiceAccount
	// 入力されたidをデコードする
	decID, err := base64.StdEncoding.DecodeString(id)
	if err != nil {
		log.Fatal(err)
	}
	// :以下の数字を取得
	searchID := string(decID[strings.Index(string(decID), ":")+1:])
	if err := r.DB.Find(&serviceAccount, searchID).Error; err != nil {
		return nil, err
	}

	serviceAccount.ID = id
	return &serviceAccount, nil
}

func (r *queryResolver) ServiceAccounts(ctx context.Context, limit int, offset *int) (*model.ServiceAccountPagination, error) {
	var serviceAccounts []*model.ServiceAccount
	if err := r.DB.Find(&serviceAccounts).Error; err != nil {
		return nil, err
	}
	for _, serviceAccount := range serviceAccounts {
		serviceAccountID := []byte("ServiceAccount:" + serviceAccount.ID) // プライマリキーを型名とセットで記述
		// エンコードする
		encServiceAccountID := base64.StdEncoding.EncodeToString(serviceAccountID)
		serviceAccount.ID = encServiceAccountID
	}
	serviceAccountsPagination := &model.ServiceAccountPagination{
		Nodes: serviceAccounts,
		PageInfo: &model.PaginationInfo{ //仮のデータ
			Page:             1,
			PaginationLength: 1,
			HasNextPage:      false,
			HasPreviousPage:  false,
			Count:            len(serviceAccounts),
			TotalCount:       1,
		},
	}
	return serviceAccountsPagination, nil
}

func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	var user model.User
	// 入力されたidをデコードする
	decID, err := base64.StdEncoding.DecodeString(id)
	if err != nil {
		log.Fatal(err)
	}
	// :以下の数字を取得
	searchID := string(decID[strings.Index(string(decID), ":")+1:])
	if err := r.DB.Find(&user, searchID).Error; err != nil {
		return nil, err
	}
	user.ID = id
	return &user, nil
}

func (r *queryResolver) Users(ctx context.Context, limit int, offset *int, name *string, email *string, password *string) (*model.UserPagination, error) {
	var users []*model.User
	if err := r.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	for _, user := range users {
		userID := []byte("User:" + user.ID) // プライマリキーを型名とセットで記述
		// エンコードする
		encUserID := base64.StdEncoding.EncodeToString(userID)
		user.ID = encUserID
	}
	userPagination := &model.UserPagination{
		Nodes: users,
		PageInfo: &model.PaginationInfo{ //仮のデータ
			Page:             1,
			PaginationLength: 1,
			HasNextPage:      false,
			HasPreviousPage:  false,
			Count:            len(users),
			TotalCount:       1,
		},
	}
	return userPagination, nil
}

func (r *queryResolver) PersonalTag(ctx context.Context, id string) (*model.PersonalTag, error) {
	var personalTag model.PersonalTag
	// 入力されたidをデコードする
	decID, err := base64.StdEncoding.DecodeString(id)
	if err != nil {
		log.Fatal(err)
	}
	// :以下の数字を取得
	searchID := string(decID[strings.Index(string(decID), ":")+1:])
	if err := r.DB.Find(&personalTag, searchID).Error; err != nil {
		return nil, err
	}
	personalTag.ID = id

	return &personalTag, nil
}

func (r *queryResolver) PersonalTags(ctx context.Context, limit int, offset *int) (*model.PersonalTagPagination, error) {
	var personalTags []*model.PersonalTag
	if err := r.DB.Find(&personalTags).Error; err != nil {
		return nil, err
	}
	for _, personalTag := range personalTags {
		personalTagID := []byte("PersonalTag:" + personalTag.ID) // プライマリキーを型名とセットで記述
		// エンコードする
		encPersonalTagID := base64.StdEncoding.EncodeToString(personalTagID)
		personalTag.ID = encPersonalTagID
	}
	personalTagPagination := &model.PersonalTagPagination{
		Nodes: personalTags,
		PageInfo: &model.PaginationInfo{ //仮のデータ
			Page:             1,
			PaginationLength: 1,
			HasNextPage:      false,
			HasPreviousPage:  false,
			Count:            len(personalTags),
			TotalCount:       1,
		},
	}
	return personalTagPagination, nil
}

func (r *queryResolver) ReviewTag(ctx context.Context, id string) (*model.ReviewTag, error) {
	var reviewTag model.ReviewTag
	// 入力されたidをデコードする
	decID, err := base64.StdEncoding.DecodeString(id)
	if err != nil {
		log.Fatal(err)
	}
	// :以下の数字を取得
	searchID := string(decID[strings.Index(string(decID), ":")+1:])
	if err := r.DB.Find(&reviewTag, searchID).Error; err != nil {
		return nil, err
	}
	reviewTag.ID = id
	return &reviewTag, nil
}

func (r *queryResolver) ReviewTags(ctx context.Context, limit int, offset *int) (*model.ReviewTagPagination, error) {
	var reviewTags []*model.ReviewTag
	if err := r.DB.Find(&reviewTags).Error; err != nil {
		return nil, err
	}
	for _, reviewTag := range reviewTags {
		reviewTagID := []byte("ReviewTag:" + reviewTag.ID) // プライマリキーを型名とセットで記述
		// エンコードする
		encReviewTagID := base64.StdEncoding.EncodeToString(reviewTagID)
		reviewTag.ID = encReviewTagID
	}
	reviewTagPagination := &model.ReviewTagPagination{
		Nodes: reviewTags,
		PageInfo: &model.PaginationInfo{ //仮のデータ
			Page:             1,
			PaginationLength: 1,
			HasNextPage:      false,
			HasPreviousPage:  false,
			Count:            len(reviewTags),
			TotalCount:       1,
		},
	}
	return reviewTagPagination, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
