rm: rm.go
	go build -o rm rm.go

clean:
	rm -f rm
	rm -f testfiles/output_test*_program*

test:
	cd testfiles && ../../testing/identitcal_ouput_tests
