cat: cat.go
	go build -o cat cat.go

clean:
	rm -f cat
	rm -f testfiles/output_test*_program*

test:
	cd testfiles && ../../testing/identitcal_ouput_tests
