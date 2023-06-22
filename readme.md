# Supabase Go

This project make an POST and GET to Supabase API.

Run the app:

```bash
air
```

## Endpoints

### POST /users
Post an user
```bash
curl --request POST \
  --url http://localhost:8000/user \
  --header 'Content-Type: application/json' \
  --data '{
	"name": "juan",
	"email": "a@.co",
	"nascimento": "vhvhjv",
	"telefone": "31099699"
}'
```
### GET /users
Get all users
```bash
curl --request GET \
  --url http://localhost:8000/user
```