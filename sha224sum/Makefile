sha224sum:
	go build -o sha224sum sha224sum.go

clean:
	rm -f sha224sum
	rm -f testfiles/output_test*_program*

test:
	cd testfiles && ../../testing/compare_ouput_argument_cases
