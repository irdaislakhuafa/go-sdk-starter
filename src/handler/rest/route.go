package rest

func (r *rest) RegisterRoutes() {
	// server health and testing purpose
	r.svr.Get("/ping", r.Ping)

	api := r.svr.Group("api")
	v1 := api.Group("/v1")
	{
		v1.Post("/todos", r.CreateTodo)
		v1.Get("/todos", r.ListTodo)
	}
}
