uptime: uptime.go
	go build -o uptime uptime.go

clean:
	rm -f uptime
	rm -f testfiles/output_test*_program*

test:
	cd testfiles && ../../testing/identitcal_ouput_tests
