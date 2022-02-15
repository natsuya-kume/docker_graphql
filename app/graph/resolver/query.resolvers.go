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

func (r *queryResolver) Admin(ctx context.Context, id string) (*model.Admin, error) {
	var admin model.Admin

	decID := utils.Decode(id)
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
		admin.ID = utils.Encode("Admin:", admin.ID)
	}
	page := utils.Page(limit, offset)
	paginationLength := utils.PaginationLength(len(admins), limit)
	hasNextpage := utils.HasNextPage(len(admins), limit, offset)
	hasPreviouspage := utils.HasPreviousPage(offset)

	var totalCount int64
	r.DB.Model(&admins).Group("id").Count(&totalCount)

	r.DB.Limit(limit).Offset(*offset).Find(&admins)

	adminPagination := &model.AdminPagination{
		Nodes: admins,
		PageInfo: &model.PaginationInfo{
			Page:             page,
			PaginationLength: paginationLength,
			HasNextPage:      hasNextpage,
			HasPreviousPage:  hasPreviouspage,
			Count:            len(admins),
			TotalCount:       int(totalCount),
		},
	}
	return adminPagination, nil
}

func (r *queryResolver) Service(ctx context.Context, id string) (*model.Service, error) {
	var service model.Service

	decID := utils.Decode(id)
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
		service.ID = utils.Encode("Service:", service.ID)
	}
	page := utils.Page(limit, offset)
	paginationLength := utils.PaginationLength(len(services), limit)
	hasNextpage := utils.HasNextPage(len(services), limit, offset)
	hasPreviouspage := utils.HasPreviousPage(offset)

	var totalCount int64
	r.DB.Model(&services).Group("id").Count(&totalCount)

	r.DB.Limit(limit).Offset(*offset).Find(&services)

	servicePagination := &model.ServicePagination{
		Nodes: services,
		PageInfo: &model.PaginationInfo{ //仮のデータ
			Page:             page,
			PaginationLength: paginationLength,
			HasNextPage:      hasNextpage,
			HasPreviousPage:  hasPreviouspage,
			Count:            len(services),
			TotalCount:       int(totalCount),
		},
	}
	return servicePagination, nil
}

func (r *queryResolver) ServiceAccount(ctx context.Context, id string) (*model.ServiceAccount, error) {
	var serviceAccount model.ServiceAccount

	decID := utils.Decode(id)
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
		serviceAccount.ID = utils.Encode("ServiceAccount:", serviceAccount.ID)
	}
	page := utils.Page(limit, offset)
	paginationLength := utils.PaginationLength(len(serviceAccounts), limit)
	hasNextpage := utils.HasNextPage(len(serviceAccounts), limit, offset)
	hasPreviouspage := utils.HasPreviousPage(offset)

	var totalCount int64
	r.DB.Model(&serviceAccounts).Group("id").Count(&totalCount)

	r.DB.Limit(limit).Offset(*offset).Find(&serviceAccounts)
	serviceAccountsPagination := &model.ServiceAccountPagination{
		Nodes: serviceAccounts,
		PageInfo: &model.PaginationInfo{ //仮のデータ
			Page:             page,
			PaginationLength: paginationLength,
			HasNextPage:      hasNextpage,
			HasPreviousPage:  hasPreviouspage,
			Count:            len(serviceAccounts),
			TotalCount:       int(totalCount),
		},
	}
	return serviceAccountsPagination, nil
}

func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	var user model.User

	decID := utils.Decode(id)
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
		user.ID = utils.Encode("User:", user.ID)
	}
	page := utils.Page(limit, offset)
	paginationLength := utils.PaginationLength(len(users), limit)
	hasNextpage := utils.HasNextPage(len(users), limit, offset)
	hasPreviouspage := utils.HasPreviousPage(offset)

	var totalCount int64
	r.DB.Model(&users).Group("id").Count(&totalCount)

	r.DB.Limit(limit).Offset(*offset).Find(&users)
	userPagination := &model.UserPagination{
		Nodes: users,
		PageInfo: &model.PaginationInfo{ //仮のデータ
			Page:             page,
			PaginationLength: paginationLength,
			HasNextPage:      hasNextpage,
			HasPreviousPage:  hasPreviouspage,
			Count:            len(users),
			TotalCount:       int(totalCount),
		},
	}
	return userPagination, nil
}

func (r *queryResolver) PersonalTag(ctx context.Context, id string) (*model.PersonalTag, error) {
	var personalTag model.PersonalTag

	decID := utils.Decode(id)
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
		personalTag.ID = utils.Encode("PersonalTag:", personalTag.ID)
	}
	page := utils.Page(limit, offset)
	paginationLength := utils.PaginationLength(len(personalTags), limit)
	hasNextpage := utils.HasNextPage(len(personalTags), limit, offset)
	hasPreviouspage := utils.HasPreviousPage(offset)

	var totalCount int64
	r.DB.Model(&personalTags).Group("id").Count(&totalCount)

	r.DB.Limit(limit).Offset(*offset).Find(&personalTags)
	personalTagPagination := &model.PersonalTagPagination{
		Nodes: personalTags,
		PageInfo: &model.PaginationInfo{ //仮のデータ
			Page:             page,
			PaginationLength: paginationLength,
			HasNextPage:      hasNextpage,
			HasPreviousPage:  hasPreviouspage,
			Count:            len(personalTags),
			TotalCount:       int(totalCount),
		},
	}
	return personalTagPagination, nil
}

func (r *queryResolver) ReviewTag(ctx context.Context, id string) (*model.ReviewTag, error) {
	var reviewTag model.ReviewTag

	decID := utils.Decode(id)
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
		reviewTag.ID = utils.Encode("ReviewTag:", reviewTag.ID)
	}
	page := utils.Page(limit, offset)
	paginationLength := utils.PaginationLength(len(reviewTags), limit)
	hasNextpage := utils.HasNextPage(len(reviewTags), limit, offset)
	hasPreviouspage := utils.HasPreviousPage(offset)

	var totalCount int64
	r.DB.Model(&reviewTags).Group("id").Count(&totalCount)

	r.DB.Limit(limit).Offset(*offset).Find(&reviewTags)
	reviewTagPagination := &model.ReviewTagPagination{
		Nodes: reviewTags,
		PageInfo: &model.PaginationInfo{ //仮のデータ
			Page:             page,
			PaginationLength: paginationLength,
			HasNextPage:      hasNextpage,
			HasPreviousPage:  hasPreviouspage,
			Count:            len(reviewTags),
			TotalCount:       int(totalCount),
		},
	}
	return reviewTagPagination, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
