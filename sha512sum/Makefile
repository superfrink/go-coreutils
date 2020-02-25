sha512sum:
	go build -o sha512sum sha512sum.go

clean:
	rm -f sha512sum
	rm -f testfiles/output_test*_program*

test:
	cd testfiles && ../../testing/compare_ouput_argument_cases
