sha256sum: sha256sum.go
	go build -o sha256sum sha256sum.go

clean:
	rm -f sha256sum
	rm -f testfiles/output_test*_program*

test:
	cd testfiles && ../../testing/identitcal_ouput_tests
