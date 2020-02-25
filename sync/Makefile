sync:
	go build -o sync sync.go

clean:
	rm -f sync
	rm -f testfiles/output_test*_program*

test:
	cd testfiles && ../../testing/compare_ouput_argument_cases
