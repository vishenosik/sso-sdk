package api

import (
	"context"
	"net/http"

	// pkg
	"github.com/go-chi/chi/v5"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"

	// internal pkg

	"github.com/vishenosik/gocherry/pkg/errors"
	_http "github.com/vishenosik/gocherry/pkg/http"

	//internal
	authentication_v1 "github.com/vishenosik/sso-sdk/gen/grpc/v1/authentication"
)

type User struct {
	ID       string `json:"id"`
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Authentication interface {
	//
	LoginByEmail(
		ctx context.Context,
		email,
		password,
		appID string,
	) (string, error)
	//
	RegisterUser(
		ctx context.Context,
		user *User,
	) (string, error)
	//
	IsAdmin(
		ctx context.Context,
		userID string,
	) (bool, error)
}

type AuthenticationApi struct {
	authentication_v1.UnimplementedAuthenticationServer
	auth Authentication
}

type authapi = AuthenticationApi

func NewAuthenticationApi(auth Authentication) *authapi {
	return &authapi{
		auth: auth,
	}
}

// ping godoc
//
//	@Summary 	Регистрация пользователя
//	@Tags 		system
//	@Router 	/api/ping [get]
//	@Produce 	html
//	@Success 	200 {string}  string    "ok"
//	@Failure 	406 {string}  string    "not ok"
func (a *authapi) Routers(r chi.Router) {

	r.Group(func(r chi.Router) {
		r.Route(a.registerUser())
	})
}

func (a *authapi) RegisterService(server *grpc.Server) {
	authentication_v1.RegisterAuthenticationServer(server, a)
}

/*
HTTP handlers
*/

func (a *authapi) registerUser() (string, func(chi.Router)) {

	versionMiddleware, versionHandler := _http.DotVersionMiddlewareHandler("1.0")

	return routeAuth("register"), func(r chi.Router) {
		r.Use(
			versionMiddleware,
		)
		r.Post(_http.BlankRoute, versionHandler(_http.HandlersMap{
			"1.0": a.registerUser_1_0(),
		}))
	}
}

func (a *authapi) registerUser_1_0() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

var routeAuth = _http.MethodFunc("auth")

/*
gRPC handlers
*/

// IsAdmin Checks if the user with the given ID is an admin.
//
// Parameters:
//
//	ctx: The context of the request.
//	userID: The ID of the user to check.
//
// Returns:
//
//	isAdmin: A boolean value indicating if the user is an admin.
//	err: An error if an issue occurs during the check.
func (a *authapi) IsAdmin(
	ctx context.Context,
	request *authentication_v1.IsAdminRequest,
) (*authentication_v1.IsAdminResponse, error) {

	const message = "admin check failed"
	fail := errors.FailWrapErrorStatus((*authentication_v1.IsAdminResponse)(nil), message)

	errmp := grpcErrorsMap(map[error]codes.Code{})

	// log := logger.With(
	// 	attrs.Operation(authentication_v1.Authentication_IsAdmin_FullMethodName),
	// 	attrs.UserID(request.GetUserId()),
	// )

	isAdmin, err := a.auth.IsAdmin(ctx, request.GetUserId())
	if err != nil {
		// log.Error(message, attrs.Error(err))
		return fail(errmp.Get(err))
	}

	return &authentication_v1.IsAdminResponse{
		IsAdmin: isAdmin,
	}, nil

}

// Login Logs the user in using the provided credentials.
//
// Parameters:
//
//	ctx: The context of the request.
//	request: The login request containing the user credentials.
//	appID: The ID of the application making the request.
//
// Returns:
//
//	token: string representing user token.
//	err: error if an issue occurs during the login process.
func (a *authapi) Login(
	ctx context.Context,
	request *authentication_v1.LoginRequest,
) (*authentication_v1.LoginResponse, error) {

	// serviceRequest := entities.LoginRequest{
	// 	Email:    request.GetEmail(),
	// 	Password: request.GetPassword(),
	// }

	// token, err := srv.auth.Login(ctx, serviceRequest, request.GetAppId())
	// if err != nil {
	// 	log.Error(message, attrs.Error(err))
	// 	return fail(entities.ServiceErrorsToGrpcCodes.Get(err))
	// }

	return &authentication_v1.LoginResponse{
		Token: "token",
	}, nil

}

// Register Registers a new user in the system.
//
// Parameters:
//
//	ctx: The context of the request.
//	request: The registration request containing the user information.
//
// Returns:
//
//	userID: A string representing the user ID.
//	err: An error if an issue occurs during the registration process.
func (a *authapi) Register(
	ctx context.Context,
	request *authentication_v1.RegisterRequest,
) (*authentication_v1.RegisterResponse, error) {

	// serviceRequest := models.RegisterRequest{
	// 	Nickname: "me",
	// 	Email:    request.GetEmail(),
	// 	Password: request.GetPassword(),
	// }

	// userID, err := srv.auth.RegisterNewUser(ctx, serviceRequest)
	// if err != nil {
	// 	log.Error(message, attrs.Error(err))
	// 	return fail(models.ServiceErrorsToGrpcCodes.Get(err))
	// }

	return &authentication_v1.RegisterResponse{
		UserId: "userID",
	}, nil

}
