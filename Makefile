wrun:
	go env -w GOPRIVATE=github.com/nkocsea/skylib_skylog
	go run src/main.go

run:
	go env -w GOPRIVATE=github.com/nkocsea/skylib_skylog
	reflex -r '\.go' -s -- sh -c "go run src/main.go"

tidy:
	go clean -modcache
	rm -Rf go.sum
	go env -w GOPRIVATE=github.com/nkocsea/skylib_skylog
	go mod tidy

tags:
	git ls-remote --tags

commit:
	git status
	git add .
	git commit -m"$m"
	git push
	git tag $t
	git push --tags