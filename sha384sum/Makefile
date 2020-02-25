sha384sum:
	go build -o sha384sum sha384sum.go

clean:
	rm -f sha384sum
	rm -f testfiles/output_test*_program*

test:
	cd testfiles && ../../testing/compare_ouput_argument_cases
