



test_parsing: ./pkg/parsing/tests/parsing_test.go 
	go test -v ./pkg/parsing/tests/

builtin_test:
	go test -v ./pkg/builtin

test_rhombi:
ifdef silent
		go test ./tests/ 
else
		go test -v ./tests/
endif



