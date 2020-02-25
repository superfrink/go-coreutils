dirname:
	go build -o dirname dirname.go

clean:
	rm -f dirname
	rm -f testfiles/output_test*_program*

test:
	cd testfiles && ../../testing/compare_ouput_argument_cases
