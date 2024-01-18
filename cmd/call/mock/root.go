package mock

import (
	"io"
	"log"
	"net/http"
	"os"
	"svix-poc/lib/app/cmd"
	"sync"

	"github.com/samber/lo"
	"github.com/spf13/cobra"
)

var Root = cmd.New("mock", cmd.Add(
	cmd.New("reset", cmd.Alias("r"), cmd.Run(runResetMocks)),
	cmd.New("list", cmd.Alias("l"), cmd.Run(runListMocks)),
	cmd.New("create", cmd.Alias("c"), cmd.Run(runCreateMocks)),
))

func runResetMocks(_ *cobra.Command, _ []string) {
	res := lo.Must(http.Post("http://localhost:3000/reset", "", nil))
	defer res.Body.Close()
	log.Println(string(lo.Must(io.ReadAll(res.Body))))
}

func runListMocks(_ *cobra.Command, _ []string) {
	res := lo.Must(http.Get("http://localhost:3000/mocks"))
	defer res.Body.Close()
	log.Println(string(lo.Must(io.ReadAll(res.Body))))
}

func runCreateMocks(_ *cobra.Command, files []string) {
	wg := sync.WaitGroup{}
	wg.Add(len(files))
	for _, file := range files {
		go func(file string) {
			defer wg.Done()
			if f, err := os.Open(file); err != nil {
				log.Println(err)
			} else if res, err := http.Post("http://localhost:3000/mocks", "application/x-yaml", f); err != nil {
				log.Println(err)
			} else {
				log.Println(string(lo.Must(io.ReadAll(res.Body))))
			}
		}(file)
	}
	wg.Wait()
}
