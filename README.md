# file-eraser
This randomize files before removing.<br>
If it failed to randomize, the file would not be removed.<br>
Multithread.<br>
Its randomness depends on crypto/rand.Reader in Go standard library.<br>
If there are duplicate paths, it processed once.
<br>
<br>

# Installation
	go install github.com/lsejx/file-eraser@latest
<br><br>

# Option
	-h	help
	-r	recursive (for directory)
	-i	interactive (confirm before erasing)
	-k	keep (randomize, but don't remove)
<br><br>


# Usage
	file-eraser -h
	file-eraser file1
	file-eraser -r file1 dir1
	file-eraser -ri dir1
	file-eraser -k file1
