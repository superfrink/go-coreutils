uname:
	go build -o uname uname.go

clean:
	rm -f uname
	rm -f testfiles/output_test*_program*

test:
	cd testfiles && ../../testing/compare_ouput_argument_cases
