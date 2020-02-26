sha224sum: sha224sum.go
	go build -o sha224sum sha224sum.go

clean:
	rm -f sha224sum
	rm -f testfiles/output_test*_program*

test:
	cd testfiles && ../../testing/identitcal_ouput_tests
