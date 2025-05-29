include .env

runserver:
	go run cmd/main.go
	
github-push:
	@echo "pushing to GitHub..."
	@echo "" 
	@git push https://${GITHUB_USERNAME}:${GITHUB_PAT}@github.com/${GITHUB_USERNAME}/${GITHUB_REPO_NAME} master
	@echo ""

register-user:
	curl -X POST 'http://localhost:5000/users/register/' -H "Content-Type: application/json" -d '{"username": "sa123123m", "password": "abc123"}'

login-user:
	curl -X POST 'http://localhost:5000/users/login/' -H "Content-Type: application/json" -d '{"username": "sam", "password": "abc123", "userToken": "X6HPDdXXUf"}'

create-message:
	curl -X POST 'http://localhost:5000/messages/create/' -H "Content-Type: application/json" -d '{"chat_id": "66566c4e9f70b053ad1e8451", "sender_id": "66566c4e9f70b053ad1e8453", "receiver_id": "66566c4e9f70b053ad1e8452", "text_content": "hello world"}'
	
create-chat:
	curl -X POST 'http://localhost:5000/chats/create/' -H "Content-Type: application/json" -d '{"participants": ["665755ae6fb2a8e4e4d3a0c1", "665755ae6fb2a8e4e4d3a0c2"]}'
	
delete-chat:
	curl -X DELETE 'http://localhost:5000/chats/delete/' -H "Content-Type: application/json" -d '{"chat_id": "683829b83e03d29feaacf67a"}'