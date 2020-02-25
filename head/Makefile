head:
	go build -o head head.go

clean:
	rm -f head
	rm -f testfiles/output_test*_program*

test:
	cd testfiles && ../../testing/compare_ouput_argument_cases
