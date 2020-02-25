stat:
	go build -o stat stat.go

clean:
	rm -f stat
	rm -f testfiles/output_test*_program*

test:
	cd testfiles && ../../testing/compare_ouput_argument_cases
