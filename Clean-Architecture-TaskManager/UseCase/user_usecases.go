package usecase

import (
	"github.com/Ermi9s.Golang-Learning-phase/Clean-Architecture-TaskManager/domain"
	"golang.org/x/crypto/bcrypt"
)


func (userusecase *UseCaseData)GetUser(id string) (domain.User, error) {
	doc , err := userusecase.Repo.GetUserDocumentById(id)
	if err != nil {
		return domain.User{},err
	}
	return doc, nil
}

func (userusecase *UseCaseData)GetUsers() ([]domain.User, error) {
	filter := make(map[string]string)
	decoded, err := userusecase.Repo.GetUserDocumentByFilter(filter)
	if err != nil {
		return []domain.User{} , err
	}
	return decoded,nil
}

func (userusecase *UseCaseData)CreateUser(model domain.User) (domain.User, error) {
	user := model
	hasshedPasskey,err := bcrypt.GenerateFromPassword([]byte(user.Password) , bcrypt.DefaultCost);
	if err != nil {
		return domain.User{},err
	}
	user.Password = string(hasshedPasskey)
	id,err := userusecase.Repo.InsertUserDocument(user)

	if err!= nil {
		return domain.User{} , err
	}
	new_model,err := userusecase.GetUser(id)
	if err != nil {
		return domain.User{},err
	}
	
	return new_model,nil
}

func (userusecase *UseCaseData)UpdateUser(id string , model domain.User) (domain.User, error) {
	err := userusecase.Repo.UpdateUserDocumentById(id , model)
	if err != nil {
		return domain.User{},err
	}
	return model , nil
}

func (userusecase *UseCaseData)DeleteUser(id string) error {
	err := userusecase.Repo.DeleteUserDocument(id)
	if err != nil {
		return err
	}
	return nil
}

func (userusecase *UseCaseData)LogIn(model domain.AuthUser) (domain.User , error) {
	filter := make(map[string]string)
	filter["username"] = model.UserName
	filter["email"] = model.Email
	filter["password"] = model.Password

	result,err := userusecase.Repo.GetUserDocumentByFilter(filter)
	if err != nil {
		return domain.User{} , err
	}
	user := result[0]
	err = bcrypt.CompareHashAndPassword([]byte(model.Password) , []byte(user.Password))
	if err != nil {
		return domain.User{} , err
	}

	return user,nil
}


func (usecase *UseCaseData)Promote(id string)(domain.User , error) {
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