exit:
	go build -o exit exit.go

clean:
	rm -f exit
	rm -f testfiles/output_test*_program*

test:
	cd testfiles && ../../testing/compare_ouput_argument_cases
