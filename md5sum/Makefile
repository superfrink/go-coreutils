md5sum:
	go build -o md5sum md5sum.go

clean:
	rm -f md5sum
	rm -f testfiles/output_test*_program*

test:
	cd testfiles && ../../testing/identitcal_ouput_tests
