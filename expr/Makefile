expr:
	go build -o expr expr.go

clean:
	rm -f expr
	rm -f testfiles/output_test*_program*

test:
	cd testfiles && ../../testing/compare_ouput_argument_cases
