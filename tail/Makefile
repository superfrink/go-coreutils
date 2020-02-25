tail:
	go build -o tail tail.go

clean:
	rm -f tail
	rm -f testfiles/output_test*_program*

test:
	cd testfiles && ../../testing/compare_ouput_argument_cases
