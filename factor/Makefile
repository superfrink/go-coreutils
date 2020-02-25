factor:
	go build -o factor factor.go

clean:
	rm -f factor
	rm -f testfiles/output_test*_program*

test:
	cd testfiles && ../../testing/compare_ouput_argument_cases
