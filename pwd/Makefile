pwd:
	go build -o pwd pwd.go

clean:
	rm -f pwd
	rm -f testfiles/output_test*_program*

test:
	cd testfiles && ../../testing/compare_ouput_argument_cases
