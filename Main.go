package main

//only need FMT to start
import ("fmt"
        "time"
        "log"
        "strings"
        "io/ioutil"
)

func trackTime(start time.Time, name string) {

  elapsed := time.Since(start)
  log.Printf("%s took %s", name, elapsed)
}

func Perm(a []rune, f func([]rune)) {
	perm(a, f, 0)
}

func perm(a []rune, f func([]rune), i int) {

	if i > len(a) {
		f(a)
		return
	}
	perm(a, f, i+1)

	for j := i + 1; j < len(a); j++ {

		a[i], a[j] = a[j], a[i]
		perm(a, f, i+1)
		a[i], a[j] = a[j], a[i]

	}
}

func main() {

  //track program time
  start := time.Now()

  Perm([]rune("time"), func(a []rune) {

    //get test string
  	var word = strings.ToLower(string(a))

    // read the whole file at once
    b, err := ioutil.ReadFile("words.txt")
    if err != nil {
        panic(err)
    }

    s := string(b)

    if strings.Contains(s, word) == true {

      s := strings.Fields(s)
      //grab word from list
      for _, WORD := range s {
  		    if strings.HasSuffix(WORD, word) {
            if strings.HasPrefix(WORD, word) {
              fmt.Println(word)
              break
            }
          }
  	  }
    }

  })

  trackTime(start, "\nUnscrambling this word")
}
