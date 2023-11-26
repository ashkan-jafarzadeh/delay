package rest

func (s *Server) SetupRoutes() {
	api := s.app.Group("/api/v1")

	api.Post("delay/:orderId", s.Delay)
	api.Put("assign", s.Assign)
	api.Get("vendor/report", s.VendorReport)
}
