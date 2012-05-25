---
date: '2007-12-13 14:02:00'
layout: post
slug: enterprise-java-and-the-changing-of-the-guard
status: publish
title: Enterprise Java and the Changing of the Guard
wordpress_id: '47'
categories:
- java
- thespringexperience2007
---

Greetings from [The Spring Experience 2007](http://www.thespringexperience.com) in Hollywood, FL. So far this has been a great conference - I'm currently waiting for my third session of the day to start. As much as I enjoy all of the hype and eye candy at JavaOne, I really get a lot more out of these smaller conferences as you're not running around stressed out trying to weave through thousands of geeks as you move from session to session (or more accurately, from session to queue!).  
  
I wanted to post a few insights that I gained from Rod Johnson's keynote last night (the title of this entry). The basis for Rod's presentation was a Gartner report entitled "[Trends in Platform Middleware: Disruption Is in Sight](http://www.gartner.com/DisplayDocument?ref=g_search&id=525420&subref=simplesearch)." To summarize, there are many converging forces in the enterprise Java space that are causing or will probably cause a great deal of disruption in the way we develop enterprise Java applications. These forces include our collective experience with application servers, the rediscovery of object-oriented or POJO-based programming, open source innovation, non-Java challengers (.NET, RoR), the rise of SOA, and Rich Internet Applications. Rod sees a big movement from the "old guard," which he described as the "J2EE Stovepipe Architecture" which was based primarily around the idea of distributed objects, to a world in which many of the highest volume Java applications don't even use a Java EE application server. An interesting point was the stagnation in job requirements for the big name application servers coupled with the BZ Research statistic that 64% of Java application developers are using Tomcat for production deployment. Yet another force for disruption is the emergence of OSGi, which allows developers to modularize applications and provides facilities for versioned components, fine-grained redeployment, and library conflict resolution - all features that Java EE currently does not and perhaps cannot address. Open source provides yet more disruption, as developers have become participants in rather than spectators of the development of Java enterprise solutions. Innovation is now primarily coming from the open source world. All of this culminated in the rather telegraphed conclusion that Rod (and also the Gartner report apparently) believes that Spring is uniquely positioned to provide the best value for ongoing Enterprise Java development in the face of the current trends. Of course, what would you expect him to say? :-) [  
](http://www.gartner.com/DisplayDocument?ref=g_search&id=525420&subref=simplesearch)  
I do tend to agree with him. The pace at which Spring is innovating and addressing the Java enterprise pain points far outpaces the pace of the JCP. It just takes too long to design specs by committee. One of the more wise things that Spring does is to learn from and in most cases accommodate the changing landscape of Java EE and integrate the best features directly into the framework. I for one am glad to be developing applications with Spring, as it makes it very easy to do the things I want to do and write quality code that isn't directly dependent on an infrastructure framework. Generally, it stays out of my way.  
  
More to come on the sessions I've attended thus far.
