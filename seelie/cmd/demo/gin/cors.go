package gin

import (
	"time"

	"github.com/gin-contrib/cors"
)

// Cors ...
func Cors() {
	v := cors.New(cors.Config{
		AllowOrigins:     []string{"*.mihoyo.com", "*.hoyoverse.com"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowHeaders:     []string{"Access-Control-Allow-Origin", "Content-Type", "X-Authorization-Token", "Accept", "Origin"},
		AllowCredentials: true,
		MaxAge:           time.Second * time.Duration(3600),
		AllowWildcard:    true,
	})
	_ = v
	time.Sleep(time.Hour)
}
