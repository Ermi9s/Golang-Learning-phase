package usecase

import (
	"github.com/Ermi9s.Golang-Learning-phase/Clean-Architecture-TaskManager/domain"
	"golang.org/x/crypto/bcrypt"
)

type User_usecase struct {
	User_repo domain.User_Repository_interface 
}

func New_User_Usecase(user_repo domain.User_Repository_interface) domain.User_Usecase_interface {
	return &User_usecase{
		User_repo: user_repo,
	}
}

func (userusecase *User_usecase)GetUser(id string) (domain.User, error) {
	doc , err := userusecase.User_repo.GetUserDocumentById(id)
	if err != nil {
		return domain.User{},err
	}
	return doc, nil
}

func (userusecase *User_usecase)GetUsers() ([]domain.User, error) {
	filter := make(map[string]string)
	decoded, err := userusecase.User_repo.GetUserDocumentByFilter(filter)
	if err != nil {
		return []domain.User{} , err
	}
	return decoded,nil
}

func (userusecase *User_usecase)CreateUser(model domain.User) (domain.User, error) {
	user := model
	hasshedPasskey,err := bcrypt.GenerateFromPassword([]byte(user.Password) , bcrypt.DefaultCost);
	if err != nil {
		return domain.User{},err
	}
	user.Password = string(hasshedPasskey)
	id,err := userusecase.User_repo.InsertUserDocument(user)

	if err!= nil {
		return domain.User{} , err
	}
	new_model,err := userusecase.GetUser(id)
	if err != nil {
		return domain.User{},err
	}
	
	return new_model,nil
}

func (userusecase *User_usecase)UpdateUser(id string , model domain.User) (domain.User, error) {
	err := userusecase.User_repo.UpdateUserDocumentById(id , model)
	if err != nil {
		return domain.User{},err
	}
	return model , nil
}

func (userusecase *User_usecase)DeleteUser(id string) error {
	err := userusecase.User_repo.DeleteUserDocument(id)
	if err != nil {
		return err
	}
	return nil
}

func (userusecase *User_usecase)LogIn(model domain.AuthUser) (domain.User , error) {
	filter := make(map[string]string)
	filter["username"] = model.UserName
	filter["email"] = model.Email

	result,err := userusecase.User_repo.GetUserDocumentByFilter(filter)
	if err != nil {
		return domain.User{} , err
	}

	user := result[0]
	err = bcrypt.CompareHashAndPassword([]byte(user.Password) , []byte(model.Password))
	if err != nil {
		return domain.User{} , err
	}

	return user,nil
}


func (usecase *User_usecase)Promote(id string)(domain.User , error) {
	user , err := usecase.GetUser(id)
	if err != nil {
		return domain.User{},err
	}
	user.Is_admin = true

	user,err = usecase.UpdateUser(id , user)
	if err != nil {
		return domain.User{},err
	}
	return user , nil
}