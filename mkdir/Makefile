mkdir:
	go build -o mkdir mkdir.go

clean:
	rm -f mkdir
	rm -f testfiles/output_test*_program*

test:
	cd testfiles && ../../testing/compare_ouput_argument_cases
