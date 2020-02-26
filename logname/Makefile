logname: logname.go
	go build -o logname logname.go

clean:
	rm -f logname
	rm -f testfiles/output_test*_program*

test:
	cd testfiles && ../../testing/identitcal_ouput_tests
