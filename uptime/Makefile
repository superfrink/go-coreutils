uptime:
	go build -o uptime uptime.go

clean:
	rm -f uptime
	rm -f testfiles/output_test*_program*

test:
	cd testfiles && ../../testing/compare_ouput_argument_cases
