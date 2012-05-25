---
date: '2007-05-11 22:17:00'
layout: post
slug: thursday-was-slow
status: publish
title: Thursday was slow...
wordpress_id: '32'
categories:
- ajax
- jsf
- NetBeans
- restaurants
- ruby
---

for me, not for JavaOne. Of course I was a good little programmer and used the schedule builder to sign up for all of my sessions. I edited them a bit after the first two keynotes. Even then, I had one lone session scheduled for Thursday morning (that wasn't really directly applicable to my work, it just looked interesting), and then two sessions in the afternoon, the first starting at 4:10. I decided to skip out on the morning session and do a little shopping at Pier 39. The highlight was a "Bucket of Boat Trash" and clam chowder at the Bubba Gump Shrimp Company. And yes, we Ole Miss Rebels are represented as far out as the end of Pier 39:  
  
[![](http://lh4.google.com/image/matt.stine/RkU1ErgYqKI/AAAAAAAAAMA/6BJ0mkhqgn8/s288/DSCN1117.JPG)](http://picasaweb.google.com/matt.stine/SanFranciscoJavaOne2007/photo#5063511710562298018)  
  
The two sessions that I attended were:  
  
- RubyTooling: State of the Art  
- Using Ajax with POJC (Plain Old JavaServer Faces Components)  
  
I attended the first session simply to get a little more detail on all of the hype surrounding Ruby tooling support in NetBeans 6. What I got was even more than I bargained for. The project leaders actually walked us through not only the features that were available, but how they were implemented. I had never really thought about the problems with providing code completion and refactoring with a dynamically-typed language. It was really cool to see the thought process that went into their solutions. I'd love to hear a similar discussion from the JetBrains guys, as the Ruby support in IntelliJ IDEA is quite good as well.  
  
For me the second session was the best of the conference for me up until that point, and arguably it still is after attending Friday's sessions. Craig McClanahan, of Struts fame, was the main speaker and was joined by his colleagues Matthew Bohm and Jayashri Visvanathan. What made this session so good for me was that they presented a problem - "How can I add Ajax behaviors to my JavaServer™ Faces technology based application, without throwing away my investment in existing component libraries?" - and then provided three different solutions to that problem - low, medium, and high level.  
  
The low level consisted of simply using the HTML event pass-through attributes that are implemented by most standard JSF components (onClick, onBlur, etc.). One could use an existing JavaScript framework such as Dojo to send an XMLHttpRequest and then map that request to a Servlet or JSF handler using a technology such as Shale Remoting. The response could be sent back as JSON data which could then be transformed into the desired UI update.  
  
The medium level consisted of actually extending existing JSF components and adding the desired Ajax behavior. Due to time constraints they didn't cover this solution in detail, but they did provide a link to a detailed discussion in the Java BluePrints catalog: https://blueprints.dev.java.net/bpcatalog/ee5/ajax/extendingRenderFunctionality.html  
  
The high level solutions addressed the following needs (copied directly from the slides):  
  
● Partial page submit—gather up a particular set of  
input element values, and send them to a bit of server  
side business logic  
● Partial page refresh—the business logic needs to  
refresh the content of one or more subtrees of the  
client side DOM  
● Synchronization—the benefits of synchronizing the  
server side state  
● Don’t repeat yourself (DRY)—reuse existing  
components and renderers for partial page updates  
  
To address these issues, the speakers highlighted two add-on frameworks:  
  
● Ajax4JSF (http://labs.jboss.com/portal/jbossajax4jsf)  
● Dynamic Faces (https://jsf-extensions.dev.java.net/)  
  
I was quite impressed with both of these frameworks. One of my colleagues is currently implementing Ajax behavior in a Facelets-based application using Ajax4JSF and he is quite happy with it. Dynamic Faces looked really awesome, especially its tooling support in NetBeans (actually I'm quite impressed with the overall JSF support in NetBeans - I'll definitely be adding it to my tool belt). The highlight of the presentation was Matt's video of him building an entire currency trading application in 28 minutes - except it was "fast-played" to finish in 3 1/2 minutes and set to techno music. Matt wowed us with his dancing abilities while we watched true RAD. The crowd went wild!
