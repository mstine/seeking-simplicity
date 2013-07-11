---
layout: post
slug: blue-green-deployments-on-cloudfoundry
title: Blue-Green Deployments on Cloud Foundry
date: 2013-07-10 20:10
comments: true
status: publish
categories: 
- cloudfoundry
- cloud
---

One of the great things about Cloud Foundry is that it is a great enabler. Tall words. But what do they mean? Essentially, Cloud Foundry (and any other well-designed PaaS) enables us to do things as developers and operators that would be extremely difficult in a traditional deployment environments. One particularly valuable area of enablement is our new found ability to practice [Continous Delivery](http://continuousdelivery.com/), meaning that we continuously prove our ability to deliver working software by continuously treating each code commit to a system as if it could be deployed to a production environment. We do this by shipping these commits through what's called a "deployment pipeline," which consists of a series of build-test-deploy cycles that prove out a commit's suitability for production deployment. At the end of the pipeline we can either deploy automatically to our production environment (i.e. continuous deployment), or we can have a business decision to deploy that "deployable artifact" or not.

One particular practice associated with Continuous Delivery is [Blue-Green Deployments](http://martinfowler.com/bliki/BlueGreenDeployment.html). Martin Fowler describes these very well at the link provided, but I'll summarize briefly here:

* Cut-over to a new software version is tricky, and must be quick in order to minimize downtime events.
* Blue-green deployments ensure the parallel existence of two, identical (as possible) production environments.
* At any given point, only one (e.g. blue) services production traffic.
* New deploys are made to the other (e.g. green) environment. Final smoke testing is performed here.
* When green is determined ready, we begin routing traffic to it.
* We then stop routing traffic to blue.

Of course, there are several concerns that must be dealt with at the application level in order for this to work (datastores should support concurrent usage by two app versions, long running requests may be killed, etc.). What we'll focus on in this post is how Cloud Foundry supports the mechanics summarized above.

We'll begin with a basic Spring application named `ms-spr-demo`. This app takes users to a simple web page announcing the ubiquitous "Hello World!" message. We'll utilize the `cf` command-line interface to push the application:

``` bash
$ cf push --path build/libs/cf-demo.war
Name> ms-spr-demo

Instances> 1

Memory Limit> 512M

Creating ms-spr-demo... OK

1: ms-spr-demo
2: none
Subdomain> ms-spr-demo

1: cfapps.io
2: mattstine.com
3: none
Domain> 1

Creating route ms-spr-demo.cfapps.io... OK
Binding ms-spr-demo.cfapps.io to ms-spr-demo... OK

Create services for application?> n

Save configuration?> y

Saving to manifest.yml... OK
Uploading ms-spr-demo... OK
Starting ms-spr-demo... OK
-----> Downloaded app package (9.5M)
Installing java.
Downloading JDK...
Copying openjdk-1.7.0_25.tar.gz from the buildpack cache ...
Unpacking JDK to .jdk
Downloading Tomcat: apache-tomcat-7.0.41.tar.gz
Copying apache-tomcat-7.0.41.tar.gz from the buildpack cache ...
Unpacking Tomcat to .tomcat
Copying mysql-connector-java-5.1.12.jar from the buildpack cache ...
Copying postgresql-9.0-801.jdbc4.jar from the buildpack cache ...
Copying auto-reconfiguration-0.6.8.jar from the buildpack cache ...
-----> Uploading droplet (48M)
-----> Uploaded droplet
Checking ms-spr-demo...
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
  1/1 instances: 1 running
OK
```

The end result of this `cf push` event is that an application is now serving requests at `http://ms-spr-demo.cfapps.io`. The following graphic shows the state of our system, with the CF Router sending traffic to our application:

{% img center /images/blue-green/BlueGreen1.png %} 

Next, we make a slight change to our application. Rather that saying "Hello World!" we decide to make it say "Goodbye World!" We build a new war file for the application. Rather than letting `cf` prompt us this time, we'll make use of the `manifest.yml` file that we saved from our previous push. However, we'll rename the application and provide a new route. Take a look:

``` 
---
applications:
- name: ms-spr-demo-green
  memory: 512M
  instances: 1
  url: ms-spr-demo-green.cfapps.io
  path: build/libs/cf-demo.war
```

As you can see, we're calling our new application version `ms-spr-demo-green` and we're mapping it to the URL `ms-spr-demo-green.cfapps.io`. Let's push the application:

``` bash
Using manifest file manifest.yml

Creating ms-spr-demo-green... OK

1: ms-spr-demo-green
2: none
Subdomain> ms-spr-demo-green

1: cfapps.io
2: mattstine.com
3: none
Domain> 1

Creating route ms-spr-demo-green.cfapps.io... OK
Binding ms-spr-demo-green.cfapps.io to ms-spr-demo-green... OK
Uploading ms-spr-demo-green... OK
Starting ms-spr-demo-green... OK
-----> Downloaded app package (9.5M)
Installing java.
Downloading JDK...
Copying openjdk-1.7.0_25.tar.gz from the buildpack cache ...
Unpacking JDK to .jdk
Downloading Tomcat: apache-tomcat-7.0.41.tar.gz
Copying apache-tomcat-7.0.41.tar.gz from the buildpack cache ...
Unpacking Tomcat to .tomcat
Copying mysql-connector-java-5.1.12.jar from the buildpack cache ...
Copying postgresql-9.0-801.jdbc4.jar from the buildpack cache ...
Copying auto-reconfiguration-0.6.8.jar from the buildpack cache ...
-----> Uploading droplet (48M)
-----> Uploaded droplet
Checking ms-spr-demo-green...
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
  1/1 instances: 1 running
OK
```

We now have two instances of the application running, each of them using distinct routes:

{% img center /images/blue-green/BlueGreen2.png %} 

Now it's time for the magic to happen. We'll map our original URL route (`ms-spr-demo.cfapps.io`) to our "green" instance. This is accomplished very simply by using `cf`:

``` bash
cf map --app ms-spr-demo-green --host ms-spr-demo --domain cfapps.io
Binding ms-spr-demo.cfapps.io to ms-spr-demo-green... OK
```

The CF router immediately begins to load balance requests between each instance of the application, as shown here:

{% img center /images/blue-green/BlueGreen3.png %}

Now our router will send requests to `ms-spr-demo.cfapps.io` to both instances of the application, while `ms-spr-demo-green.cfapps.io` only services the "green" instance. Once we determine that all is well, it's time to stop routing requests to the "blue" instance. This is just as simple using `cf`:

``` bash
cf unmap --url ms-spr-demo.cfapps.io --app ms-spr-demo
Unbinding ms-spr-demo.cfapps.io from ms-spr-demo... OK
```

Our "blue" instance is now no longer receiving any web traffic:

{% img center /images/blue-green/BlueGreen4.png %}

We can now decomission our "blue" instance, or we can leave it around for a period of time in case we decide we need to roll back our changes. The important thing is that our customers suffered absolutely no downtime!



