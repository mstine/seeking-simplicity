---
date: '2007-12-14 15:40:00'
layout: post
slug: ajax-integration-guide-for-2008-jeremy-grelle
status: publish
title: 'Ajax Integration Guide for 2008: Jeremy Grelle'
wordpress_id: '51'
categories:
- ajax
- spring
- thespringexperience2007
---

This talk looked really exciting going in. Jeremy is a member of the Spring Web Products team and the lead of Spring Faces, as well as a member of the JSF 2.0 Expert Group. He describes himself as the resident Ajax freak at SpringSource. Jeremy began by discussing the current Ajax landscape, highlighting what he believes are the important features a successful framework must deliver, and then identifying the frameworks he considers to be the cream of the crop. No surprises here on the list: Prototype/Scriptaculous, jQuery, YUI, Ext, and Dojo. He then gave a brief high-level overview of each framework - I really didn't learn anything new here, as I've been following these frameworks for quite awhile. He mentioned the commonality that all of these frameworks are about more than asynchronous client-server communication - they're about improving your client-side code and user experience. See Dion's May 2007 comment that Ajax "was never about the acronym - it's about building killer websites." What is good about all of these frameworks is that you can use all of them in an unobtrusive manner and progressively enhance your UI. This gives you the ability to only Ajaxify what you need and also get graceful degradation.  
  
In the second portion of his talk, Jeremy talked about connecting to Spring. I was really hoping that the "goods" would be delivered here. He started by mentioning DWR, which I think is a really nice framework. It merited 2 slides and no demo (even though there was a slide mentioning a demo). I don't think Jeremy really cares for it or GWT, as he sort of lumped them together and didn't mention them again. I do appreciate his point that neither really allows you to harvest any existing controlling infrastructure that we may have built using Spring MVC or another framework.  
  
Jeremy then moved into what he considers the better way to do it, which is the progressive enhancement I mentioned above. He took the existing Spring Booking MVC + Web Flow example application and ajaxified portions of it. He did this using an interesting combination of Apache Tiles, some Javascript "pseudo-AOP" he lifted from the SpringFaces project, and a bit of Dojo. It seemed to work really well, but it wasn't immediately clear how this type of technology was going to be made available to us as developers. He said something along the lines of "you don't have to be using JSF to use this Javascript, so we probably need to change the name to something else." I hope they figure out soon so I can get my hands on it. :-)  
  
He closed by briefly discussing REST's place in the future of Spring Web. The ideal scenario, according to Jeremy, is that we will be able to request different representations of the same resource. For example, we could request HTML for full page display or partial DOM updates, and JSON for intelligent widgets like the Dojo table.  
  
It looks like a lot of exciting things will happen in 2008. I hope I don't have to wait too long.
