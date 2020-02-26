uname: uname.go
	go build -o uname uname.go

clean:
	rm -f uname
	rm -f testfiles/output_test*_program*

test:
	cd testfiles && ../../testing/identitcal_ouput_tests
