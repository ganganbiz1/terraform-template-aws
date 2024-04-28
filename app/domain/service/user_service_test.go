package service_test

import (
	"context"
	"testing"

	"github.com/ganganbiz1/go-echo-gorm-template/app/domain"
	"github.com/ganganbiz1/go-echo-gorm-template/app/domain/entity"
	mock_repository "github.com/ganganbiz1/go-echo-gorm-template/app/domain/repository/mock"
	"github.com/ganganbiz1/go-echo-gorm-template/app/domain/service"
	"github.com/go-playground/assert/v2"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

const (
	msgTestFormat = "got: %v, want: %v"
)

func Test_User_Create(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userRepoMock := mock_repository.NewMockIfUserRepository(ctrl)

	service := service.NewUserService(userRepoMock)

	type args struct {
		e *entity.User
	}

	user := &entity.User{
		ID:            0,
		Email:         "test@test.com",
		Name:          "テスト太郎",
		CognitoUserID: "",
	}

	tests := []struct {
		testName string
		args     args
		setMock  func()
		wantErr  error
	}{
		{
			testName: "正常系",
			args:     args{e: user},
			setMock: func() {
				userRepoMock.EXPECT().Create(gomock.Any(), user).Return(nil)
			},
			wantErr: nil,
		},
		{
			testName: "異常系",
			args:     args{e: user},
			setMock: func() {
				userRepoMock.EXPECT().Create(gomock.Any(), user).Return(domain.ErrInternal)
			},
			wantErr: domain.ErrInternal,
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			tt.setMock()
			err := service.Create(context.Background(), tt.args.e)
			if tt.wantErr != nil {
				require.Error(t, err, msgTestFormat, err, tt.wantErr)
			} else {
				require.NoError(t, err, msgTestFormat, err, nil)
			}
			assert.Equal(t, err, tt.wantErr)
		})
	}
}

func Test_User_Search(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userRepoMock := mock_repository.NewMockIfUserRepository(ctrl)

	service := service.NewUserService(userRepoMock)

	type args struct {
		id int
	}

	expectedUser := &entity.User{
		ID:            1,
		Email:         "test@test.com",
		Name:          "テスト太郎",
		CognitoUserID: "abcdefg",
	}

	tests := []struct {
		testName string
		args     args
		setMock  func()
		want     *entity.User
		wantErr  error
	}{
		{
			testName: "正常系",
			args:     args{id: 1},
			setMock: func() {
				userRepoMock.EXPECT().Get(gomock.Any(), gomock.Any()).Return(expectedUser, nil)
			},
			want:    expectedUser,
			wantErr: nil,
		},
		{
			testName: "異常系: 取得エラー",
			args:     args{id: 1},
			setMock: func() {
				userRepoMock.EXPECT().Get(gomock.Any(), gomock.Any()).Return(nil, domain.ErrInternal)
			},
			want:    nil,
			wantErr: domain.ErrInternal,
		},
		{
			testName: "異常系: NotFoundエラー",
			args:     args{id: 1},
			setMock: func() {
				userRepoMock.EXPECT().Get(gomock.Any(), gomock.Any()).Return(nil, domain.ErrNotFound)
			},
			want:    nil,
			wantErr: domain.ErrNotFound,
		},
		{
			testName: "異常系: NotFound",
			args:     args{id: 1},
			setMock: func() {
				userRepoMock.EXPECT().Get(gomock.Any(), gomock.Any()).Return(nil, nil)
			},
			want:    nil,
			wantErr: domain.ErrNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			tt.setMock()
			got, err := service.Search(context.Background(), tt.args.id)
			if tt.wantErr != nil {
				require.Error(t, err, msgTestFormat, err, tt.wantErr)
			} else {
				require.NoError(t, err, msgTestFormat, err, nil)
			}
			assert.Equal(t, got, tt.want)
			assert.Equal(t, err, tt.wantErr)
		})
	}
}
