cokro-binance/
├── cmd/
│ └── app/
│ └── main.go
│
├── config/
│ └── config.go
│
├── domain/
│ └── spot/
│ └── general/
│ ├── entity/
│ │ └── general_entity.go
│ ├── repository.go // kontrak repository
│ └── service.go // kontrak usecase (service)
│
├── internal/
│ └── spot/
│ └── general/
│ ├── repository/
│ │ ├── model/
│ │ │ ├── general_model.go
│ │ │ └── mapper.go
│ │ └── repository_impl.go
│ │
│ ├── usecase/
│ │ ├── usecase.go // interface GeneralUsecase
│ │ └── usecase_impl.go // implementasi
│ │
│ ├── dto/
│ │ ├── request.go
│ │ └── response.go
│ │
│ ├── controller.go
│ └── router.go
│
├── pkg/
│ ├── httpclient/
│ │ └── client.go
│ ├── utils/
│ │ └── convert.go
│ └── logger/
│
└── shared/
├── error/
└── response/
