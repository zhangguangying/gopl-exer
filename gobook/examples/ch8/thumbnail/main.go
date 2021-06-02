package thumbnail

import (
	"gopl.io/ch8/thumbnail"
	"log"
)

func makeThumbnails(filenames []string)  {
	for _, f := range filenames {
		if _, err := thumbnail.ImageFile(f); err != nil {
			log.Println(err)
		}
	}
}

func makeThumbnails2(filenames []string) {
	for _, f := range filenames {
		go thumbnail.ImageFile(f)
	}
}

func makeThumbnails3(filenames []string) {
	ch := make(chan struct{})
	for _, f := range filenames {
		go func(f string) {
			thumbnail.ImageFile(f)
			ch <- struct{}{}
		}(f)
	}
	for range filenames {
		<- ch
	}
}