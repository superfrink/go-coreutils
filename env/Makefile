env:
	go build -o env env.go

clean:
	rm -f env
	rm -f testfiles/output_test*_program*

test:
	cd testfiles && ../../testing/compare_ouput_argument_cases
