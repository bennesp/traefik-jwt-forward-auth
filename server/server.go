package server

import (
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/bennesp/traefik-jwt-forward-auth/services/jwt"
	"github.com/bennesp/traefik-jwt-forward-auth/sources"
	"github.com/gin-gonic/gin"
	ginlogrus "github.com/toorop/gin-logrus"
)

type Server struct {
	source     sources.Source
	jwtWrapper *jwt.JwtWrapper
}

func New(source sources.Source, jwtWrapper *jwt.JwtWrapper) *Server {
	return &Server{
		source:     source,
		jwtWrapper: jwtWrapper,
	}
}

func (ctx *Server) Start(address string) {
	log.Infof("Starting server on port %s...", address)

	logger := log.New()
	r := gin.New()
	r.Use(ginlogrus.Logger(logger), gin.Recovery())
	r.GET("/", ctx.handleGet)

	http.ListenAndServe(address, r)
}
