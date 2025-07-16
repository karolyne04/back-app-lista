package api

import (
	"net/http"

	"karol/internal/interfaces/api/handler"
	"karol/internal/interfaces/api/middleware"
)

// Router gerencia as rotas da API
type Router struct {
	authHandler    *handler.AuthHandler
	userHandler    *handler.UserHandler
	authMiddleware *middleware.AuthMiddleware

	shoppingListHandler *handler.ShoppingListHandler
	listItemHandler     *handler.ListItemHandler
	categoryHandler     *handler.CategoryHandler
	tipHandler          *handler.TipHandler
}

// NewRouter cria uma nova instância de Router
func NewRouter(
	authHandler *handler.AuthHandler,
	userHandler *handler.UserHandler,
	authMiddleware *middleware.AuthMiddleware,
	shoppingListHandler *handler.ShoppingListHandler,
	listItemHandler *handler.ListItemHandler,
	categoryHandler *handler.CategoryHandler,
	tipHandler *handler.TipHandler,
) *Router {
	return &Router{
		authHandler:         authHandler,
		userHandler:         userHandler,
		authMiddleware:      authMiddleware,
		shoppingListHandler: shoppingListHandler,
		listItemHandler:     listItemHandler,
		categoryHandler:     categoryHandler,
		tipHandler:          tipHandler,
	}
}

// Setup configura as rotas da API
func (r *Router) Setup() http.Handler {
	mux := http.NewServeMux()

	// Rotas públicas
	mux.HandleFunc("/api/auth/login", r.authHandler.Login)
	mux.HandleFunc("/api/users", r.userHandler.Create)
	mux.HandleFunc("/api/categories", r.categoryHandler.List)
	mux.HandleFunc("/api/tips", r.tipHandler.List)

	// Rotas protegidas
	protectedMux := http.NewServeMux()
	protectedMux.HandleFunc("/api/users/me", r.userHandler.GetProfile)
	// ShoppingList
	protectedMux.HandleFunc("/api/lists", r.shoppingListHandler.List) // GET
	protectedMux.HandleFunc("/api/lists", r.shoppingListHandler.Create) // POST
	protectedMux.HandleFunc("/api/lists/finish", r.shoppingListHandler.Finish) // POST (exemplo)
	protectedMux.HandleFunc("/api/lists/update", r.shoppingListHandler.Update) // PUT
	protectedMux.HandleFunc("/api/lists/delete", r.shoppingListHandler.Delete) // DELETE
	// ListItem
	protectedMux.HandleFunc("/api/lists/items/add", r.listItemHandler.Add) // POST
	protectedMux.HandleFunc("/api/lists/items/update", r.listItemHandler.Update) // PUT
	protectedMux.HandleFunc("/api/lists/items/remove", r.listItemHandler.Remove) // DELETE
	protectedMux.HandleFunc("/api/lists/items/check", r.listItemHandler.Check) // POST

	// Aplicar middleware de autenticação às rotas protegidas
	mux.Handle("/api/users/me", r.authMiddleware.Authenticate(protectedMux))
	mux.Handle("/api/lists", r.authMiddleware.Authenticate(protectedMux))
	mux.Handle("/api/lists/finish", r.authMiddleware.Authenticate(protectedMux))
	mux.Handle("/api/lists/update", r.authMiddleware.Authenticate(protectedMux))
	mux.Handle("/api/lists/delete", r.authMiddleware.Authenticate(protectedMux))
	mux.Handle("/api/lists/items/add", r.authMiddleware.Authenticate(protectedMux))
	mux.Handle("/api/lists/items/update", r.authMiddleware.Authenticate(protectedMux))
	mux.Handle("/api/lists/items/remove", r.authMiddleware.Authenticate(protectedMux))
	mux.Handle("/api/lists/items/check", r.authMiddleware.Authenticate(protectedMux))

	return mux
}