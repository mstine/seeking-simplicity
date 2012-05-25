---
date: '2007-12-14 15:20:00'
layout: post
slug: whats-new-in-spring-25-juergen-hoeller
status: publish
title: 'What''s New in Spring 2.5? : Juergen Hoeller'
wordpress_id: '50'
categories:
- java
- spring
- thespringexperience2007
---

This was my second talk of TSE 2007. I have to admit I chose it by process of elimination - none of the second session talks particularly jumped out and grabbed me like Chris Richardson's talk.  
  
Juergen is the project lead for the Spring Framework, so he was the obvious choice to give this talk. Juergen split it up into three sections:  


  1. Platforms
  2. Annotation Configuration
  3. AspectJ Support
To break down part one as quickly as possible, Spring supports virtually EVERYTHING. You get JDK 1.4 all the way to Java 6, including all of the new API's in Java 6. It fully supports Java EE 5 while remaining backward compatible all the way to J2EE 1.3.  You can now deploy a Spring ApplicationContext as a RAR file, and you also get full JCA 1.5 support (if you want to know what that means, don't ask me :-)). Quite notable was the fact that they have worked with IBM to support WebSphere's proprietary transaction manager. Also interesting was Spring's continuing strong support for OSGi as an alternative enterprise runtime.  
  
The annotation configuration part of the talk was quite fascinating. I'm something of an annotations junkie - I have to be careful about that. At any rate, it was so good that I was interested enough to attend Juergen and Mark Fisher's later talk that was completely dedicated to the subject. I'll leave the details for my entry on that talk.  
  
The final portion, on AspectJ support, was equally fascinating. The first new feature was the ability to advise specific beans by name rather than by type using AspectJ. This was made even sweeter by the fact that you can use pattern matching in your  definition.  Very exciting was the ability to do AspectJ load-time weaving, meaning you can use the power of AspectJ at runtime without involving the AspectJ compiler. Unfortunately, this isn't available across all appservers. I didn't get a chance to ask and comfirm, but it seems that JBoss (our primary appserver) does not support this. GlassFish, Weblogic, OC4J, and Tomcat were all mentioned as being able to handle it.  
  
The most interesting part of this talk for me was the ability to annotate a class as @Configurable, and then do dependency injection on it even thought it isn't managed by Spring. You could do this in Spring 2.0, but you had to use the AspectJ compiler. Now, with load-time weaving, you can do this at runtime! I really could have used this recently when I wanted to inject a single dependency into a class that was really overkill to manage as a Spring bean. Since I'm using Spring 2.5 in this project, I can go back and try this feature. Exciting!
