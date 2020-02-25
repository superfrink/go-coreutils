testing:
	go build -o testing testing.go

clean:
	rm -f testing
	rm -f testfiles/output_test*_program*

test:
	cd testfiles && ../../testing/compare_ouput_argument_cases
