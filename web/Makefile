.PHONY: css-watch
css-watch:
	tailwindcss -i css/input.css -o assets/css/styles.css --minify --watch

.PHONY: dev
dev:
	templ generate
	tailwindcss -i css/input.css -o assets/css/styles.css --minify
	go run .
