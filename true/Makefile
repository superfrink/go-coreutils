true:
	go build -o true true.go

clean:
	rm -f true
	rm -f testfiles/output_test*_program*

test:
	cd testfiles && ../../testing/compare_ouput_argument_cases
