base64:
	go build -o base64 base64.go

clean:
	rm -f base64
	rm -f testfiles/output_test*_program*

test:
	cd testfiles && ../../testing/compare_ouput_argument_cases
