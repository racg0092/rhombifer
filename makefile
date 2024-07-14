



test_parsing: ./pkg/parsing/tests/parsing_test.go 
	go test -v ./pkg/parsing/tests/



test_rhombi:
ifdef silent
		go test ./tests/ 
else
		go test -v ./tests/
endif
