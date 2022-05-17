create-migration:
	go run github.com/prisma/prisma-client-go migrate dev --name init
migrate:
	go run github.com/prisma/prisma-client-go migrate dev
db-gen:
	go run github.com/prisma/prisma-client-go generate
gql-gen:
	go run github.com/99designs/gqlgen generate
serve:
	go run server.go