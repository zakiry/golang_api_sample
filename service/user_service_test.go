package service

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/zakiry/golang_api_sample/entity"
	"github.com/zakiry/golang_api_sample/mock_database"
	"github.com/zakiry/golang_api_sample/proto"
)

func TestUserService_GetUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userId := int64(1)

	mockUser := mock_database.NewMockUser(ctrl)
	mockUser.EXPECT().GetById(userId).Return(entity.User{Id: 1, Name: "piyo piyo", Age: 23})

	service := userService{user: mockUser}

	actual, _ := service.GetUser(nil, &sample_proto.GetUserRequest{Id: 1})

	assert.Equal(t, int64(1), actual.Id)
	assert.Equal(t, "piyo piyo", actual.Name)
	assert.Equal(t, int32(23), actual.Age)
}
