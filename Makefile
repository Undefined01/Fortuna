build: backend frontend

.PHONY: backend frontend test

backend:
	make -C ./backend

frontend:
	make -C ./frontend

test:
	make -C ./backend backend_test
