---
date: '2007-12-13 17:32:00'
layout: post
slug: improving-application-design-with-a-rich-domain-model-chris-richardson
status: publish
title: 'Improving Application Design with a Rich Domain Model: Chris Richardson'
wordpress_id: '48'
categories:
- java
- thespringexperience2007
---

This was my very first session of the conference. I've really been looking forward to it. I became pretty excited about object-oriented programming when I first really learned it in my computer simulation course at Ole Miss. In that course we built discrete event simulation programs using collaborating Java threads. Each thread implemented an object from the domain model representing a particular simulation problem. Once I got into the "working world," I found that the architecture described by Rod Johnson as the "J2EE stove pipe" had made my OO skills essentially unusable in the projects on which I was required to work on a day-to-day basis. My recent experiences developing in a more POJO-based style, particularly with JPA/Hibernate, have allowed me to get closer to my preferred OO-style, but the business logic still lived in the service-tier. My domain model was still essentially "anemic" to borrow from Martin Fowler.  
  
When I read the description for this session I was elated! Finally I could see that someone was building real enterprise applications using object-oriented techniques! I've never quite been able to articulate what I felt was wrong with the way we're building applications in my group. I was fortunate enough to have breakfast with Chris Richardson prior to the session, and he said that addressing my inability to articulate those problems was one of the goals of his talk. Awesome! On to the session...  
  
Chris began with a tour through his object-based programming experience, beginning in LISP, followed by C++ and then Java. Around 1999 he got into EJB programming, and described his experience as "Applications were still built from objects, but those objects were very different..." Basically the EJB paradigm caused us to abandon object-oriented programming in favor of procedural programming. Why? Well, for all of the pros surrounding EJB, it made writing object-oriented code difficult, and in some cases, impossible. The EJB/procedural style works very well for simple business logic, but it doesn't scale well as business logic becomes more complex. The result is a few monolithic transactional classes containing hundreds to thousands of lines of code embedded in very long methods. Hence, the legacy of EJB (and I quote):  


  * Java is an object-oriented language, AND...  

  * Object-oriented design is a better way to tackle complexity, YET...  

  * Many complex enterprise Java applications are written in a procedural style  

Chris then moved on to discuss the rich domain model pattern, which is really nothing new at all. It's just good OO design. Most of your classes (or entities) correspond to real world concepts, and the business logic is spread amongst them. Classes are true objects: they contain state and behavior. The goal is to push as much business logic as possible down into the domain objects, which is exactly where it belongs. What do you get? Better maintainability, better testability, better reusability, and a better shared understanding of your domain. Not only that, but your code is quantifiably simpler! Who doesn't want that? The main drawback to this is that you have to have good OO design skills to make it happen - which is exactly what the EJB programming model has made scarce.  
  
Chris followed this up by discussing the building blocks of domain models, concepts he derived from Eric Evans' book Domain Driven Design. In a nutshell, they are Entities, Value Objects, Aggregates, Services, Factories, and Repositories - the reader can get the book to find out what these are. Interesting notes for me include the fact that Repositories are nearly equivalent to DAO's, especially the way I typically implement them. The concept of a Service is far different from what I'm used to. I "grew up" writing services according to the EJB Session Bean/Facade model, where essentially all of your business logic resides in service methods. Chris defined a service as a class implementing only logic that cannot be put in a single entity. They are actually quite thin!  
  
Next came the discussion of frameworks and their role. In short, frameworks act as an enabler to rich domain models when used properly. Your domain model should be implemented purely as Plain Old Java Objects (POJO's) in that they don't implement any infrastructure interfaces or call infrastructure API's. The jury is still out on annotations - Chris argues that they still violate the POJO concept, I'm not so sure. You then wire your model together using dependency injection, handle crosscutting concerns (transaction management, security, logging, etc.) with Aspect-Oriented Programming (AOP), and use object/relational mapping (ORM) for persistence. What you often run into as obstacles are that many of these frameworks, particularly web and ORM frameworks, force you to introduce "smells" into your code for the framework to use them. Examples include requirements for public default constructors and JavaBean-style setters.  
  
Chris closed with an awesome discussion of common code smells and the refactorings used to eliminate them. I was most interested in Feature Envy, which is where you have methods that are far too interested in data belonging to other classes. This is very common in session facades, and the healing factor is to push that logic down into the appropriate domain classes. Another interesting smell was Primitive Obsession, where code uses built-in types (such as String and Integer) instead of application classes to represent state. The refactoring in this case is to introduce Value Objects, which are immutable, validated objects representing domain concepts (e.g. a shipping address).  
  
Chris closed with the charge to begin refactoring our procedural designs into a rich domain model on MONDAY! Oh, how I can't wait to begin. Good job Chris!
