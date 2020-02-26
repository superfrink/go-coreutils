stat: stat.go
	go build -o stat stat.go

clean:
	rm -f stat
	rm -f testfiles/output_test*_program*

test:
	cd testfiles && ../../testing/identitcal_ouput_tests
