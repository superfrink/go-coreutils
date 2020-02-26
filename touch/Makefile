touch: touch.go
	go build -o touch touch.go

clean:
	rm -f touch
	rm -f testfiles/output_test*_program*

test:
	cd testfiles && ../../testing/identitcal_ouput_tests
