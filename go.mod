module starter-kit

go 1.12

require (
	github.com/appleboy/gin-jwt/v2 v2.6.2
	github.com/gin-gonic/contrib v0.0.0-20190923054218-35076c1b2bea
	github.com/gin-gonic/gin v1.4.0
	github.com/go-resty/resty v0.0.0-00010101000000-000000000000
	github.com/go-stack/stack v1.8.0 // indirect
	github.com/google/go-cmp v0.4.0 // indirect
	github.com/nicksnyder/go-i18n/v2 v2.0.3
	github.com/sirupsen/logrus v1.4.2
	github.com/xdg/scram v0.0.0-20180814205039-7eeb5667e42c // indirect
	github.com/xdg/stringprep v1.0.0 // indirect
	go.mongodb.org/mongo-driver v1.1.1
	golang.org/x/net v0.0.0-20190628185345-da137c7871d7 // indirect
	golang.org/x/sync v0.0.0-20190911185100-cd5d95a43a6e // indirect
	golang.org/x/text v0.3.2
	gopkg.in/resty.v1 v1.12.0 // indirect
)

replace github.com/go-resty/resty => gopkg.in/resty.v1 v1.11.0
