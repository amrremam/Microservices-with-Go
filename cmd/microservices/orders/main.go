package orders

import (
	"log"
	"net/http"
	"os"
)

func main() {
	log.Println("Starting orders microservice")

	ctx := cmd.Context()

	r, closeFn := createOrdersMicroservice()
	defer closeFn()

	server := &http.Server{Addr: os.Getenv("SHOP_ORDERS_SERVICE_BIND_ADDR"), Handler: r}
	go func() {
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			panic(err)
		}
	}()

	<-ctx.Done()
	log.Println("Closing orders microservice")

	if err := server.Close(); err != nil {
		panic(err)
	}
}


func createOrderMicroservice() (router *chi.Mux, closeFn func()) {
	cmd.WaitForService(os.Getenv("SHOP_RABBITMQ_ADDR"))
	shopHTTPClient := orders_infra_product.NewHTTPClient(os.Getenv("SHOP_SERVICE_ADDR"))

	r := cmd.CreateRouter()

	orders_public_http.AddRoutes(r, ordersService, ordersRepo)
	orders_private_http.AddRoutes(r, ordersService, ordersRepo)

	return r, func(){

	}
}
