package domain

type Repository_interface interface {
	GetTaskDocumentById(id string) (Task, error)
	GetTaskDocumentByFilter(filter map[string]string) ([]Task, error)
	UpdateTaskDocumentById(id string, update Task) error
	InsertTaskDocument(object Task) (string, error)
	DeleteTaskDocument(id string) error
	GetUserDocumentById(id string) (User , error)
	GetUserDocumentByFilter(filter map[string]string)([]User , error)
	UpdateUserDocumentById(id string , update User)  error
	InsertUserDocument(object User) (string, error)
	DeleteUserDocument(id string) error
}

type Usecase_interface interface {
	GetTask(id string) (Task, error)
	GetTasks(filter map[string]string) ([]Task, error)
	CreateTask(model Task) (Task, error)
	UpdateTask(id string, model Task) (Task, error)
	DeleteTask(id string) error
	GetUser(id string) (User, error)
	GetUsers() ([]User, error)
	CreateUser(model User) (User, error)
	UpdateUser(id string , model User) (User, error)
	DeleteUser(id string) error
	LogIn(model AuthUser) (User , error)
	Promote(id string)(User , error)
}

