deps:
	echo "Dependencies"
	go mod download
	echo "Dependencies Completed"
compile:
	echo "Build Process for Windows, Linux and MacOS"
	GOOS=windows GOARCH=amd64 go build -o bin/windows/main.exe main.go
	GOOS=linux GOARCH=amd64 go build -o bin/linux/main main.go
	GOOS=darwin GOARCH=amd64 go build -o bin/macos/main main.go
	echo "Build Process Completed"
run:
	echo "Run Process"
	./bin/
	echo "Run Process Completed"
run_windows:
	echo "Running Windows"
	./bin/windows/main.exe
run_linux:
	echo "Running Linux"
	./bin/linux/main
run_macos:
	echo "Running MacOS"
	./bin/macos/main
stop:
	echo "Stop Process"
	killall go run main.go
	echo "Stop Process Completed"
vendor:
	echo "Vendor Process"
	go mod vendor
	echo "Vendor Process Completed"
clean:
	echo "Clean Process"
	go clean
	rm -rf bin/
	echo "Clean Process Completed"

