type ServerHooks struct {
	// RequestReceived is called as soon as a request enters the Twirp
	// server at the earliest available moment.
	RequestReceived func(context.Context) (context.Context, error)

	// RequestRouted is called when a request has been routed to a
	// particular method of the Twirp server.
	RequestRouted func(context.Context) (context.Context, error)

	// ResponsePrepared is called when a request has been handled and a
	// response is ready to be sent to the client.
	ResponsePrepared func(context.Context) context.Context

	// ResponseSent is called when all bytes of a response (including an error
	// response) have been written. Because the ResponseSent hook is terminal, it
	// does not return a context.
	ResponseSent func(context.Context)

	// Error hook is called when an error occurs while handling a request. The
	// Error is passed as argument to the hook.
	Error func(context.Context, Error) context.Context
}