# A little bit about Go Kit

Isn't Go just a perfect language for microservices? Peter Bourgon certainly thought so when he started a project that was supposed to make it popular for this stuff as well. After a few months it grew into a pretty cool and interesting tool for Go developers. Let's see what you need to do to use it and how it can make your work easier.

### What is Go Kit?

Some say that this is a framework, some that it's a set of good practices. I'd say it's a bit of both, and more. It all starts with understanding, that your service is not only the business logic that does all the magic for you, but also logging, monitoring, message transport (incoming and outgoing), etc. On Go Kit website it's described as an onion ([source](https://gokit.io/faq/)):

<img src="https://raw.githubusercontent.com/mycodesmells/gokit-example/master/res/onion.png"/>

The thing is, that you should never reinvent the wheel (in software development and in life in general!). Make use of things that exist already, because you not only save your time coming up with a solution, but also save yourself from making errors that some people have already overcame. Another important aspect is consistency - once you have some solution in place and start using it with all services, you have one common way of dealing with similar problems. This might not seem important for three-service environment, but once you have dozens of them, you'd understand how time-consuming it might get.

### Service example

Our example service is just as simple as it can get. We will have an HTTP server that takes a year of birth as an input parameter and calculate current age based on that:

    http.HandleFunc("/age", func(w http.ResponseWriter, req *http.Request) {
        ... // decode request
        age, _ := as.CalculateAge(ageRequest.YearOfBirth)
        ... // encode response
    })
    fmt.Printf("%v", http.ListenAndServe(":8000", nil))
 
While the business logic looks like this:
 
    type ageService struct{}
    
    func (ageService) CalculateAge(yearOfBirth int) (int, error) {
    	year := time.Now().Year()
    
    	if yearOfBirth > year {
    		return 0, errNotBornYet
    	}
    
    	return year - yearOfBirth, nil
    }
    
### Adding Go Kit

Go Kit is based on endpoints (`endpoint.Endpoint`), which wrap your business logic and take care of request decoding and reqsponse encoding. Plus, this encoding code can be easily shared and used in different endpoints (services):

    func makeCalculateAgeEndpoint(as AgeService) endpoint.Endpoint {
    	return func(ctx context.Context, request interface{}) (interface{}, error) {
    		req := request.(calculateAgeRequest)
    		age, err := as.CalculateAge(req.YearOfBirth)
    		if err != nil {
    			return calculateAgeResponse{age, err.Error()}, nil
    		}
    		return calculateAgeResponse{age, ""}, nil
    	}
    }
    
Our update `main` function would look like this:

    func main() {
    	ctx := context.Background()
    	as := ageService{}
    
    	ageHandler := httptransport.NewServer(
    		ctx,
    		makeCalculateAgeEndpoint(as),
    		decodeAgeRequest,
    		encodeResponse,
    	)
    
    	http.Handle("/age", ageHandler)
    	log.Fatal(http.ListenAndServe(":8001", nil))
    }

But when we run our new server, we don't feel any difference. Our service is still not logging anything, so it's far from being perfect. Let's do something about it with some basic logging.
  
### Basic logging

The fastest way to add logging is add an aspect-like logger, that is to log something before calling a method, calling it and adding another log after the call ends:
 
    type middleware func(endpoint.Endpoint) endpoint.Endpoint
    
    func loggingMiddleware(logger log.Logger) middleware {
    	return func(next endpoint.Endpoint) endpoint.Endpoint {
    		return func(ctx context.Context, request interface{}) (interface{}, error) {
    			logger.Log("msg", "calling endpoint")
    			defer logger.Log("msg", "called endpoint")
    			return next(ctx, request)
    		}
    	}
    }
    
This also is a bit similar to Node JS middleware pattern, with additional stuff happening after the `next(..)` call finishes. Once we start our server and make a request we get some logs:

    $ ageStep2 // starting the server
    msg=HTTP addr=:8002
    method=CalculateAge msg="calling endpoint"
    method=CalculateAge msg="called endpoint"

It's an improvement, but not a major one. We still don't know anything about the details of our code, why stop now?

### Enhanced logging

In order to make our logs more specific, we need to create a custom-shaped logger that will call our service specifically. What we gain from this, is an access to input and output data so that they can be logged and provide more useful information. 

    type loggingMiddleware struct {
    	logger log.Logger
    	next   AgeService
    }
    
    func (mw loggingMiddleware) CalculateAge(y int) (output int, err error) {
    	defer func(begin time.Time) {
    		mw.logger.Log(
    			"method", "calculateAge",
    			"input", y,
    			"output", output,
    			"err", err,
    			"took", time.Since(begin),
    		)
    	}(time.Now())
    
    	output, err = mw.next.CalculateAge(y)
    	return
    }

With the addition of some call duration details and all of the sudden you get a pretty decent-looking logs:

    $ ageStep3
    msg=HTTP addr=:8003
    method=calculateAge input=1960 output=56 err=null took=63.527µs
    
### Summary

With Go-Kit and middleware-based processing you have an ability to add verious functionalities around your sevice, such as logging, monitoring or circuit breaking (turning the service off if for some reason it doesn't behave as it supposed to). Plus, you are not limited to one way of data transport (in this example I'm using JSON, but you can quickly change it to eg. gRPC). Sky is the limit here, as more and more tools are supported (open source is awesome!). I highly recommend you to create your services with Go Kit, as it might save a ton of work/time in a long run.

Source code of this example is available [on Github](https://github.com/mycodesmells/gokit-example).
