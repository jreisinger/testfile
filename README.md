`tesfile` module creates (and optionally removes) a temporary file with a given content:

```
import "github.com/jreisinger/testfile"

func main() {
    	content = []byte("line1\nline2\n")
    	tf := testfile.New(content)

	testFile, err := os.Open(tf.Name)
	if err != nil {
		log.Fatal(err)
	}
	defer testFile.Close()
	defer tf.Remove()
    
    	// do something with testFile ...
}
```

If you don't `Remove()` the test file you can find it in [TempDir](https://golang.org/pkg/os/#TempDir) (usually `/tmp` on Unix systems) starting with `testfile`.
