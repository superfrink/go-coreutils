wc: wc.go
	go build -o wc wc.go

clean:
	rm -f wc
	rm -f testfiles/output_test*_program*

test:
	cd testfiles && ../../testing/identitcal_ouput_tests
