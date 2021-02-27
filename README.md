fetch dependencies from go.mod :
- go mod download

run server :
- go run cmd/simple-service/main.go

curl create user :
curl --location --request POST 'http://127.0.0.1:1323/users' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "tri"
}'

curl get user :
curl --location --request GET 'http://127.0.0.1:1323/users/1'

curl update user :
curl --location --request PUT 'http://127.0.0.1:1323/users/1' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "tri budiyono"
}'

curl delete user :
curl --location --request DELETE 'http://127.0.0.1:1323/users/1'