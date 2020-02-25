rmdir:
	go build -o rmdir rmdir.go

clean:
	rm -f rmdir
	rm -f testfiles/output_test*_program*

test:
	cd testfiles && ../../testing/compare_ouput_argument_cases
