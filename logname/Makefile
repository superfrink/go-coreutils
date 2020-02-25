logname:
	go build -o logname logname.go

clean:
	rm -f logname
	rm -f testfiles/output_test*_program*

test:
	cd testfiles && ../../testing/compare_ouput_argument_cases
