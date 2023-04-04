package routes

import (
	"product-crud/handler"
	"product-crud/mock"
	"testing"
)

func TestNewRoutes(t *testing.T) {
	mockService := &mock.ServiceMock{}

	todoHandler := handler.NewProductHandler(mockService)
	got := NewRoutes(todoHandler)

	expectedRoutes := []string{
		"/products",
		"/products/:id",
		"/products",
		"/products/:id",
		"/products/:id",
	}

	t.Run("all routes are present", func(t *testing.T) {
		for i, v := range got {
			if v.Pattern != expectedRoutes[i] {
				t.Errorf("pattern expectations mismatched: \n want: %v \n got: %v", v.Pattern, expectedRoutes[i])
			}
		}
	})
}
