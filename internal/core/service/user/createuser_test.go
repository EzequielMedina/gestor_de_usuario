package user_test

import (
	"gestor_de_usuario/internal/adapter/handler/api/request"
	"gestor_de_usuario/internal/core/domain"
	"gestor_de_usuario/internal/core/service/user"
	"gestor_de_usuario/mocks"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"testing"
)

type createTestedInput struct {
	user *request.UserRequest
}
type createTestedOutput struct {
	user *domain.User
	err  error
}

func TestUserService_CreateUser(t *testing.T) {

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
		mocks             func(userRepo *mocks.MockUserRepository, util *mocks.MockUtilService, input *request.UserRequest, output createTestedOutput)
		input             *request.UserRequest
		expected          createTestedOutput
		assertionFunction func(subT *testing.T, user *domain.User, err error)
	}{
		"email empty": {
			mocks: func(userRepo *mocks.MockUserRepository, util *mocks.MockUtilService, input *request.UserRequest, output createTestedOutput) {
				//userRepo.EXPECT().CreateUser(gomock.Eq(input)).Return(output.user, output.err)
			},
			input: &request.UserRequest{
				Name:     userName,
				LastName: userLastName,
				Email:    "",
				Password: userPassword,
			},
			expected: createTestedOutput{
				user: nil,
				err:  domain.ErrEmailRequired,
			},
			assertionFunction: func(subT *testing.T, user *domain.User, err error) {
				assert.Equal(subT, domain.ErrEmailRequired, err)
				assert.Nil(subT, user)
			},
		},
		"validate email": {
			mocks: func(userRepo *mocks.MockUserRepository, util *mocks.MockUtilService, input *request.UserRequest, output createTestedOutput) {
				//userRepo.EXPECT().CreateUser(input).Return(output.user, output.err)
				util.EXPECT().IsValidEmail(input.Email).Return(false)
			},
			input: &request.UserRequest{
				Name:     userName,
				LastName: userLastName,
				Email:    "email",
				Password: userPassword,
			},

			expected: createTestedOutput{
				user: nil,
				err:  domain.ErrInvalidEmail,
			},
			assertionFunction: func(subT *testing.T, user *domain.User, err error) {
				assert.Equal(subT, domain.ErrInvalidEmail, err)
				assert.Nil(subT, user)
			},
		},
		"exist email": {
			mocks: func(userRepo *mocks.MockUserRepository, util *mocks.MockUtilService, input *request.UserRequest, output createTestedOutput) {
				userRepo.EXPECT().GetUserByEmail(input.Email).Return(&userOutput, nil)
				util.EXPECT().IsValidEmail(input.Email).Return(true)
			},
			input: &request.UserRequest{
				Name:     userName,
				LastName: userLastName,
				Email:    userEmail,
				Password: userPassword,
			},
			expected: createTestedOutput{
				user: nil,
				err:  domain.ErrEmailAlreadyExists,
			},
			assertionFunction: func(subT *testing.T, user *domain.User, err error) {
				assert.Equal(subT, domain.ErrEmailAlreadyExists, err)
				assert.Nil(subT, user)
			},
		},
		"email internal error": {
			mocks: func(userRepo *mocks.MockUserRepository, util *mocks.MockUtilService, input *request.UserRequest, output createTestedOutput) {
				util.EXPECT().IsValidEmail(input.Email).Return(true)
				userRepo.EXPECT().GetUserByEmail(input.Email).Return(nil, domain.ErrInternal)
			},
			input: &request.UserRequest{
				Name:     userName,
				LastName: userLastName,
				Email:    userEmail,
				Password: userPassword,
			},
			expected: createTestedOutput{
				user: nil,
				err:  domain.ErrInternal,
			},
			assertionFunction: func(subT *testing.T, user *domain.User, err error) {
				assert.Equal(subT, domain.ErrInternal, err)
				assert.Nil(subT, user)
			},
		},
		"name empty": {
			mocks: func(userRepo *mocks.MockUserRepository, util *mocks.MockUtilService, input *request.UserRequest, output createTestedOutput) {
				util.EXPECT().IsValidEmail(input.Email).Return(true)
				userRepo.EXPECT().GetUserByEmail(input.Email).Return(nil, nil)
			},
			input: &request.UserRequest{
				Name:     "",
				LastName: userLastName,
				Email:    userEmail,
				Password: userPassword,
			},
			expected: createTestedOutput{
				user: nil,
				err:  domain.ErrNameRequired,
			},
			assertionFunction: func(subT *testing.T, user *domain.User, err error) {
				assert.Equal(subT, domain.ErrNameRequired, err)
				assert.Nil(subT, user)
			},
		},
		"last name empty": {
			mocks: func(userRepo *mocks.MockUserRepository, util *mocks.MockUtilService, input *request.UserRequest, output createTestedOutput) {
				util.EXPECT().IsValidEmail(input.Email).Return(true)
				userRepo.EXPECT().GetUserByEmail(input.Email).Return(nil, nil)
			},
			input: &request.UserRequest{
				Name:     userName,
				LastName: "",
				Email:    userEmail,
				Password: userPassword,
			},
			expected: createTestedOutput{
				user: nil,
				err:  domain.ErrLastNameRequired,
			},
			assertionFunction: func(subT *testing.T, user *domain.User, err error) {
				assert.Equal(subT, domain.ErrLastNameRequired, err)
				assert.Nil(subT, user)
			},
		},
		"password empty": {
			mocks: func(userRepo *mocks.MockUserRepository, util *mocks.MockUtilService, input *request.UserRequest, output createTestedOutput) {
				util.EXPECT().IsValidEmail(input.Email).Return(true)
				userRepo.EXPECT().GetUserByEmail(input.Email).Return(nil, nil)
			},
			input: &request.UserRequest{
				Name:     userName,
				LastName: userLastName,
				Email:    userEmail,
				Password: "",
			},
			expected: createTestedOutput{
				user: nil,
				err:  domain.ErrPasswordRequired,
			},
			assertionFunction: func(subT *testing.T, user *domain.User, err error) {
				assert.Equal(subT, domain.ErrPasswordRequired, err)
				assert.Nil(subT, user)
			},
		},
		"password length": {
			mocks: func(userRepo *mocks.MockUserRepository, util *mocks.MockUtilService, input *request.UserRequest, output createTestedOutput) {
				util.EXPECT().IsValidEmail(input.Email).Return(true)
				userRepo.EXPECT().GetUserByEmail(input.Email).Return(nil, nil)
				util.EXPECT().IsValidPassword(input.Password).Return(false)
			},
			input: &request.UserRequest{
				Name:     userName,
				LastName: userLastName,
				Email:    userEmail,
				Password: gofakeit.Password(true, true, true, true, false, 7),
			},
			expected: createTestedOutput{
				user: nil,
				err:  domain.ErrInvalidPassword,
			},
			assertionFunction: func(subT *testing.T, user *domain.User, err error) {
				assert.Equal(subT, domain.ErrInvalidPassword, err)
				assert.Nil(subT, user)
			},
		},
		"hashPassword error": {
			mocks: func(userRepo *mocks.MockUserRepository, util *mocks.MockUtilService, input *request.UserRequest, output createTestedOutput) {
				util.EXPECT().IsValidEmail(input.Email).Return(true)
				userRepo.EXPECT().GetUserByEmail(input.Email).Return(nil, nil)
				util.EXPECT().IsValidPassword(input.Password).Return(true)
				util.EXPECT().HashPassword(input.Password).Return("", domain.ErrInternal)

			},
			input: &request.UserRequest{
				Name:     userName,
				LastName: userLastName,
				Email:    userEmail,
				Password: userPassword,
			},
			expected: createTestedOutput{
				user: nil,
				err:  domain.ErrInternal,
			},
			assertionFunction: func(subT *testing.T, user *domain.User, err error) {
				assert.Equal(subT, domain.ErrInternal, err)
				assert.Nil(subT, user)
			},
		},
		"create user internal error": {
			mocks: func(userRepo *mocks.MockUserRepository, util *mocks.MockUtilService, input *request.UserRequest, output createTestedOutput) {

				util.EXPECT().IsValidEmail(input.Email).Return(true)
				userRepo.EXPECT().GetUserByEmail(input.Email).Return(nil, nil)
				util.EXPECT().IsValidPassword(input.Password).Return(true)
				util.EXPECT().HashPassword(input.Password).Return("hashedPassword", nil)
				userRepo.EXPECT().CreateUser(gomock.Any()).Return(nil, output.err)
			},
			input: &request.UserRequest{
				Name:     userName,
				LastName: userLastName,
				Email:    userEmail,
				Password: userPassword,
			},
			expected: createTestedOutput{
				user: nil,
				err:  domain.ErrInternal,
			},
			assertionFunction: func(subT *testing.T, user *domain.User, err error) {
				assert.Equal(subT, domain.ErrInternal, err)
				assert.Nil(subT, user)
			},
		},
		"create user success": {
			mocks: func(userRepo *mocks.MockUserRepository, util *mocks.MockUtilService, input *request.UserRequest, output createTestedOutput) {
				util.EXPECT().IsValidEmail(input.Email).Return(true)
				userRepo.EXPECT().GetUserByEmail(input.Email).Return(nil, nil)
				util.EXPECT().IsValidPassword(input.Password).Return(true)
				util.EXPECT().HashPassword(input.Password).Return("hashedPassword", nil)
				userRepo.EXPECT().CreateUser(gomock.Any()).Return(&output.user, output.err)
			},
			input: &request.UserRequest{
				Name:     userName,
				LastName: userLastName,
				Email:    userEmail,
				Password: userPassword,
			},
			expected: createTestedOutput{
				user: &userOutput,
				err:  nil,
			},
			assertionFunction: func(subT *testing.T, user *domain.User, err error) {
				assert.Equal(subT, userOutput.FirstName, user.FirstName)
				assert.Equal(subT, userOutput.LastName, user.LastName)
				assert.Equal(subT, userOutput.Email, user.Email)
				// Validar que la contraseña ha sido hasheada
				assert.NotEqual(subT, userPassword, user.Password)
				// Opcional: si quieres comprobar que el hash es válido, puedes hacer:
				//errHash := util.ComparePassword(user.Password, userOutput.Password)
				//assert.NoError(subT, errHash)
			},
		},
	}

	for name, test := range testTable {
		t.Run(name, func(subT *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockUserRepository := mocks.NewMockUserRepository(ctrl)
			mockUtil := mocks.NewMockUtilService(ctrl)
			test.mocks(mockUserRepository, mockUtil, test.input, test.expected)

			s := user.NewUserService(mockUserRepository, mockUtil)

			createUser, err := s.CreateUser(test.input)
			test.assertionFunction(subT, createUser, err)

		})
	}

}
