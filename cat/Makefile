cat:
	go build -o cat cat.go

clean:
	rm -f cat
	rm -f testfiles/output_test*_program*

test:
	cd testfiles && ../../testing/compare_ouput_argument_cases
