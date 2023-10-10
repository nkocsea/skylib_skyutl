wrun:
	go env -w GOPRIVATE=suntech.com.vn/skylib/skylog.git
	go run src/main.go

run:
	go env -w GOPRIVATE=suntech.com.vn/skylib/skylog.git
	reflex -r '\.go' -s -- sh -c "go run src/main.go"

tidy:
	go clean -modcache
	rm -Rf go.sum
	go env -w GOPRIVATE=suntech.com.vn/skylib/skylog.git
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