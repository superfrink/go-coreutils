touch:
	go build -o touch touch.go

clean:
	rm -f touch
	rm -f testfiles/output_test*_program*

test:
	cd testfiles && ../../testing/compare_ouput_argument_cases
