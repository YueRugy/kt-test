package main

import (
	"flag"
	"fmt"
	kitlog "github.com/go-kit/kit/log"
	tranhttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"golang.org/x/time/rate"
	. "kt-test/service"
	"kt-test/util"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

func main() {

	name := flag.String("name", "", "服务名称")
	port := flag.Int("p", 0, "端口")
	flag.Parse()
	//fmt.Println(*name)
	//fmt.Println(*port)
	if *name == "" {
		log.Fatal("请输入name")
	}
	if *port == 0 {
		log.Fatal("请输入port")
	}
	util.SetNameAndPort(*name, *port)

	var logger kitlog.Logger
	{
		logger = kitlog.NewLogfmtLogger(os.Stdout)
		logger = kitlog.WithPrefix(logger, "kit-test", "kit-1.0")
		logger = kitlog.With(logger, "time", kitlog.DefaultTimestampUTC)
		logger = kitlog.With(logger, "caller", kitlog.DefaultCaller)
	}
	us := UserService{}
	l := rate.NewLimiter(1, 3)
	e := RateLimit(l)(UserLogger(logger)(GetEndpoint(&us)))
	//e := service.GetEndpoint(&us)
	options := []tranhttp.ServerOption{
		tranhttp.ServerErrorEncoder(MyErrorEncoder),
	}
	hand := tranhttp.NewServer(e, DecodeUserRequest, EncodeUserResponse, options...)
	r := mux.NewRouter()
	{
		r.Methods("GET", "DELETE").Path(`/user/{uid:\d+}`).Handler(hand)
		r.Methods("GET").Path("/health").HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			writer.Header().Set("Content-type", "application/json")
			_, _ = writer.Write([]byte(`{"status":"ok"}`))
		})
	}
	r.Handle(`/user/{uid:\d+}`, hand)
	errChan := make(chan error)
	go func() {
		util.Register()
		err := http.ListenAndServe(":"+strconv.Itoa(*port), r)
		if err != nil {
			errChan <- err
		}
	}()
	go func() {
		sig_c := make(chan os.Signal)
		signal.Notify(sig_c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-sig_c)
	}()
	getErr := <-errChan
	util.Deregister()
	log.Fatal(getErr)
}
