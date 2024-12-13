server:
	cd backend && go run .

client:
	cd frontend && npm run dev

.PHONY: server client
