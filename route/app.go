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
}

func App() AppModel {

	//domain
	healthDomain := &domain.HealthDomainCtx{}
	authDomain := &domain.AuthDomainCtx{}
	userDomain := &domain.UserDomainCtx{}
	roleDomain := &domain.RoleDomainCtx{}
	userListDomain := &domain.UserListDomainCtx{}

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

	return AppModel{
		Health:   healthHandler,
		User:     userHandler,
		Auth:     authHandler,
		Role:     roleHandler,
		UserList: userListHandler,
	}
}
