whoami:
	go build -o whoami whoami.go

clean:
	rm -f whoami
	rm -f testfiles/output_test*_program*

test:
	cd testfiles && ../../testing/compare_ouput_argument_cases
