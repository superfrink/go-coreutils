sha1sum:
	go build -o sha1sum sha1sum.go

clean:
	rm -f sha1sum
	rm -f testfiles/output_test*_program*

test:
	cd testfiles && ../../testing/identitcal_ouput_tests
