include .env

runserver:
	go run cmd/main.go
	
github-push:
	@echo "pushing to GitHub..."
	@echo "" 
	@git push https://${GITHUB_USERNAME}:${GITHUB_PAT}@github.com/${GITHUB_USERNAME}/${GITHUB_REPO_NAME} master
	@echo ""

register-user:
	curl -X POST 'http://localhost:5000/users/register/' -H "Content-Type: application/json" -d '{"username": "sam", "password": "abc123"}'

login-user:
	curl -X POST 'http://localhost:5000/users/login/' -H "Content-Type: application/json" -d '{"username": "sam", "password": "abc123"}'
	