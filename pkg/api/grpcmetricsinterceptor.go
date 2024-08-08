package api

import googleGrpc "google.golang.org/grpc"

type Interceptor struct {
}

func (i *Interceptor) Intercept(srv any, ss googleGrpc.ServerStream, info *googleGrpc.StreamServerInfo, handler googleGrpc.StreamHandler) error {

}
