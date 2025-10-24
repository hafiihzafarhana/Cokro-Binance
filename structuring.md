cokro-binance/
├── cmd/
│ └── app/
│ └── main.go
│
├── config/
│ ├── config.go
│ └── env.yaml
│
├── domain/
│ ├── spot/
│ │ ├── market/
│ │ │ ├── entity/
│ │ │ │ └── market.go
│ │ │ ├── repository.go
│ │ │ ├── service.go
│ │ │ └── error.go
│ │ └── order/
│ │ ├── entity/
│ │ │ └── order.go
│ │ ├── repository.go
│ │ ├── service.go
│ │ └── error.go
│ └── futures/...
│
├── internal/
│ └── spot/
│ ├── bootstrap/
│ │ └── spot_bootstrap.go
│ ├── market/
│ │ ├── repository/
│ │ │ └── impl.go
│ │ ├── usecase/
│ │ │ ├── usecase_impl.go // implementasi
│ │ │ └── mapper.go // mapper
│ │ │ └── usecase.go // interface
│ │ ├── controller.go
│ │ ├── router.go
│ │ └── dto/
│ │ ├── request.go
│ │ └── response.go
│ └── order/
│ └── ...
│
├── pkg/
│ ├── httpclient/
│ ├── logger/
│ ├── helper/
│ ├── utils/
│ ├── grpcclient/
│ └── websocketclient/
│
├── shared/
│ ├── constant/
│ ├── middleware/
│ ├── error/
│ └── response/
│
└── test/
├── mocks/
└── integration/
