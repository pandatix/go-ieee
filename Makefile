.PHONY: tests
tests:
	@echo "--- Integration tests ---"
	go test ./api -run=^Test_I_ -coverprofile=integration.out -json | tee -a gotest.json

.PHONY: clean
clean:
	rm gotest.json integration.out
