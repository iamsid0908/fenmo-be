package route

import (
	"core/domain"
	"core/handler"
	"core/service"
)

type AppModel struct {
	Health   handler.HealthHandler
	User     handler.UserHandler
	Auth     handler.AuthHandler
	Role     handler.RoleHandler
	UserList handler.UserListHandler
	Expense  handler.ExpenseHandler
	Category handler.CategoryHandler
}

func App() AppModel {

	//domain
	healthDomain := &domain.HealthDomainCtx{}
	authDomain := &domain.AuthDomainCtx{}
	userDomain := &domain.UserDomainCtx{}
	roleDomain := &domain.RoleDomainCtx{}
	userListDomain := &domain.UserListDomainCtx{}
	expenseDomain := &domain.ExpenseDomainCtx{}
	categoryDomain := &domain.CategoryDomainCtx{}

	//service
	healthService := service.HealthService{
		HealthDomain: healthDomain,
	}
	userService := service.UserService{
		UserDomain: userDomain,
	}
	authService := service.AuthService{
		AuthDomain: authDomain,
		UserDomain: userDomain,
	}
	roleService := service.RoleService{
		RoleDomain: roleDomain,
	}
	userListService := service.UserListService{
		UserListDomain: userListDomain,
	}
	expenseService := service.ExpenseService{
		ExpenseDomain: expenseDomain,
	}
	categoryService := service.CategoryService{
		CategoryDomain: categoryDomain,
	}
	//handler
	healthHandler := handler.HealthHandler{
		HealthService: healthService,
	}
	userHandler := handler.UserHandler{
		UserService: userService,
	}
	authHandler := handler.AuthHandler{
		AuthService: authService,
	}
	roleHandler := handler.RoleHandler{
		RoleService: roleService,
	}
	userListHandler := handler.UserListHandler{
		UserListService: userListService,
	}
	expenseHandler := handler.ExpenseHandler{
		ExpenseService: expenseService,
	}
	categoryHandler := handler.CategoryHandler{
		CategoryService: categoryService,
	}
	return AppModel{
		Health:   healthHandler,
		User:     userHandler,
		Auth:     authHandler,
		Role:     roleHandler,
		UserList: userListHandler,
		Expense:  expenseHandler,
		Category: categoryHandler,
	}
}
