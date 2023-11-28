package main

import (
	
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)


/*
	你被要求编写一个叫做 WebsiteRacer 的函数，用来对比请求两个 URL 来「比赛」
	返回先响应的 URL。如果两个 URL 在 10 秒内都未返回结果，那么应该返回一个 error。
*/
func TestRacer(t *testing.T) {
	// t.Run("", func(t *testing.T) {
	// 	slowServer := makeDelayedServer(20 * time.Millisecond)
	// 	fastServer := makeDelayedServer(0 * time.Millisecond)

	// 	defer slowServer.Close()
	// 	defer fastServer.Close()

	// 	slowURL := slowServer.URL
	// 	fastURL := fastServer.URL

	// 	fmt.Println(slowURL)
	// 	fmt.Println(fastURL)

	// 	got := RacerA(slowURL, fastURL)

	// 	want := fastURL

	// 	if got != want {
	// 		t.Errorf("got [%s] want [%s]", got, want)
	// 	}
	// })

	// t.Run("", func(t *testing.T) {
	// 	serverA := makeDelayedServer(11 * time.Second)
	// 	serverB := makeDelayedServer(12 * time.Second)

	// 	defer serverA.Close()
	// 	defer serverB.Close()

	// 	_, err := RacerB(serverA.URL, serverB.URL)

	// 	if err == nil {
	// 		t.Error("expected an error but didn't get one")
	// 	}
	// })

	// t.Run("", func(t *testing.T) {
	// 	slowServer := makeDelayedServer(20 * time.Millisecond)
	// 	fastServer := makeDelayedServer(0 * time.Millisecond)

	// 	defer slowServer.Close()
	// 	defer fastServer.Close()

	// 	slowURL := slowServer.URL
	// 	fastURL := fastServer.URL

	// 	want := fastURL
	// 	got, err := RacerC(slowURL, fastURL)

	// 	if err != nil {
	// 		t.Fatalf("did not expect an error but got one %v", err)
	// 	}

	// 	if got != want {
	// 		t.Errorf("got [%s] want [%s]", got, want)
	// 	}
	// })

	t.Run("returns an error if a server doesn't respond within 10s", func(t *testing.T) {
		server := makeDelayedServer(25 * time.Millisecond)

		defer server.Close()

		_, err := ConfigurableRacer(server.URL, server.URL, 20 *time.Millisecond)

		if err == nil {
			t.Error("expected an error but didn't get one")
		}
	})
}

func makeDelayedServer(t time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(t)
		w.WriteHeader(http.StatusOK)
	}))
}
