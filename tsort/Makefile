tsort:
	go build -o tsort tsort.go

clean:
	rm -f tsort
	rm -f testfiles/output_test*_program*

test:
	cd testfiles && ../../testing/compare_ouput_argument_cases
