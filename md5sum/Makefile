md5sum:
	go build -o md5sum md5sum.go

clean:
	rm -f md5sum
	rm -f testfiles/output_test*_program*

test:
	cd testfiles && ../../testing/compare_ouput_argument_cases
