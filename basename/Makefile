basename: basename.go
	go build -o basename basename.go

clean:
	rm -f basename
	rm -f testfiles/output_test*_program*

test:
	cd testfiles && ../../testing/identitcal_ouput_tests
