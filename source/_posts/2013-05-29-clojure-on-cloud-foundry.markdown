---
layout: post
title: "Clojure on Cloud Foundry"
date: 2013-05-29 22:10
slug: clojure-on-cloud-foundry
comments: true
status: publish
categories: 
- cloudfoundry
- clojure
---

I was inspired by Brian McClain's [post on bringing Haskell to Cloud Foundry](http://catdevrandom.me/blog/2013/05/16/buildpacks-in-cloud-foundry-v2/) using Cloud Foundry v2 buildpacks, so I decided to go on a buildpack journey of my own. Since Clojure is the language I most enjoying "toying around with," I thought I'd try to deploy a simple Clojure web application using the [Heroku Clojure Buildpack](https://github.com/mstine/heroku-buildpack-clojure).

To reiterate some of the coolness around buildpacks, they are what allows a PaaS like Cloud Foundry or Heroku to support various runtimes without first building that support into the core platform. If your favorite language or framework runtime isn't available, there's absolutely nothing stopping you from providing your own buildpack to support it. Stuff can get crazy -- McClain has even [hinted at bringing FORTRAN support to Cloud Foundry](https://twitter.com/BrianMMcClain/status/339479905974751232).

I decided for my experiment to build a VERY basic "Hello World" style application using [Ring](https://github.com/ring-clojure/ring), which is "a Clojure web applications library inspired by Python's WSGI and Ruby's Rack." The easiest way to get started building Ring applications is to utilize the popular Clojure build tool [Leiningen](http://github.com/technomancy/leiningen).

First I started by creating a new project:

``` bash
$ lein new hello-cf
$ cd hello-cf
``` 

The next task was to add the Ring dependencies to my `project.clj` file:

``` clojure
(defproject hello-cf "0.1.0-SNAPSHOT"
  :min-lein-version "2.0.0"
  :description "Hello Clojure on Cloud Foundry!"
  :license {:name "Eclipse Public License"
            :url "http://www.eclipse.org/legal/epl-v10.html"}
  :dependencies [[org.clojure/clojure "1.4.0"]
                 [ring/ring-core "1.1.8"]
                 [ring/ring-jetty-adapter "1.1.8"]])
```

Then it was time to create the Ring application itself, by editing `src/hello_cf/core.clj`:

``` clojure
(ns hello-cf.core
  (:use [ring.adapter.jetty :only [run-jetty]]))

(defn handler [request]
  {:status 200
   :headers {"Content-Type" "text/html"}
   :body "Hello Cloud Foundry from heroku-buildpack-clojure!"})

(defn -main [port]
  (run-jetty handler {:port (Integer. port)}))
```

Let's break this down a bit. The `handler` function will handle any HTTP request that hits our application, and will return an "OK" response containing a pleasant message indicating that we've succeeded. That's really about it. Our application is complete. We can test it out by running the following:

``` bash
$ lein trampoline run -m hello-cf.core 8080
2013-05-29 22:42:52.576:INFO:oejs.Server:jetty-7.6.1.v20120215
2013-05-29 22:42:52.804:INFO:oejs.AbstractConnector:Started SelectChannelConnector@0.0.0.0:8080
```

Hitting `http://localhost:8080` in the browser confirms that we're well on our way. Now it's time to trying pushing the application to Cloud Foundry. As Brian stated in his blog, one of the stellar aspects of Cloud Foundry buildpacks is that they are approximately the same as Heroku buildpacks. Practically, this means that one *should* be able to utilize a Heroku buildpack on Cloud Foundry with minimal or no modifications. Let's put that theory to the test, shall we? Before we do, let's create a `Procfile` quickly to let the buildpack know what we want to run:

```
web: lein with-profile production trampoline run -m hello-cf.core $PORT
```

And on with the push:

``` bash
cf push hello-cf --buildpack=git://github.com/heroku/heroku-buildpack-clojure.git
Using manifest file manifest.yml

Creating hello-cf... OK

1: hello-cf
2: none
Subdomain> hello-cf

1: mstine.cf-app.com
2: none
Domain> mstine.cf-app.com

Binding hello-cf.mstine.cf-app.com to hello-cf... OK
Uploading hello-cf... OK
Starting hello-cf... OK
-----> Downloaded app package (12K)
Initialized empty Git repository in /tmp/buildpacks/heroku-buildpack-clojure.git/.git/
Installing heroku-buildpack-clojure.git.
-----> Installing OpenJDK 1.6...done
-----> Installing Leiningen
       Downloading: leiningen-2.1.2-standalone.jar
       Writing: lein script
-----> Building with Leiningen
       Running: lein with-profile production compile :all
       Retrieving lein-standalone-repl/lein-standalone-repl/0.1.5/lein-standalone-repl-0.1.5.pom from clojars
       Retrieving lein-standalone-repl/lein-standalone-repl/0.1.5/lein-standalone-repl-0.1.5.jar from clojars
       Performing task 'compile' with profile(s): 'production'
       Retrieving org/clojure/clojure/1.4.0/clojure-1.4.0.pom from
       ...
       Compiling hello-cf.core
-----> Uploading staged droplet (66M)
-----> Uploaded droplet
Checking hello-cf...
Staging in progress...
Staging in progress...
Staging in progress...
Staging in progress...
Staging in progress...
Staging in progress...
Staging in progress...
  0/1 instances: 1 starting
  0/1 instances: 1 down
  0/1 instances: 1 starting
  0/1 instances: 1 flapping
Application failed to start.
```

Drat. Let's take a quick look at the logs to see what may be awry:

``` bash
Reading logs/stderr.log... OK
/home/vcap/app/.lein/bin/lein: line 42: java: command not found
```

Ah-hah! Looks like the existing buildpack is making some assumptions about the structure of our application that no longer hold true on Cloud Foundry. So, I followed in Brian's footsteps and [forked away](https://github.com/mstine/heroku-buildpack-clojure). One small [commit](https://github.com/mstine/heroku-buildpack-clojure/commit/fd2c46cc23267fa2d808123d2fd58f4295da4b85) looks like it ought to fix the problem. Let's give it another try:

``` bash
cf push hello-cf --buildpack=git://github.com/mstine/heroku-buildpack-clojure.git
Using manifest file manifest.yml

Not applying manifest changes without --reset
See `cf diff` for more details.

Uploading hello-cf... OK
Changes:
  buildpack: 'git://github.com/heroku/heroku-buildpack-clojure.git' -> 'git://github.com/mstine/heroku-buildpack-clojure.git'
Updating hello-cf... OK
Stopping hello-cf... OK

Starting hello-cf... OK
-----> Downloaded app package (8.0K)
-----> Downloaded app buildpack cache (17M)
Initialized empty Git repository in /tmp/buildpacks/heroku-buildpack-clojure.git/.git/
Installing heroku-buildpack-clojure.git.
-----> Installing OpenJDK 1.6...done
-----> Using cached Leiningen 2.1.2
       Writing: lein script
-----> Building with Leiningen
       Running: lein with-profile production compile :all
       Performing task 'compile' with profile(s): 'production'
       Compiling hello-cf.core
-----> Uploading staged droplet (66M)
-----> Uploaded droplet
Checking hello-cf...
Staging in progress...
Staging in progress...
Staging in progress...
Staging in progress...
Staging in progress...
Staging in progress...
Staging in progress...
Staging in progress...
Staging in progress...
Staging in progress...
Staging in progress...
  0/1 instances: 1 starting
  0/1 instances: 1 starting
  0/1 instances: 1 starting
  0/1 instances: 1 starting
  0/1 instances: 1 starting
  0/1 instances: 1 starting
  0/1 instances: 1 starting
  0/1 instances: 1 starting
  0/1 instances: 1 starting
  0/1 instances: 1 starting
  0/1 instances: 1 starting
  0/1 instances: 1 starting
  0/1 instances: 1 starting
  0/1 instances: 1 starting
  0/1 instances: 1 starting
  0/1 instances: 1 starting
  0/1 instances: 1 starting
  1/1 instances: 1 running
OK
```

BOOM!

I quickly pointed my browser, and as yesterday's tweet indicates, success:

<blockquote class="twitter-tweet" data-partner="tweetdeck"><p>Yeah...that just happened. cc: <a href="https://twitter.com/cloudfoundry">@cloudfoundry</a> <a href="https://twitter.com/search?q=%23clojure&amp;src=hash">#clojure</a> <a href="https://twitter.com/search?q=%23buildpacks&amp;src=hash">#buildpacks</a> <a href="https://twitter.com/search?q=%23winning&amp;src=hash">#winning</a> <a href="http://t.co/9lJHqmWQPw">pic.twitter.com/9lJHqmWQPw</a></p>&mdash; Matt Stine (@mstine) <a href="https://twitter.com/mstine/statuses/339248683151417344">May 28, 2013</a></blockquote>
<script async src="//platform.twitter.com/widgets.js" charset="utf-8"></script>  

Score another win for Cloud Foundry's buildpack support. I'm now toying with the idea of doing something of a world tour of LISP on Cloud Foundry. My next candidate may be [Scheme](https://github.com/evhan/heroku-buildpack-chicken). 
