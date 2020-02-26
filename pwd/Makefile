pwd: pwd.go
	go build -o pwd pwd.go

clean:
	rm -f pwd
	rm -f testfiles/output_test*_program*

test:
	cd testfiles && ../../testing/identitcal_ouput_tests
