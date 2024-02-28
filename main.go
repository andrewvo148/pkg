package main

import (
	"compress/gzip"
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"sort"
	"time"
)

type Order struct {
	ID          string    `json:"id"`
	DateOrdered time.Time `json:"date_ordered"`
	CustomerID  string
	Items       []Item
}
type Item struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

type HelloHandler struct{}

func (hh HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello!\n"))
}

func main() {
	people := []Person{
		{"Pat", "Patterson", 37},
		{"Tracy", "Bobbert", 23},
		{"Fred", "Fredson", 18},
	}
	fmt.Println(people)
	sort.Slice(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})
	fmt.Println(people)

	twoBase := makeMult(2)
	threeBase := makeMult(3)
	for i := 0; i < 3; i++ {
		fmt.Println(twoBase(i), threeBase(i))
	}
}

type Counter struct {
	total       int
	lastUpdated time.Time
}

func (c *Counter) Increment() {
	c.total++
	c.lastUpdated = time.Now()
}

func (c *Counter) String() string {
	return fmt.Sprintf("total: %d, last updated: %v", c.total, c.lastUpdated)
}
func countLetters(r io.Reader) (map[string]int, error) {
	buf := make([]byte, 2048)
	out := map[string]int{}
	for {
		n, err := r.Read(buf)
		for _, b := range buf[:n] {
			if (b >= 'A' && b <= 'Z') || (b >= 'a' && b <= 'z') {
				out[string(b)]++
			}
		}
		if err == io.EOF {
			return out, nil
		}
		if err != nil {
			return nil, err
		}
	}
}

func buildGzipReader(fileName string) (*gzip.Reader, func(), error) {
	r, err := os.Open(fileName)
	if err != nil {
		return nil, nil, err
	}
	gr, err := gzip.NewReader(r)
	if err != nil {
		return nil, nil, err
	}

	return gr, func() {
		gr.Close()
		r.Close()
	}, nil
}

func longRunningThingManager(ctx context.Context, data string) (string, error) {
	type wrapper struct {
		result string
		err    error
	}
	ch := make(chan wrapper, 1)
	go func() {
		// do the long running thing
		result, err := longRunningThing(ctx, data)
		ch <- wrapper{
			result: result,
			err:    err,
		}
	}()
	select {
	case data := <-ch:
		return data.result, data.err
	case <-ctx.Done():
		return "", ctx.Err()
	}
}

func longRunningThing(ctx context.Context, data string) (string, error) {
	time.Sleep(10 * time.Second)
	return "", nil
}

func makeMult(base int) func(int) int {
	return func(factor int) int {
		return base * factor
	}
}
