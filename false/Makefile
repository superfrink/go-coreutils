false:
	go build -o false false.go

clean:
	rm -f false
	rm -f testfiles/output_test*_program*

test:
	cd testfiles && ../../testing/compare_ouput_argument_cases
