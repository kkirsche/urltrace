// Copyright © 2016 Kevin Kirsche <kevin.kirsche@verizon.com> <kev.kirsche@gmail.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

var (
	timeout int
	fullURL bool
)

// TransportWrapper wraps the http.Transport structure to allow us to record the
// URLs which we are redirected through
type TransportWrapper struct {
	*http.Transport
}

// RoundTrip executes a single HTTP transaction, returning
// a Response for the provided Request.
func (t *TransportWrapper) RoundTrip(req *http.Request) (*http.Response, error) {
	// Use the default transport we set
	transport := t.Transport

	if transport == nil {
		transport = http.DefaultTransport.(*http.Transport)
	}

	resp, err := transport.RoundTrip(req)
	if err != nil {
		return resp, err
	}

	// Log the status code and the URL used
	if fullURL {
		log.Printf("Status: %d, Full URL: %s\n", resp.StatusCode, req.URL.String())
	} else {
		log.Printf("Status: %d, Base URL: %s\n", resp.StatusCode, req.URL.Host)
	}

	return resp, err
}

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "urltrace",
	Short: "urltrace allows a user to trace a URL's redirects",
	Long: `urltrace is designed to allow a user to trace the redirect path of a
URL and record that so that they can identify any URLs which are necessary to
reach a given URL. The command may be used like so:`,
	Run: func(cmd *cobra.Command, args []string) {
		log.SetPrefix("[URL Tracer] ")

		log.Println("Creating Transport Wrapper")
		t := &TransportWrapper{
			Transport: http.DefaultTransport.(*http.Transport),
		}

		log.Printf("Creating HTTP Client with %d second timeout\n", timeout)
		timeoutString := strconv.Itoa(timeout)
		timeoutDuration, err := time.ParseDuration(timeoutString + "s")
		if err != nil {
			log.Panicln(err)
		}

		client := &http.Client{
			Transport: t,
			Timeout:   timeoutDuration * time.Second,
		}

		_, err = client.Get(strings.Join(args, ""))
		if err != nil {
			log.Panicln(err)
		}
	},
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports Persistent Flags, which, if defined here,
	// will be global for your application.

	RootCmd.PersistentFlags().BoolVarP(&fullURL, "full-url", "f", false, "Display the entire URL, not the host portion.")
	RootCmd.PersistentFlags().IntVarP(&timeout, "timeout", "t", 10, "Sets the timeout in seconds for a requested URL")
}
