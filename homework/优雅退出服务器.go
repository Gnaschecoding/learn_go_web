package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// Server 代表一个 HTTP 服务器，一个实例监听一个端口
type Server struct {
	httpServer *http.Server
}

// NewServer 创建一个新的 Server 实例
func NewServer(name, addr string) *Server {
	mux := &serverMux{
		reject:   false,
		ServeMux: http.NewServeMux(),
	}
	return &Server{
		httpServer: &http.Server{
			Addr:    addr,
			Handler: mux,
		},
	}
}

// Handle 注册一个处理函数到 Server 的路由中
func (s *Server) Handle(pattern string, handler http.Handler) {
	s.httpServer.Handler.(*serverMux).Handle(pattern, handler)
}

// serverMux 既可以看做是装饰器模式，也可以看做委托模式
type serverMux struct {
	reject bool
	*http.ServeMux
}

func (s *serverMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if s.reject {
		w.WriteHeader(http.StatusServiceUnavailable)
		_, _ = w.Write([]byte("服务已关闭"))
		return
	}
	s.ServeMux.ServeHTTP(w, r)
}

// App 代表应用本身
type App struct {
	servers []*Server
	cbs     []ShutdownCallback
	timeout time.Duration
}

// ShutdownCallback 定义退出回调函数类型
type ShutdownCallback func(ctx context.Context)

// Option 定义选项函数类型
type Option func(*App)

// WithShutdownCallbacks 用于注册退出回调
func WithShutdownCallbacks(cbs ...ShutdownCallback) Option {
	return func(app *App) {
		app.cbs = cbs
	}
}

// WithTimeout 用于设置优雅退出的超时时间
func WithTimeout(timeout time.Duration) Option {
	return func(app *App) {
		app.timeout = timeout
	}
}

// NewApp 创建一个新的 App 实例
func NewApp(servers []*Server, opts ...Option) *App {
	app := &App{
		servers: servers,
		timeout: 5 * time.Second, // 默认超时时间为 5 秒
	}
	for _, opt := range opts {
		opt(app)
	}
	return app
}

// StartAndServe 启动应用并监听系统信号
func (a *App) StartAndServe() {
	// 启动所有服务器
	for _, s := range a.servers {
		go func(server *Server) {
			fmt.Printf("Starting server on %s\n", server.httpServer.Addr)
			if err := server.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				fmt.Printf("Failed to start server on %s: %v\n", server.httpServer.Addr, err)
			}
		}(s)
	}

	// 监听系统信号
	c := make(chan os.Signal, 1)
	signals := []os.Signal{
		os.Interrupt,
		os.Kill,
		syscall.SIGKILL,
		//		syscall.SIGSTOP,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGQUIT,
		syscall.SIGILL,
		syscall.SIGTRAP,
		syscall.SIGABRT,
		//	syscall.SIGSYS,
		syscall.SIGTERM,
	}
	signal.Notify(c, signals...)

	// 第一次收到信号，开始优雅退出
	<-c
	fmt.Println("Received shutdown signal, starting graceful shutdown...")
	a.gracefulShutdown()

	// 再次监听信号，超时控制
	go func() {
		select {
		case <-c:
			fmt.Println("Received second shutdown signal, force exiting...")
			os.Exit(1)
		case <-time.After(a.timeout):
			fmt.Println("Graceful shutdown timed out, force exiting...")
			os.Exit(1)
		}
	}()
}

// gracefulShutdown 执行优雅退出操作
func (a *App) gracefulShutdown() {
	// 拒绝新请求
	for _, s := range a.servers {
		s.httpServer.Handler.(*serverMux).reject = true
	}

	// 等待已有请求执行完毕
	fmt.Printf("Waiting for %s to complete existing requests...\n", a.timeout)
	time.Sleep(a.timeout)

	// 关闭所有服务器
	ctx, cancel := context.WithTimeout(context.Background(), a.timeout)
	defer cancel()
	for _, s := range a.servers {
		if err := s.httpServer.Shutdown(ctx); err != nil {
			fmt.Printf("Failed to shutdown server on %s: %v\n", s.httpServer.Addr, err)
		}
	}

	// 执行退出回调
	var wg sync.WaitGroup
	for _, cb := range a.cbs {
		wg.Add(1)
		go func(callback ShutdownCallback) {
			defer wg.Done()
			ctx, cancel := context.WithTimeout(context.Background(), a.timeout)
			defer cancel()
			callback(ctx)
		}(cb)
	}
	wg.Wait()

	fmt.Println("Graceful shutdown completed.")
}

// StoreCacheToDBCallback 示例退出回调函数
func StoreCacheToDBCallback(ctx context.Context) {
	fmt.Println("Storing cache to database...")
	// 模拟耗时操作
	select {
	case <-time.After(2 * time.Second):
		fmt.Println("Cache stored to database.")
	case <-ctx.Done():
		fmt.Println("Store cache to database timed out.")
	}
}

func main() {
	s1 := NewServer("business", "localhost:8080")
	s1.Handle("/", http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		_, _ = writer.Write([]byte("hello"))
	}))
	s2 := NewServer("admin", "localhost:8081")
	app := NewApp([]*Server{s1, s2}, WithShutdownCallbacks(StoreCacheToDBCallback), WithTimeout(10*time.Second))
	app.StartAndServe()
}
