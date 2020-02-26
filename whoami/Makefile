whoami: whoami.go
	go build -o whoami whoami.go

clean:
	rm -f whoami
	rm -f testfiles/output_test*_program*

test:
	cd testfiles && ../../testing/identitcal_ouput_tests
