include .env

runserver:
	go run cmd/main.go
	
github-push:
	@echo "Enter the branch name (e.g. master, main, dev): "; \
	read branch; \
	echo "pushing to GitHub..."; \
	git push https://${GITHUB_USERNAME}:${GITHUB_PAT}@github.com/${GITHUB_USERNAME}/${GITHUB_REPO_NAME}.git $$branch; \ 
	echo ""

register-user:
	curl -X POST 'http://localhost:5000/users/register/' -H "Content-Type: application/json" -d '{"username": "sam", "password": "admin"}'

login-user:
	curl -X POST 'http://localhost:5000/users/login/' -H "Content-Type: application/json" -d '{"username": "sam", "password": "abc123", "userToken": "X6HPDdXXUf"}'
	
delete-message:
	curl -X DELETE 'http://localhost:5000/messages/delete/68384d27196d5b09876d9dbb' 

create-chat:
	curl -X POST 'http://localhost:5000/chats/create/' \
	-d '{"participants": ["68613234ea4da55bd33bccf5", "6861325aea4da55bd33bccf6"]}'
		
delete-chat:
	curl -X DELETE 'http://localhost:5000/chats/delete/683829b83e03d29feaacf67a' 
	
get-chat:
	curl -X POST 'http://localhost:5000/chats/get/' -H "Content-Type: application/json" -d '{"chat_id": "68382f5c0a2e91ac833b7a38"}'

get-user-chats:
	curl -X POST 'http://localhost:5000/users/chats/' -H "Content-Type: application/json" -d '{"user_id": "682722b817686ba22654c79f"}'
