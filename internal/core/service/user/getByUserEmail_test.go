package user

import (
	"gestor_de_usuario/internal/core/domain"
	"gestor_de_usuario/mocks"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"testing"
)

type getByUserEmailTestedInput struct {
	email string
}
type getByUserEmailTestedOutput struct {
	user *domain.User
	err  error
}

func TestUserService_GetUserByEmail(t *testing.T) {
	userName := gofakeit.Name()
	userLastName := gofakeit.LastName()
	userEmail := gofakeit.Email()
	userPassword := gofakeit.Password(true, true, true, true, false, 8)

	userOutput := domain.User{
		ID:        gofakeit.UUID(),
		FirstName: userName,
		LastName:  userLastName,
		Email:     userEmail,
		Password:  userPassword,
		CreatedAt: gofakeit.Date(),
		UpdatedAt: gofakeit.Date(),
		Active:    true,
	}

	testTable := map[string]struct {
		mocks     func(userRepor *mocks.MockUserRepository, util *mocks.MockUtilService, input *getByUserEmailTestedInput, output getByUserEmailTestedOutput)
		input     *getByUserEmailTestedInput
		expedted  getByUserEmailTestedOutput
		assertion func(t *testing.T, user *domain.User, err error)
	}{
		"email empty": {
			mocks: func(userRepor *mocks.MockUserRepository, util *mocks.MockUtilService, input *getByUserEmailTestedInput, output getByUserEmailTestedOutput) {

			},
			input: &getByUserEmailTestedInput{
				email: "",
			},
			expedted: getByUserEmailTestedOutput{
				user: nil,
				err:  domain.ErrEmailRequired,
			},
			assertion: func(subT *testing.T, user *domain.User, err error) {
				assert.Equal(subT, domain.ErrEmailRequired, err)
				assert.Nil(subT, user)
			},
		},
		"email invalid format": {
			mocks: func(userRepor *mocks.MockUserRepository, util *mocks.MockUtilService, input *getByUserEmailTestedInput, output getByUserEmailTestedOutput) {
				util.EXPECT().IsValidEmail(input.email).Return(false)
			},
			input: &getByUserEmailTestedInput{
				email: "email",
			},
			expedted: getByUserEmailTestedOutput{
				user: nil,
				err:  domain.ErrInvalidEmail,
			},
			assertion: func(subT *testing.T, user *domain.User, err error) {
				assert.Equal(subT, domain.ErrInvalidEmail, err)
				assert.Nil(subT, user)
			},
		},
		"email not exist": {
			mocks: func(userRepor *mocks.MockUserRepository, util *mocks.MockUtilService, input *getByUserEmailTestedInput, output getByUserEmailTestedOutput) {
				util.EXPECT().IsValidEmail(input.email).Return(true)
				userRepor.EXPECT().GetUserByEmail(input.email).Return(nil, domain.ErrEmailNotFound)

			},
			input: &getByUserEmailTestedInput{
				email: gofakeit.Email(),
			},
			expedted: getByUserEmailTestedOutput{
				user: nil,
				err:  domain.ErrEmailNotFound,
			},
			assertion: func(subT *testing.T, user *domain.User, err error) {
				assert.Equal(subT, domain.ErrEmailNotFound, err)
				assert.Nil(subT, user)
			},
		},
		"email exist": {
			mocks: func(userRepor *mocks.MockUserRepository, util *mocks.MockUtilService, input *getByUserEmailTestedInput, output getByUserEmailTestedOutput) {
				util.EXPECT().IsValidEmail(input.email).Return(true)
				userRepor.EXPECT().GetUserByEmail(input.email).Return(output.user, output.err)
			},
			input: &getByUserEmailTestedInput{
				userEmail,
			},
			expedted: getByUserEmailTestedOutput{
				user: &userOutput,
				err:  nil,
			},
			assertion: func(subT *testing.T, user *domain.User, err error) {
				assert.Nil(subT, err)
				assert.NotNil(subT, user)
				assert.Equal(subT, userEmail, user.Email)
			},
		},
	}

	for name, tt := range testTable {
		t.Run(name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockUserRepositoy := mocks.NewMockUserRepository(ctrl)
			mockUtilService := mocks.NewMockUtilService(ctrl)

			tt.mocks(mockUserRepositoy, mockUtilService, tt.input, tt.expedted)
			s := NewUserService(mockUserRepositoy, mockUtilService)

			userDB, err := s.GetUserByEmail(tt.input.email)
			tt.assertion(t, userDB, err)

		})
	}

}
