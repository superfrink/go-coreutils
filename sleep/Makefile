sleep:
	go build -o sleep sleep.go

clean:
	rm -f sleep
	rm -f testfiles/output_test*_program*

test:
	cd testfiles && ../../testing/compare_ouput_argument_cases
