package cmd

import (
	"log"
	"net/http"

	"fmt"

	"github.com/spf13/cobra"
)

var servCmd = &cobra.Command{
	Use:   "serv",
	Short: "Start Ftp Server",
	Long:  `Start ftp server in current or selected directory`,
	Run:   servCmdFunc,
}

var (
	username  string
	password  string
	directory string

	port    int
	address string
)

func servCmdFunc(cmd *cobra.Command, args []string) {

	addr := "0.0.0.0:9090"
	if port != 9090 || address != "0.0.0.0" {
		addr = fmt.Sprintf("%s:%d", address, port)
	}

	http.Handle("/", handleAuth(http.FileServer(http.Dir(directory))))
	log.Printf("Serving %s on : %s:%d\n", directory, address, port)
	log.Fatal(http.ListenAndServe(addr, nil))
}

func handleAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if username != "" && password != "" {
			w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
			u, p, authOK := r.BasicAuth()
			if authOK == false {
				http.Error(w, "Not authorized", 401)
				return
			}
			if u != username || p != password {
				http.Error(w, "Not authorized", 401)
				return
			}
		}
		next.ServeHTTP(w, r)
	})
}

func init() {
	servCmd.Flags().StringVarP(&username, "username", "u", "", "username for auth")
	servCmd.Flags().StringVarP(&password, "password", "p", "", "password for auth")
	servCmd.Flags().IntVarP(&port, "port", "P", 9090, "port number to listen on")
	servCmd.Flags().StringVarP(&address, "ip", "i", "0.0.0.0", "Ip address to listen on")
	servCmd.Flags().StringVarP(&directory, "directory", "d", ".", "Source directory to read from")
	rootCmd.AddCommand(servCmd)
}
