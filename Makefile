.PHONY: run stop clean

run:
	go mod download
	[ -f .env ] || cp .env.example .env
	docker-compose up -d
	air

stop:
	docker-compose down

clean:
	docker-compose down -v
	rm -rf tmp
