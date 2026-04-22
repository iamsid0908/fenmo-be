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
}

func App() AppModel {

	//domain
	healthDomain := &domain.HealthDomainCtx{}
	authDomain := &domain.AuthDomainCtx{}
	userDomain := &domain.UserDomainCtx{}
	roleDomain := &domain.RoleDomainCtx{}
	userListDomain := &domain.UserListDomainCtx{}
	expenseDomain := &domain.ExpenseDomainCtx{}

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

	return AppModel{
		Health:   healthHandler,
		User:     userHandler,
		Auth:     authHandler,
		Role:     roleHandler,
		UserList: userListHandler,
		Expense:  expenseHandler,
	}
}
