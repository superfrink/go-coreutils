ls:
	go build -o ls ls.go

clean:
	rm -f ls
	rm -f testfiles/output_test*_program*

test:
	cd testfiles && ../../testing/compare_ouput_argument_cases
