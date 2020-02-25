wc:
	go build -o wc wc.go

clean:
	rm -f wc
	rm -f testfiles/output_test*_program*

test:
	cd testfiles && ../../testing/compare_ouput_argument_cases
