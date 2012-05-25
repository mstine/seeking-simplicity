---
date: '2007-05-10 18:49:00'
layout: post
slug: wednesday-was-ajax-day
status: publish
title: Wednesday was AJAX Day
wordpress_id: '31'
categories:
- ajax
- GWT
- jMaki
- jsf
---

Not officially, but nearly every session I attended had something to do with AJAX:  
  
- Creating Amazing Web Interfaces with Ajax  
- jMaki: Web 2.0 App Building Made Easy  
- Fast, Beautiful, Easy: Pick Three - Building Web User Interfaces in the Java Programming Language with Google Web Toolkit  
- Killer JavaScript Technology Frameworks for Java Platform Developers: An Exploration of Prototype, Script.aculo.us, and Rico  
  
I have to say that I was rather impressed by what I saw.  
  
The first talk was by Ben and Dion, the Ajaxian guys. It was an appropriate way to start, as they gave a quick history overview of Ajax. One nice point they made was that Ajax really isn't about the acronym - it never was - it's about building killer websites. Who cares what the actual technology behind it is. They discussed a couple of what they seemed to consider the better frameworks available - Dojo and ExtJS. They then explored some amazing up and coming features, including offline support and 2D client side graphics manipulation.  
  
I was rather impressed with jMaki - in short it is a wrapper around many of the popular JavaScript frameworks available (Dojo, Yahoo UI, Script.aculo.us, Spry, Google), and makes them accessible to Java, PHP, and Ruby. It has excellent tool support in NetBeans and Eclipse. It provides protection from changes in the API's of all of these projects - you can mix and match frameworks and only be concerned about one API - jMaki's. It does the work of linking all of the widgets together and communicating amongst them and with the server side.  
  
The GWT talk was easily my favorite of the day. I'm extremely impressed with what these guys have done. I hadn't had much opportunity to look at GWT until now, and I really wish I had. I was initially skeptical about writing an entire application in Java and letting it generate HTML and JavaScript. I guess these guys knew that, because they're development philosophy addresses my concerns quite nicely:  
  
  
To radically improve the web experience for  
users by enabling developers to use existing  
Java tools to build no-compromise AJAX for  
any modern browser  
  
From what I can see, they deliver on their mission. They've optimized their code for speed and for browser specificity (e.g. from what I understand, if your client is using Firefox, you get Firefox optimized JavaScript, same for IE, etc.). You can use all of your favorite IDE features to build the code, including the debugger. I really want to try to make use of this toolkit in the near future.  
  
The final talk was less informative for me, but only because I had experience with most of the technologies already. The killer part of this was how the speaker extended existing JSF components and added Script.aculo.us effects. It really made his version of Yahoo maps shine.  
  
Ajax isn't going anywhere but up. I just left yet another Ajax/JSF session, which for me was the best session of the conference so far. In a later entry I'll tell you why.
