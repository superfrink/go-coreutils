sha256sum:
	go build -o sha256sum sha256sum.go

clean:
	rm -f sha256sum
	rm -f testfiles/output_test*_program*

test:
	cd testfiles && ../../testing/compare_ouput_argument_cases
