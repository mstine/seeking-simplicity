---
date: '2009-06-02 15:56:01'
layout: post
slug: java-technology-on-google-app-engine-communityone-2009-lightning-talk
status: publish
title: 'Java™ Technology on Google App Engine: #CommunityOne 2009 Lightning Talk'
wordpress_id: '184'
categories:
- javaone2009
---

My second lightning talk at Community One West 2009 revolved around the still relatively recent announcement [that Java is now supported on Google's App Engine](http://code.google.com/appengine/docs/java/overview.html).

What exactly is [Google App Engine](http://code.google.com/appengine)? It's none other than a way that you can run your Java technology-based applications on Google's massive infrastructure.

As far as the "Geekxecutive Summary," Google has provided:



	
  * [A Java 6 Virtual Machine](http://code.google.com/appengine/docs/java/runtime.html) with a [Class Whitelist](http://code.google.com/appengine/docs/java/jrewhitelist.html). Not all classes in the standard JDK library are available. A notable missing member is the java.awt.* library, so if you're using that for image processing, find an alternative (see one later) before porting to the App Engine.

	
  * The Java Servlet Standard by way of the [Jetty container](http://www.mortbay.org/jetty/).

	
  * [A secured sandbox environment](http://code.google.com/appengine/docs/java/runtime.html#The_Sandbox), meaning 1) no spawned threads; 2) no arbitrary socket connections (you can use HTTP/HTTPS, see below); 3) No writing to the local filesystem - makes sense because you don't know where you are on any given request.



Several services are provided to you by Google as a Java developer:


	
  * [The Datastore](http://code.google.com/appengine/docs/java/datastore/) - a schemaless object datastore, with a query engine and atomic transactions. This is accessible via Java Data Objects (JDO) 2.3, Java Persistence API (JPA) 1.0, or Google's own low-level API's.

	
  * A [Memcache](http://code.google.com/appengine/docs/java/memcache/) service for caching whatever objects you like, accessible via JCache JSR 107.

	
  * [A URL fetch service](http://code.google.com/appengine/docs/java/urlfetch/) for accessing data on other systems, accessible via java.net.URLConnection.

	
  * [A Mail (JavaMail) service](http://code.google.com/appengine/docs/java/mail/).

	
  * [An Image processing service](http://code.google.com/appengine/docs/java/images/), a likely candidate for replacing AWT for image processing. Be careful, it binds you to Google's services. You may want to write a wrapper for it.

	
  * [Google Accounts for authentication](http://code.google.com/appengine/docs/java/users/). Should you leverage this, anyone with a Google account can authenticate into your app.



The nice thing about Google's services is that where possible they always provided a way in through standards-based Java API's.

Now, on to tools:


	
  * [An Apache Ant component](http://code.google.com/appengine/docs/java/tools/ant.html) for your automated build process.

	
  * [A very nice plugin](http://code.google.com/appengine/docs/java/tools/eclipse.html) for the Eclipse IDE.

	
  * [The App Engine SDK](http://code.google.com/appengine/docs/java/gettingstarted/installing.html) - Java API's into all of Google's services.

	
  * [A Development Server](http://code.google.com/appengine/docs/java/tools/devserver.html) so that you can run a simulated Google environment on your local machine.

	
  * [AppCfg](http://code.google.com/appengine/docs/java/tools/uploadinganapp.html) for command-line interaction with the App Engine platform.



Probably the most significant piece to this puzzle is the fact that Google is not simply supporting Java the language, but Java the platform. It's a standard JVM, so other languages running on the JVM will happily run on Google App Engine as well. At this writing, I knew of the following languages that are known to run:


	
  * Java

	
  * Groovy (and Grails as of 1.1.1)

	
  * JRuby

	
  * Scala

	
  * Clojure

	
  * Beanshell

	
  * JavaScript (Rhino)

	
  * Jython



Here's the screencast from my talk. Enjoy!

[youtube=http://www.youtube.com/watch?v=hotbwYStwH8&hl;=en&fs;=1]
