yes:
	go build -o yes yes.go

clean:
	rm -f yes
	rm -f testfiles/output_test*_program*

test:
	cd testfiles && ../../testing/compare_ouput_argument_cases
