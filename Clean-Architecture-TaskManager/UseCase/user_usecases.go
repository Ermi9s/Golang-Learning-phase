package usecase

import (
	"github.com/Ermi9s.Golang-Learning-phase/Clean-Architecture-TaskManager/domain"
)


func (userusecase *UseCaseData)GetUser(id string) (domain.Model, error) {
	doc , err := userusecase.repo.GetDocumentById("Users" , id)
	if err != nil {
		return &domain.Task{},err
	}
	return doc, nil
}

func (userusecase *UseCaseData)GetUsers() ([]domain.User, error) {
	filter := make(map[string]string)
	decoded, err := userusecase.repo.GetDocumentByFilter("Users" , filter)
	if err != nil {
		return []domain.User{} , err
	}
	result := []domain.User{}
	for _,val := range decoded {
		new := val.(*domain.User)
		result = append(result, *new)
	}
	return result,nil
}

func (userusecase *UseCaseData)CreateUser(model domain.Model) (domain.Model, error) {
	err := userusecase.repo.InsertDocument("Users" , model)
	if err != nil {
		return &domain.User{} , err
	}
	return model,nil
}

func (userusecase *UseCaseData)UpdateUser(id string , model domain.Model) (domain.Model, error) {
	err := userusecase.repo.UpdateDocumentById("Users", id , model)
	if err != nil {
		return &domain.User{},err
	}
	return model , nil
}

func (userusecase *UseCaseData)DeleteUser(id string) error {
	err := userusecase.repo.DeleteDocument("Users" , id)
	if err != nil {
		return err
	}
	return nil
}

func (userusecase *UseCaseData)LogIn(model domain.AuthUser) (domain.Model , error) {
	filter := make(map[string]string)
	filter["username"] = model.UserName
	filter["email"] = model.Email
	filter["password"] = model.Password

	result,err := userusecase.repo.GetDocumentByFilter("Users" , filter)
	user := result[0]
	if err != nil {
		return &domain.User{} , err
	}
	return user , nil
}
func (usecase *UseCaseData)Promote(id string)(domain.User , error) {
	iuser , err := usecase.GetUser(id)
	if err != nil {
		return domain.User{},err
	}
	user := iuser.(*domain.User)
	user.Is_admin = true

	iuser,err = usecase.UpdateUser(id , user)
	if err != nil {
		return domain.User{},err
	}

	return *user , nil
}