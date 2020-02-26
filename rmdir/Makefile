rmdir: rmdir.go
	go build -o rmdir rmdir.go

clean:
	rm -f rmdir
	rm -f testfiles/output_test*_program*

test:
	cd testfiles && ../../testing/identitcal_ouput_tests
