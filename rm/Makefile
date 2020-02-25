rm:
	go build -o rm rm.go

clean:
	rm -f rm
	rm -f testfiles/output_test*_program*

test:
	cd testfiles && ../../testing/compare_ouput_argument_cases
