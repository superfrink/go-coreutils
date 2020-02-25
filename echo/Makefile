echo:
	go build -o echo echo.go

clean:
	rm -f echo
	rm -f testfiles/output_test*_program*

test:
	cd testfiles && ../../testing/compare_ouput_argument_cases
