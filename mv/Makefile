mv:
	go build -o mv mv.go

clean:
	rm -f mv
	rm -f testfiles/output_test*_program*

test:
	cd testfiles && ../../testing/compare_ouput_argument_cases
