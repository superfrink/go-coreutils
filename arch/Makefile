arch: arch.go
	go build -o arch arch.go

clean:
	rm -f arch
	rm -f testfiles/output_test*_program*

test:
	cd testfiles && ../../testing/identitcal_ouput_tests
